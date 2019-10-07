package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/domain/user"
	"github.com/yezarela/go-lambda/models"
	"github.com/yezarela/go-lambda/pkg/core"
	"github.com/yezarela/go-lambda/pkg/validator"

	_ "github.com/go-sql-driver/mysql"
)

var userUsecase *user.Usecase

func init() {
	db := core.OpenSQLConnection()
	userRepo := user.NewRepository(db)
	userUsecase = user.NewUsecase(db, userRepo)
}

type params struct {
	Name  string `json:"name" valid:"required"`
	Email string `json:"email" valid:"required,email~invalid email format"`
}

func handler(ctx context.Context, evt events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body params

	if len(evt.Body) <= 0 {
		return models.APIErrResponse(400, models.ErrInvalidParameters)
	}

	// Marshall request body to body variable
	err := json.Unmarshal([]byte(evt.Body), &body)
	if err != nil {
		return models.APIServerError(err)
	}

	// Validate with govalidator
	err = validator.ValidateStruct(body)
	if err != nil {
		return models.APIErrResponse(400, err)
	}

	// Create user
	userData := models.User{
		Name:  body.Name,
		Email: body.Email,
	}

	res, err := userUsecase.CreateUser(ctx, &userData)
	if err != nil {
		return models.APIServerError(err)
	}

	return models.APIResponse(200, res)
}

func main() {
	lambda.Start(handler)
}
