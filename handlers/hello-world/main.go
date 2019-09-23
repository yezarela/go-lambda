package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/models"

	_ "github.com/go-sql-driver/mysql"
)

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return models.APIResponse(200, "Hello world!")
}

func main() {
	lambda.Start(handler)
}
