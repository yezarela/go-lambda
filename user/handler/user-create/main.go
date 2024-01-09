package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/domain"
	"github.com/yezarela/go-lambda/infra/api"
	"github.com/yezarela/go-lambda/infra/database"
	_userRepo "github.com/yezarela/go-lambda/user/repository"
	_userUsecase "github.com/yezarela/go-lambda/user/usecase"
	"github.com/yezarela/go-lambda/utils/validator"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	userRepo    domain.UserRepository
	userUsecase domain.UserUsecase
)

type params struct {
	Name  string `json:"name" valid:"required"`
	Email string `json:"email" valid:"required,email~invalid email format"`
}

func handler(ctx context.Context, evt events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body params

	if len(evt.Body) <= 0 {
		return api.APIErrResponse(400, domain.ErrInvalidParameters)
	}

	// Marshall request body to body variable
	err := json.Unmarshal([]byte(evt.Body), &body)
	if err != nil {
		return api.APIServerError(err)
	}

	// Validate with govalidator
	err = validator.ValidateStruct(body)
	if err != nil {
		return api.APIErrResponse(400, err)
	}

	// Create user
	userData := domain.User{
		Name:  body.Name,
		Email: body.Email,
	}

	res, err := userUsecase.CreateUser(ctx, &userData)
	if err != nil {
		return api.APIServerError(err)
	}

	return api.APIResponse(200, res)
}

func main() {
	db = database.NewMySQLConnection(os.Getenv("DBDataSourceName"))
	userRepo = _userRepo.NewMysqlRepository(db)
	userUsecase = _userUsecase.NewUsecase(userRepo)

	lambda.Start(handler)
}
