package repository

import (
	"database/sql"
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
	_, err := r.DB.Exec(query, post.Title, post.Content, post.UserId, post.CreatedAt, post.Excerpt, post.Slug)
	return err
}

func (r *PostRepository) GetByID(id int) (*models.Post, error) {
	post := models.Post{}
	query := `SELECT * FROM posts WHERE id = ?`
	err := r.DB.Get(&post, query, id)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	comments, err := fetchComments(r.DB, id)
	post.Comments = buildCommentTree(comments)
	return &post, err
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT * FROM posts`
	err := r.DB.Select(&posts, query)
	return posts, err
}

func (r *PostRepository) Update(post *models.Post) error {
	query := `UPDATE posts SET title = ?, content = ?, user_id = ?, created_at = ?, excerpt = ?, slug = ? WHERE id = ?`
	_, err := r.DB.Exec(query, post.Title, post.Content, post.UserId, post.CreatedAt, post.Excerpt, post.Slug, post.ID)
	return err
}

func (r *PostRepository) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func fetchComments(db *sqlx.DB, postID int) ([]*models.Comment, error) {
	rows, err := db.Query("SELECT id, content, user_id, post_id, parent_id FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		var parentID sql.NullInt64
		if err := rows.Scan(&comment.Id, &comment.Content, &comment.UserId, &comment.PostId, &parentID); err != nil {
			return nil, err
		}
		if parentID.Valid {
			comment.ParentId = new(int)
			*comment.ParentId = int(parentID.Int64)
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func buildCommentTree(comments []*models.Comment) []*models.Comment {
	commentMap := make(map[int]*models.Comment)
	var roots []*models.Comment

	for _, comment := range comments {
		commentMap[comment.Id] = comment
		if comment.ParentId == nil {
			roots = append(roots, comment)
		}
	}

	for _, comment := range comments {
		if comment.ParentId != nil {
			parent := commentMap[*comment.ParentId]
			parent.Children = append(parent.Children, comment)
		}
	}

	return roots
}
