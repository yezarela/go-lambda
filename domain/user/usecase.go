package user

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/yezarela/go-lambda/model"
	"github.com/yezarela/go-lambda/pkg/tern"
)

// Usecase ...
type Usecase interface {
	ListUser(ctx context.Context, p ...ListUserParams) ([]*model.User, error)
	GetByID(ctx context.Context, id uint) (*model.User, error)
	CreateUser(ctx context.Context, p *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, p *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type usecase struct {
	db       *sql.DB
	userRepo Repository
}

// NewUsecase ...
func NewUsecase(db *sql.DB, r Repository) Usecase {
	return &usecase{
		db:       db,
		userRepo: r,
	}
}

// ListUser ...
func (m *usecase) ListUser(ctx context.Context, params ...ListUserParams) ([]*model.User, error) {
	op := "user.Usecase.ListUser"

	param := ListUserParams{
		SortBy:        "date",
		SortDirection: "desc",
		Limit:         10,
		Offset:        0,
	}

	if p := params[0]; len(params) > 0 {
		param.SortBy = tern.String(p.SortBy, param.SortBy)
		param.SortDirection = tern.String(p.SortDirection, param.SortDirection)
		param.Limit = tern.Uint(p.Limit, param.Limit)
		param.Offset = tern.Uint(p.Offset, param.Offset)
	}

	res, err := m.userRepo.ListUser(ctx, param)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// GetByID ...
func (m *usecase) GetByID(ctx context.Context, id uint) (*model.User, error) {
	op := "user.Usecase.GetByID"

	res, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// CreateUser ...
func (m *usecase) CreateUser(ctx context.Context, p *model.User) (*model.User, error) {
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
func (m *usecase) UpdateUser(ctx context.Context, p *model.User) (*model.User, error) {
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
func (m *usecase) DeleteUser(ctx context.Context, id uint) error {
	op := "user.Usecase.DeleteUser"

	existedUser, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if existedUser == nil {
		return model.ErrNotFound
	}

	return m.userRepo.DeleteUser(ctx, id)
}
