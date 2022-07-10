package main

import (
	"github.com/antixcode6/cyderes/pkg/api"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.Ingest)
}
