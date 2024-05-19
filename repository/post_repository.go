package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ucanbaklava/go-auth/models"
)

type PostRepository struct {
	DB *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	query := `INSERT INTO posts (title, content, user_id, created_at, excerpt, slug) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, post.Title, post.Content, post.UserID, post.CreatedAt, post.Excerpt, post.Slug)

	return fmt.Errorf("error at Post.Create %w", err)
}

func (r *PostRepository) GetByID(id int) (*models.Post, error) {
	post := models.Post{}
	query := `SELECT * FROM posts WHERE id = ?`
	err := r.DB.Get(&post, query, id)
	if err != nil {
		log.Fatalln(err.Error())

		return nil, fmt.Errorf("error at Post.GetByID %w", err)
	}

	comments, err := fetchComments(r.DB, id)
	post.Comments = buildCommentTree(comments)

	return &post, err
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post

	query := `SELECT * FROM posts`
	err := r.DB.Select(&posts, query)

	return posts, fmt.Errorf("error at Post.GetAll %w", err)
}

func (r *PostRepository) Update(post *models.Post) error {
	query := `UPDATE posts SET title = ?, content = ?, user_id = ?, created_at = ?, excerpt = ?, slug = ? WHERE id = ?`
	_, err := r.DB.Exec(query, post.Title, post.Content, post.UserID, post.CreatedAt, post.Excerpt, post.Slug, post.ID)

	return fmt.Errorf("error at Post.Update %w", err)
}

func (r *PostRepository) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := r.DB.Exec(query, id)

	return fmt.Errorf("error at Post.Delete %w", err)
}

func fetchComments(db *sqlx.DB, postID int) ([]*models.Comment, error) {
	rows, err := db.Query("SELECT id, content, user_id, post_id, parent_id FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, fmt.Errorf("error at Post.fetchComments %w", err)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error at Post.fetchComments %w", rows.Err())
	}

	defer rows.Close()

	var comments []*models.Comment

	for rows.Next() {
		var comment models.Comment

		var parentID sql.NullInt64

		if err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID, &parentID); err != nil {
			return nil, fmt.Errorf("error at Post.fetchComments rows.Scan %w", err)
		}

		if parentID.Valid {
			comment.ParentID = new(int)
			*comment.ParentID = int(parentID.Int64)
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func buildCommentTree(comments []*models.Comment) []*models.Comment {
	commentMap := make(map[int]*models.Comment)

	var roots []*models.Comment

	for _, comment := range comments {
		commentMap[comment.ID] = comment

		if comment.ParentID == nil {
			roots = append(roots, comment)
		}
	}

	for _, comment := range comments {
		if comment.ParentID != nil {
			parent := commentMap[*comment.ParentID]
			parent.Children = append(parent.Children, comment)
		}
	}

	return roots
}
