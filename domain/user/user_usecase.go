package user

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/yezarela/go-lambda/models"
	"github.com/yezarela/go-lambda/pkg/utils"
)

// Usecase ...
type Usecase struct {
	db       *sql.DB
	userRepo *Repository
}

// NewUsecase ...
func NewUsecase(db *sql.DB, r *Repository) *Usecase {
	return &Usecase{
		db:       db,
		userRepo: r,
	}
}

// ListUser ...
func (m *Usecase) ListUser(ctx context.Context, params ...ListUserParams) ([]*models.User, error) {
	op := "user.Usecase.ListUser"

	param := ListUserParams{
		SortBy:        "date",
		SortDirection: "desc",
		Limit:         10,
		Offset:        0,
	}

	if p := params[0]; len(params) > 0 {
		param.SortBy = utils.Strdef(p.SortBy, param.SortBy)
		param.SortDirection = utils.Strdef(p.SortDirection, param.SortDirection)
		param.Limit = utils.Uintdef(p.Limit, param.Limit)
		param.Offset = utils.Uintdef(p.Offset, param.Offset)
	}

	res, err := m.userRepo.ListUser(ctx, param)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// GetByID ...
func (m *Usecase) GetByID(ctx context.Context, id uint) (*models.User, error) {
	op := "user.Usecase.GetByID"

	res, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// CreateUser ...
func (m *Usecase) CreateUser(ctx context.Context, p *models.User) (*models.User, error) {
	op := "user.Usecase.CreateUser"

	// Start transaction
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer tx.Rollback()

	id, err := m.userRepo.CreateUser(ctx, tx, p)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	// Commit transactions / end
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	res, err := m.userRepo.GetUser(ctx, uint(id))
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// UpdateUser ...
func (m *Usecase) UpdateUser(ctx context.Context, p *models.User) (*models.User, error) {
	op := "user.Usecase.UpdateUser"

	// Start transaction
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer tx.Rollback()

	res, err := m.userRepo.UpdateUser(ctx, tx, p)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	// Commit transactions / end
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// DeleteUser ...
func (m *Usecase) DeleteUser(ctx context.Context, id uint) error {
	op := "user.Usecase.DeleteUser"

	existedUser, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if existedUser == nil {
		return models.ErrNotFound
	}

	return m.userRepo.DeleteUser(ctx, id)
}
