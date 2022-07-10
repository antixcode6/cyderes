terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.22.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }
  }

  required_version = "~> 1.0"
}

provider "aws" {
  region = var.aws_region
}
#Deploy binary to S3
resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "cyderes-lambda-project"
}

data "archive_file" "lambda_cyderes_test" {
  type = "zip"

  source_dir  = "${path.module}/../pkg/"
  output_path = "${path.module}/cyderes-s3.zip"
}
resource "aws_s3_object" "lambda_cyderes_test" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "cyderes-s3.zip"
  source = data.archive_file.lambda_cyderes_test.output_path

  etag = filemd5(data.archive_file.lambda_cyderes_test.output_path)
}
resource "aws_lambda_function" "lambda_cyderes_test" {
  function_name = "cyderes-test"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_object.lambda_cyderes_test.key

  runtime = "go1.x"
  handler = "main"

  source_code_hash = data.archive_file.lambda_cyderes_test.output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_iam_role" "lambda_exec" {
  name = "cyderes"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}
resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_apigatewayv2_api" "lambda" {
  name          = "cyderes"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.lambda.id

  name        = "serverless_lambda_stage"
  auto_deploy = true

}
resource "aws_apigatewayv2_integration" "lambda_cyderes_test" {
    request_parameters     = {
        "overwrite:querystring.name" = "req"
        }
  api_id = aws_apigatewayv2_api.lambda.id
  integration_uri    = aws_lambda_function.lambda_cyderes_test.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "lambda_cyderes_test" {
  api_id = aws_apigatewayv2_api.lambda.id
    
  route_key = "GET /api/v1/{query}"

  target    = "integrations/${aws_apigatewayv2_integration.lambda_cyderes_test.id}"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_cyderes_test.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}