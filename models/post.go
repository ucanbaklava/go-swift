package models

import "time"

type Post struct {
	ID        int        `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Content   string     `json:"content" db:"content"`
	Excerpt   string     `json:"excerpt" db:"excerpt"`
	Slug      string     `json:"slug" db:"slug"`
	UserId    int        `json:"user_id" db:"user_id"`
	UpdatedAt string     `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	IsDeleted bool       `json:"is_deleted" db:"is_deleted"`
	Comments  []*Comment `json:"comments"`
}
