package usecase

import (
	"eis-be/dto"
	"eis-be/helpers"
	"eis-be/models"
	"eis-be/repository"
	"errors"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
)

type BlogsUsecase interface {
	Browse(page, limit int, search string) (interface{}, int64, error)
	Create(blog dto.CreateBlogsRequest, claims jwt.MapClaims) error
	Find(id int) (interface{}, error)
	Update(id int, blog dto.CreateBlogsRequest, claims jwt.MapClaims) (models.Blogs, error)
	Delete(id int) error
}

type blogsUsecase struct {
	blogsRepository repository.BlogsRepository
	usersRepository repository.UsersRepository
}

func NewBlogsUsecase(blogsRepo repository.BlogsRepository, usersRepo repository.UsersRepository) *blogsUsecase {
	return &blogsUsecase{
		blogsRepository: blogsRepo,
		usersRepository: usersRepo,
	}
}

func validateCreateBlogsRequest(req dto.CreateBlogsRequest) error {
	val := reflect.ValueOf(req)
	for i := 0; i < val.NumField(); i++ {
		if helpers.IsEmptyField(val.Field(i)) {
			return errors.New("field can't be empty")
		}
	}
	return nil
}

func (s *blogsUsecase) Browse(page, limit int, search string) (interface{}, int64, error) {
	blogs, total, err := s.blogsRepository.Browse(page, limit, search)

	if err != nil {
		return nil, total, err
	}

	return blogs, total, nil
}

func (s *blogsUsecase) Create(blog dto.CreateBlogsRequest, claims jwt.MapClaims) error {
	e := validateCreateBlogsRequest(blog)

	if e != nil {
		return e
	}

	blogData := models.Blogs{
		Title:     blog.Title,
		Content:   blog.Content,
		Thumbnail: blog.Thumbnail,
		CreatedBy: uint(claims["userId"].(float64)),
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

func (s *blogsUsecase) Update(id int, blog dto.CreateBlogsRequest, claims jwt.MapClaims) (models.Blogs, error) {
	blogData, err := s.blogsRepository.Find(id)

	if err != nil {
		return models.Blogs{}, err
	}

	blogData.Title = blog.Title
	blogData.Content = blog.Content
	blogData.Thumbnail = blog.Thumbnail
	blogData.UpdatedBy = uint(claims["userId"].(float64))

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

func (s *blogsUsecase) Delete(id int) error {
	err := s.blogsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
