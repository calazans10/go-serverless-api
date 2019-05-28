package main

import (
	"github.com/aws-lambda/go-serverless-api/pkg/api"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.Handler)
}
