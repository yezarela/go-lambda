package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListUserParams struct {
	SortBy        string
	SortDirection string
	Limit         uint
	Offset        uint
}

type UserRepository interface {
	ListUser(ctx context.Context, p ListUserParams) ([]*User, error)
	GetUser(ctx context.Context, id uint) (*User, error)
	CreateUser(ctx context.Context, p *User) (int64, error)
	UpdateUser(ctx context.Context, p *User) (*User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserUsecase interface {
	ListUser(ctx context.Context, p ...ListUserParams) ([]*User, error)
	GetByID(ctx context.Context, id uint) (*User, error)
	CreateUser(ctx context.Context, p *User) (*User, error)
	UpdateUser(ctx context.Context, p *User) (*User, error)
	DeleteUser(ctx context.Context, id uint) error
}

