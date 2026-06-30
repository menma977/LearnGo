package models

import "time"

type User struct {
	ID        int        `json:"-" db:"id"`
	Uuid      string     `json:"id" db:"uuid"`
	Name      string     `json:"name" db:"name"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"password"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
