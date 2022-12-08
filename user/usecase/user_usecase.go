package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/yezarela/go-lambda/domain"
	"github.com/yezarela/go-lambda/utils/tern"
)

type usecase struct {
	userRepo domain.UserRepository
}

// NewUsecase ...
func NewUsecase(r domain.UserRepository) domain.UserUsecase {
	return &usecase{
		userRepo: r,
	}
}

// ListUser ...
func (m *usecase) ListUser(ctx context.Context, params ...domain.ListUserParams) ([]*domain.User, error) {
	op := "user.Usecase.ListUser"

	param := domain.ListUserParams{
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
func (m *usecase) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	op := "user.Usecase.GetByID"

	res, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

// CreateUser ...
func (m *usecase) CreateUser(ctx context.Context, p *domain.User) (*domain.User, error) {
	op := "user.Usecase.CreateUser"

	id, err := m.userRepo.CreateUser(ctx, p)
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
func (m *usecase) UpdateUser(ctx context.Context, p *domain.User) (*domain.User, error) {
	op := "user.Usecase.UpdateUser"

	res, err := m.userRepo.UpdateUser(ctx, p)
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
		return domain.ErrNotFound
	}

	return m.userRepo.DeleteUser(ctx, id)
}
