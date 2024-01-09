package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yezarela/go-lambda/domain"
	mock_domain "github.com/yezarela/go-lambda/domain/mock"
	"github.com/yezarela/go-lambda/utils/jsonutil"
)

func TestHandler(t *testing.T) {

	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	mockUserUsecase := mock_domain.NewMockUserUsecase(ctrl)

	userUsecase = mockUserUsecase

	tests := []struct {
		name         string
		qs           map[string]string
		usecaseParam domain.ListUserParams
		usecaseRes   []*domain.User
		wantRes      string
		wantStatus   int
	}{
		{
			name: "Successful Request",
			qs: map[string]string{
				"limit":  "10",
				"offset": "0",
			},
			usecaseParam: domain.ListUserParams{
				Limit:  10,
				Offset: 0,
			},
			usecaseRes: []*domain.User{
				{
					ID:    1,
					Name:  "test",
					Email: "test@example.com",
				},
			},
			wantRes: `[{
				"id": 1,
				"name": "test",
				"email": "test@example.com",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "0001-01-01T00:00:00Z"
			}]`,
			wantStatus: 200,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			mockUserUsecase.EXPECT().ListUser(ctx, tt.usecaseParam).Return(tt.usecaseRes, nil)

			resp, err := handler(ctx, events.APIGatewayProxyRequest{
				QueryStringParameters: tt.qs,
			})

			res, _ := jsonutil.Compact(resp.Body)
			want, _ := jsonutil.Compact(tt.wantRes)

			assert.NoError(t, err)
			assert.Equal(t, resp.StatusCode, tt.wantStatus)
			assert.Equal(t, res, want)
		})
	}
}
