package behaviourAuthor

import (
	"TaskManagement/app/helpers"
	repoAuthor "TaskManagement/app/repository/author"
	"context"
)

type AuthorRepository struct {
	ctx context.Context
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{
		ctx: context.Background(),
	}
}

func (r *AuthorRepository) GetAllAuthor() ([]repoAuthor.Author, int, error) {
	resSet, cnt, err := repoAuthor.GetAllAuthor(r.ctx)
	if err != nil {
		return nil, 0, err
	}
	return resSet, cnt, nil
}

func (r *AuthorRepository) GetAuthor(id int) (*repoAuthor.Author, error) {
	resSet, err := repoAuthor.GetAuthor(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return resSet, nil
}

func (r *AuthorRepository) AddAuthor(obj repoAuthor.Author) (interface{}, error) {
	resSet, err := repoAuthor.AddAuthor(r.ctx, obj)
	if err != nil {
		return nil, err
	}
	return resSet, nil
}

func (r *AuthorRepository) UpdateAuthor(obj map[string]interface{}) error {
	helpers.RemoveNilKeys(obj)
	err := repoAuthor.UpdateAuthor(r.ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthorRepository) DeleteAuthor(id int) error {
	err := repoAuthor.DeleteAuthor(r.ctx, id)
	if err != nil {
		return err
	}
	return nil

}
