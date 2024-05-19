package models

type Comment struct {
	Id       int        `json:"id" db:"id"`
	UserId   int        `json:"user_id" db:"user_id"`
	PostId   int        `json:"post_id" db:"post_id"`
	Content  string     `json:"comment" db:"content"`
	ParentId *int       `json:"parent_id" db:"parent_id"`
	Children []*Comment `json:"children,omitempty"`
}
