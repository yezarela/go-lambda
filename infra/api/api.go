package api

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
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

	printStackTrace(err)

	return events.APIGatewayProxyResponse{
		Body:       "Internal server error",
		StatusCode: 500,
	}, err
}

type stacktracer interface {
	StackTrace() errors.StackTrace
}

type causer interface {
	Cause() error
}

func printStackTrace(err error) {

	var errStack errors.StackTrace

	for err != nil {
		// Find the earliest error.StackTrace
		if t, ok := err.(stacktracer); ok {
			errStack = t.StackTrace()
		}
		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			break
		}
	}
	if errStack != nil {
		fmt.Println(err)
		fmt.Printf("%+v\n", errStack)
	} else {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}
