package models

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// APIResponse ...
func APIResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	bytes, _ := json.Marshal(&body)

	return events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: statusCode,
	}, nil
}

// APIErrResponse ...
func APIErrResponse(statusCode int, err error) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: statusCode,
	}, nil
}

// APIServerError ...
func APIServerError(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)

	return events.APIGatewayProxyResponse{
		Body:       "Internal server error",
		StatusCode: 500,
	}, err
}
