package models

import "time"

type Post struct {
	ID        int        `db:"id"         json:"id"`
	Title     string     `db:"title"      json:"title"`
	Content   string     `db:"content"    json:"content"`
	Excerpt   string     `db:"excerpt"    json:"excerpt"`
	Slug      string     `db:"slug"       json:"slug"`
	UserID    int        `db:"user_id"    json:"userId"`
	UpdatedAt string     `db:"updated_at" json:"updatedAt"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
	Comments  []*Comment `json:"comments"`
}
