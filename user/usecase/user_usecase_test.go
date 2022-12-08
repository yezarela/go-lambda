package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/yezarela/go-lambda/domain"
	mock_domain "github.com/yezarela/go-lambda/domain/mock"
)

func TestGetByID(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	mockUserRepo := mock_domain.NewMockUserRepository(ctrl)
	mockUser := domain.User{ID: 1}

	t.Run("Success", func(t *testing.T) {

		mockUserRepo.EXPECT().GetUser(ctx, mockUser.ID).Return(&mockUser, nil)

		usecase := NewUsecase(mockUserRepo)

		res, err := usecase.GetByID(ctx, mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.ID, res.ID)

	})

	t.Run("Error", func(t *testing.T) {

		mockUserRepo.EXPECT().GetUser(ctx, uint(12)).Return(nil, errors.New("noop"))

		usecase := NewUsecase(mockUserRepo)

		res, err := usecase.GetByID(ctx, uint(12))

		assert.Error(t, err)
		assert.Nil(t, res)
	})

}

// func TestCreateUser(t *testing.T) {
// 	ctrl, ctx := gomock.WithContext(context.Background(), t)
// 	defer ctrl.Finish()

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	t.Run("Successful", func(t *testing.T) {

// 		mock.ExpectBegin()
// 		// mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
// 		// mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
// 		mock.ExpectCommit()

// 		mockUserRepo := userMock.NewMockRepository(ctrl)
// 		mockUserID := int64(1)
// 		mockUser := model.User{ID: uint(mockUserID)}

// 		mockUserRepo.EXPECT().
// 			CreateUser(ctx, gomock.Any(), &mockUser).
// 			Return(mockUserID, nil)

// 		usecase := user.NewUsecase(db, mockUserRepo)

// 		res, err := usecase.CreateUser(ctx, &mockUser)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, res)
// 		// assert.Equal(t, mockUser.ID, res.ID)

// 	})

// 	// t.Run("Error", func(t *testing.T) {

// 	// 	mockUserRepo.EXPECT().GetUser(ctx, uint(12)).Return(nil, errors.New("noop"))

// 	// 	usecase := user.NewUsecase(nil, mockUserRepo)

// 	// 	res, err := usecase.GetByID(ctx, uint(12))

// 	// 	assert.Error(t, err)
// 	// 	assert.Nil(t, res)
// 	// })

// }

// go test domain/user/usecase_test.go
