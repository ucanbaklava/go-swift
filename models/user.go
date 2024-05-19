package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	LastLogin time.Time `json:"lastLogin"`
}
