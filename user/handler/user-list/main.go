package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/domain"
	"github.com/yezarela/go-lambda/infra/api"
	"github.com/yezarela/go-lambda/infra/database"
	_userRepo "github.com/yezarela/go-lambda/user/repository"
	_userUsecase "github.com/yezarela/go-lambda/user/usecase"

	_ "github.com/go-sql-driver/mysql"
)

var userUsecase domain.UserUsecase

func init() {
	db := database.NewMySQLConnection()
	userRepo := _userRepo.NewMysqlRepository(db)
	userUsecase = _userUsecase.NewUsecase(userRepo)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	qs := req.QueryStringParameters

	// Parse limit & offset
	limit, _ := strconv.Atoi(qs["limit"])
	offset, _ := strconv.Atoi(qs["offset"])

	param := domain.ListUserParams{
		Limit:         uint(limit),
		Offset:        uint(offset),
		SortBy:        qs["sort_by"],
		SortDirection: qs["sort_direction"],
	}

	// Get list of users
	users, err := userUsecase.ListUser(ctx, param)
	if err != nil {
		return api.APIServerError(err)
	}

	return api.APIResponse(200, users)
}

func main() {
	lambda.Start(handler)
}
