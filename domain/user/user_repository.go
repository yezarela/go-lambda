package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yezarela/go-lambda/models"
)

// Repository represent the user's repository
type Repository interface {
	ListUser(ctx context.Context, p ListUserParams) ([]*models.User, error)
	GetUser(ctx context.Context, id uint) (*models.User, error)
	CreateUser(ctx context.Context, tx *sql.Tx, p *models.User) (int64, error)
	UpdateUser(ctx context.Context, tx *sql.Tx, p *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository will create an object that represent the Repository interface
func NewUserRepository(db *sql.DB) Repository {
	return &userRepository{db}
}

func (m *userRepository) fetchUser(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*models.User{}
	for rows.Next() {
		s := models.UserScan{}

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Email,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		data := &models.User{}
		data = data.FromScan(s)

		items = append(items, data)
	}

	return items, nil
}

// ListUserParams ...
type ListUserParams struct {
	SortBy        string
	SortDirection string
	Limit         uint
	Offset        uint
}

// ListUser ...
func (m *userRepository) ListUser(ctx context.Context, param ListUserParams) ([]*models.User, error) {

	switch param.SortBy {
	case "date":
		param.SortBy = "created_at"
	default:
		param.SortBy = "created_at"
	}

	switch param.SortDirection {
	case "asc":
		param.SortDirection = "ASC"
	default:
		param.SortDirection = "DESC"
	}

	query := fmt.Sprintf(listUserQuery,
		param.SortBy,
		param.SortDirection,
		param.Limit,
		param.Offset,
	)

	items, err := m.fetchUser(ctx, query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// GetUser ...
func (m *userRepository) GetUser(ctx context.Context, id uint) (*models.User, error) {

	rows, err := m.fetchUser(ctx, getUserQuery, id)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, nil
}

// CreateUser ...
func (m *userRepository) CreateUser(ctx context.Context, tx *sql.Tx, p *models.User) (int64, error) {

	stmt, err := tx.PrepareContext(ctx, createUserQuery)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx,
		p.Name,
		p.Email,
	)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// UpdateUser ...
func (m *userRepository) UpdateUser(ctx context.Context, tx *sql.Tx, p *models.User) (*models.User, error) {

	stmt, err := tx.PrepareContext(ctx, updateUserQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		p.Name,
		p.Email,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// DeleteUser ...
func (m *userRepository) DeleteUser(ctx context.Context, id uint) error {

	stmt, err := m.db.PrepareContext(ctx, deleteUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
