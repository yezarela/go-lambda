package main

import (
	"context"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yezarela/go-lambda/domain/user"
	"github.com/yezarela/go-lambda/models"
	"github.com/yezarela/go-lambda/pkg/core"

	_ "github.com/go-sql-driver/mysql"
)

var userUsecase *user.Usecase

func init() {
	db := core.OpenSQLConnection()
	userRepo := user.NewUserRepository(db)
	userUsecase = user.NewUserUsecase(db, userRepo)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	qs := req.QueryStringParameters

	// Parse limit & offset
	limit, _ := strconv.Atoi(qs["limit"])
	offset, _ := strconv.Atoi(qs["offset"])

	param := user.ListUserParams{
		Limit:         uint(limit),
		Offset:        uint(offset),
		SortBy:        qs["sort_by"],
		SortDirection: qs["sort_direction"],
	}

	// Get list of users
	users, err := userUsecase.ListUser(ctx, param)
	if err != nil {
		return models.APIServerError(err)
	}

	return models.APIResponse(200, users)
}

func main() {
	lambda.Start(handler)
}
