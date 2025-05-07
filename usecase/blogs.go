package usecase

import (
	"eis-be/models"
	"eis-be/repository"
	"eis-be/dto"
	"eis-be/helpers"
	"reflect"
	"errors"
)

type BlogsUsecase interface {
	GetAll() (interface{}, error)
	Create(blog dto.CreateBlogsRequest) error
}

type blogsUsecase struct {
	blogsRepository		repository.BlogsRepository
}

func NewBlogsUsecase(blogsRepo repository.BlogsRepository) *blogsUsecase {
	return &blogsUsecase{
		blogsRepository: blogsRepo,
	}
}

func validateCreateBlogsRequest(req dto.CreateBlogsRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("Field can't be empty")
		}
	}
	return nil
}

func (s *blogsUsecase) GetAll() (interface{}, error) {
	blogss, err := s.blogsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return blogss, nil
}

func (s *blogsUsecase) Create(blog dto.CreateBlogsRequest) error {
	e := validateCreateBlogsRequest(blog)
	
	if e!= nil {
		return e
	}

	blogData := models.Blogs{
		Active: true,
		Title: blog.Title,
		Content:  blog.Content,
		Thumbnail: blog.Thumbnail,
	}

	err := s.blogsRepository.Create(blogData)

	if err != nil {
		return err
	}

	return nil
}
