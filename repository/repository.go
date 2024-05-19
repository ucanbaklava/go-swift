package repository

import "github.com/ucanbaklava/go-auth/models"

type IPostRepository interface {
	Create(post *models.Post) error
	GetByID(id int) (*models.Post, error)
	GetAll() ([]models.Post, error)
	Update(post *models.Post) error
	Delete(id int) error
}

type ICommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(id int) (*models.Comment, error)
	GetByPostID(postID int) ([]models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id int) error
}
