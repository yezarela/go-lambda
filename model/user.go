package model

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserScan struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Email     sql.NullString
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}

func (u *User) FromScan(s UserScan) *User {

	u.ID = uint(s.ID.Int64)
	u.Name = s.Name.String
	u.Email = s.Email.String
	u.CreatedAt = s.CreatedAt.Time
	u.UpdatedAt = s.UpdatedAt.Time

	return u
}
