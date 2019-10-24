// +build test

package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yezarela/go-lambda/domain/user"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	t.Run("Success", func(t *testing.T) {

		rows := sqlmock.
			NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
			AddRow(1, "john", "john@doe.com", time.Now(), time.Now())

		mock.ExpectQuery(user.GetUserQuery).
			WithArgs(1).
			WillReturnRows(rows)

		repo := user.NewRepository(db)

		res, err := repo.GetUser(context.Background(), uint(1))

		assert.NoError(t, err)
		assert.NotNil(t, res)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})

	t.Run("Error", func(t *testing.T) {

		mock.ExpectQuery(user.GetUserQuery).WillReturnError(errors.New("some error"))

		repo := user.NewRepository(db)

		res, err := repo.GetUser(context.Background(), uint(5))

		assert.Error(t, err)
		assert.Nil(t, res)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}

// go test domain/user/repository_test.go
