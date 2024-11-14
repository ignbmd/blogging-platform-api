package services

import (
	"context"
	"ignbmd/blogging-platform-api/internal/app/models"
	"ignbmd/blogging-platform-api/internal/app/repositories"
	"ignbmd/blogging-platform-api/internal/app/validators"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	repo      *repositories.PostRepository
	validator *validators.PostValidator
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{
		repo:      repo,
		validator: validators.NewPostValidator(),
	}
}

func (s *PostService) FindAll(ctx context.Context, searchTerm string) ([]models.Post, error) {
	return s.repo.FindAll(ctx, searchTerm)
}

func (s *PostService) FindByID(ctx context.Context, id string) (*models.Post, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, objectID)
}

func (s *PostService) Create(ctx context.Context, post *models.Post) error {

	if err := s.validator.Validate(post); err != nil {
		return err
	}

	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := s.repo.Create(ctx, post)
	return err
}

func (s *PostService) Update(ctx context.Context, id string, post *models.Post) (*models.Post, error) {
	if err := s.validator.Validate(post); err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, objectID, post)
}

func (s *PostService) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := s.repo.Delete(ctx, objectID)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
