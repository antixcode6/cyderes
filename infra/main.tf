terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
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

  source_dir  = "${path.module}/../pkg/main"
  output_path = "${path.module}/cyderes.zip"
}
resource "aws_s3_object" "lambda_cyderes_test" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "cyderes.zip"
  source = data.archive_file.lambda_cyderes_test.output_path

  etag = filemd5(data.archive_file.lambda_cyderes_test.output_path)
}
