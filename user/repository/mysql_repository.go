package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/yezarela/go-lambda/domain"
)

type repository struct {
	db *sql.DB
}

// NewMysqlRepository ...
func NewMysqlRepository(db *sql.DB) domain.UserRepository {
	return &repository{db}
}

type UserScan struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Email     sql.NullString
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}

func FromScan(s UserScan) *domain.User {
	u := domain.User{}

	u.ID = uint(s.ID.Int64)
	u.Name = s.Name.String
	u.Email = s.Email.String
	u.CreatedAt = s.CreatedAt.Time
	u.UpdatedAt = s.UpdatedAt.Time

	return &u
}

func (m *repository) fetchUser(ctx context.Context, query string, args ...interface{}) ([]*domain.User, error) {
	op := "user.Repository.fetchUser"

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer rows.Close()

	items := []*domain.User{}
	for rows.Next() {
		s := UserScan{}

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

		data := FromScan(s)

		items = append(items, data)
	}

	return items, nil
}

// ListUser ...
func (m *repository) ListUser(ctx context.Context, param domain.ListUserParams) ([]*domain.User, error) {
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

	query := fmt.Sprintf(ListUserQuery,
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
func (m *repository) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	op := "user.Repository.GetUser"

	rows, err := m.fetchUser(ctx, GetUserQuery, id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, nil
}

// CreateUser ...
func (m *repository) CreateUser(ctx context.Context, p *domain.User) (int64, error) {
	op := "user.Repository.CreateUser"

	stmt, err := m.db.PrepareContext(ctx, CreateUserQuery)
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
func (m *repository) UpdateUser(ctx context.Context, p *domain.User) (*domain.User, error) {
	op := "user.Repository.UpdateUser"

	stmt, err := m.db.PrepareContext(ctx, UpdateUserQuery)
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
func (m *repository) DeleteUser(ctx context.Context, id uint) error {
	op := "user.Repository.DeleteUser"

	stmt, err := m.db.PrepareContext(ctx, DeleteUserQuery)
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
