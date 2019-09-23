package user

import (
	"context"
	"database/sql"

	"github.com/yezarela/go-lambda/models"
	"github.com/yezarela/go-lambda/pkg/utils"
)

// Usecase represent the user's usecases
type Usecase struct {
	db       *sql.DB
	userRepo Repository
}

// NewUserUsecase will create new Usecase object
func NewUserUsecase(db *sql.DB, r Repository) *Usecase {
	return &Usecase{
		db:       db,
		userRepo: r,
	}
}

// ListUser ...
func (m *Usecase) ListUser(ctx context.Context, params ...ListUserParams) ([]*models.User, error) {
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
		return nil, err
	}

	return res, nil
}

// GetByID ...
func (m *Usecase) GetByID(ctx context.Context, id uint) (*models.User, error) {

	res, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateUser ...
func (m *Usecase) CreateUser(ctx context.Context, p *models.User) (*models.User, error) {

	// Start transaction
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := m.userRepo.CreateUser(ctx, tx, p)
	if err != nil {
		return nil, err
	}

	// Commit transactions / end
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	res, err := m.userRepo.GetUser(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateUser ...
func (m *Usecase) UpdateUser(ctx context.Context, p *models.User) (*models.User, error) {

	// Start transaction
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := m.userRepo.UpdateUser(ctx, tx, p)
	if err != nil {
		return nil, err
	}

	// Commit transactions / end
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteUser ...
func (m *Usecase) DeleteUser(ctx context.Context, id uint) error {

	existedUser, err := m.userRepo.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if existedUser == nil {
		return models.ErrNotFound
	}

	return m.userRepo.DeleteUser(ctx, id)
}
