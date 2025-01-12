package models

import (
	"errors"
	"strings"
	"time"
)

type Role string

const (
	Admin Role = "admin"
	Guru  Role = "guru"
)

func (r Role) IsValid() bool {
	return r == Admin || r == Guru
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Password  string    `json:"-" gorm:"not null;size:50"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null;size:50"`
	Role      Role      `json:"role" gorm:"type:ENUM('admin','guru');default:guru"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if !u.Role.IsValid() {
		return errors.New("invalid role")
	}
	return nil
}
