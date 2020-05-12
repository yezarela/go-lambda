package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/domain/user"
	"github.com/yezarela/go-lambda/model"
	"github.com/yezarela/go-lambda/pkg/conn"
	"github.com/yezarela/go-lambda/pkg/validator"

	_ "github.com/go-sql-driver/mysql"
)

var userUsecase user.Usecase

func init() {
	db := conn.NewSQLConnection()
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
		return model.APIErrResponse(400, model.ErrInvalidParameters)
	}

	// Marshall request body to body variable
	err := json.Unmarshal([]byte(evt.Body), &body)
	if err != nil {
		return model.APIServerError(err)
	}

	// Validate with govalidator
	err = validator.ValidateStruct(body)
	if err != nil {
		return model.APIErrResponse(400, err)
	}

	// Create user
	userData := model.User{
		Name:  body.Name,
		Email: body.Email,
	}

	res, err := userUsecase.CreateUser(ctx, &userData)
	if err != nil {
		return model.APIServerError(err)
	}

	return model.APIResponse(200, res)
}

func main() {
	lambda.Start(handler)
}
