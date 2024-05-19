package models

type Comment struct {
	ID       int        `db:"id"                   json:"id"`
	UserID   int        `db:"user_id"              json:"userId"`
	PostID   int        `db:"post_id"              json:"postId"`
	Content  string     `db:"content"              json:"comment"`
	ParentID *int       `db:"parent_id"            json:"parentId"`
	Children []*Comment `json:"children,omitempty"`
}
