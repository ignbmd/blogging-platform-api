package validators

import (
	"fmt"
	"ignbmd/blogging-platform-api/internal/app/models"
)

type PostValidator struct {
	Errors map[string]string
}

func NewPostValidator() *PostValidator {
	return &PostValidator{
		Errors: make(map[string]string),
	}
}

func (v *PostValidator) Validate(post *models.Post) error {
	v.Errors = make(map[string]string)

	if len(post.Title) < 3 || len(post.Title) > 100 {
		v.Errors["title"] = "Title must be between 3 and 100 characters"
	}

	if len(post.Content) < 10 {
		v.Errors["content"] = "Content must be at least 10 characters"
	}

	if len(post.Category) < 2 || len(post.Category) > 50 {
		v.Errors["category"] = "Category must be between 2 and 50 characters"
	}

	if post.Tags != nil {
		for i, tag := range post.Tags {
			if len(tag) < 2 || len(tag) > 20 {
				v.Errors[fmt.Sprintf("tags[%d]", i)] = "Each tag must be between 2 and 20 characters"
			}
		}
	}

	if len(v.Errors) > 0 {
		return NewValidationError(v.Errors)
	}

	return nil
}
