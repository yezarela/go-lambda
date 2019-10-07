package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/yezarela/go-lambda/models"
)

// Repository ...
type Repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (m *Repository) fetchUser(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	op := "user.Repository.fetchUser"

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, op)
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
			return nil, errors.Wrap(err, op)
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
func (m *Repository) ListUser(ctx context.Context, param ListUserParams) ([]*models.User, error) {
	op := "user.Repository.ListUser"

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
		return nil, errors.Wrap(err, op)
	}

	return items, nil
}

// GetUser ...
func (m *Repository) GetUser(ctx context.Context, id uint) (*models.User, error) {
	op := "user.Repository.GetUser"

	rows, err := m.fetchUser(ctx, getUserQuery, id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, nil
}

// CreateUser ...
func (m *Repository) CreateUser(ctx context.Context, tx *sql.Tx, p *models.User) (int64, error) {
	op := "user.Repository.CreateUser"

	stmt, err := tx.PrepareContext(ctx, createUserQuery)
	if err != nil {
		return -1, errors.Wrap(err, op)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx,
		p.Name,
		p.Email,
	)
	if err != nil {
		return -1, errors.Wrap(err, op)
	}

	return res.LastInsertId()
}

// UpdateUser ...
func (m *Repository) UpdateUser(ctx context.Context, tx *sql.Tx, p *models.User) (*models.User, error) {
	op := "user.Repository.UpdateUser"

	stmt, err := tx.PrepareContext(ctx, updateUserQuery)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		p.Name,
		p.Email,
	)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return p, nil
}

// DeleteUser ...
func (m *Repository) DeleteUser(ctx context.Context, id uint) error {
	op := "user.Repository.DeleteUser"

	stmt, err := m.db.PrepareContext(ctx, deleteUserQuery)
	if err != nil {
		return errors.Wrap(err, op)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
