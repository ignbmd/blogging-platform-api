package repositories

import (
	"context"
	"ignbmd/blogging-platform-api/config"
	"ignbmd/blogging-platform-api/internal/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		collection: config.GetDB().Collection("posts"),
	}
}

func (r *PostRepository) FindAll(ctx context.Context, searchTerm string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{}

	if searchTerm != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"title": bson.M{"$regex": searchTerm, "$options": "i"}},
				{"content": bson.M{"$regex": searchTerm, "$options": "i"}},
				{"category": bson.M{"$regex": searchTerm, "$options": "i"}},
			},
		}
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return []models.Post{}, err
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return []models.Post{}, err
	}

	if posts == nil {
		return []models.Post{}, nil
	}

	return posts, nil
}

func (r *PostRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Post, error) {
	var post models.Post
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(ctx context.Context, post *models.Post) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return r.collection.InsertOne(ctx, post)
}

func (r *PostRepository) Update(ctx context.Context, id primitive.ObjectID, post *models.Post) (*models.Post, error) {
	filterQuery := bson.M{"_id": id}
	updateQuery := bson.M{"$set": bson.M{
		"title":      post.Title,
		"content":    post.Content,
		"category":   post.Category,
		"tags":       post.Tags,
		"updated_at": time.Now(),
	}}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var updatedPost models.Post
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.collection.FindOneAndUpdate(ctx, filterQuery, updateQuery, opts).Decode(&updatedPost)
	if err != nil {
		return nil, err
	}

	return &updatedPost, nil
}

func (r *PostRepository) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	query := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return r.collection.DeleteOne(ctx, query)
}
