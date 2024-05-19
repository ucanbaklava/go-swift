package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ucanbaklava/go-auth/models"
)

type CommentRepository struct {
	DB *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	query := `INSERT INTO comments (post_id, content, author, created) VALUES (?, ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.PostId, comment.Content, comment.UserId, time.Now().String())
	return err
}

func (r *CommentRepository) GetByID(id int) (*models.Comment, error) {
	var comment models.Comment
	query := `SELECT * FROM comments WHERE id = ?`
	err := r.DB.Get(&comment, query, id)
	return &comment, err
}

func (r *CommentRepository) GetByPostID(postID int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT * FROM comments WHERE post_id = ?`
	err := r.DB.Select(&comments, query, postID)
	return comments, err
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	query := `UPDATE comments SET post_id = ?, content = ?, author = ?, created = ? WHERE id = ?`
	_, err := r.DB.Exec(query, comment.PostId, comment.Content, comment.ParentId, time.Now().String(), comment.Id)
	return err
}

func (r *CommentRepository) Delete(id int) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
