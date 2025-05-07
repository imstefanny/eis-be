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
	Find(id int) (interface{}, error)
	Update(id int, blog dto.CreateBlogsRequest) (models.Blogs, error)
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
	blogs, err := s.blogsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (s *blogsUsecase) Create(blog dto.CreateBlogsRequest) error {
	e := validateCreateBlogsRequest(blog)
	
	if e != nil {
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

func (s *blogsUsecase) Find(id int) (interface{}, error) {
	blog, err := s.blogsRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogsUsecase) Update(id int, blog dto.CreateBlogsRequest) (models.Blogs, error) {
	blogData, err := s.blogsRepository.Find(id)

	if err != nil {
		return models.Blogs{}, err
	}

	blogData.Title = blog.Title
	blogData.Content = blog.Content
	blogData.Thumbnail = blog.Thumbnail
	
	e := s.blogsRepository.Update(id, blogData)

	if e != nil {
		return models.Blogs{}, e
	}

	blogUpdated, err := s.blogsRepository.Find(id)

	if err != nil {
		return models.Blogs{}, err
	}
	return blogUpdated, nil
}
