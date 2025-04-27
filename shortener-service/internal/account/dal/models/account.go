package models

import "time"

type AccountModel struct {
	ID        string
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
