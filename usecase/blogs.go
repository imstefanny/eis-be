package usecase

import (
	"eis-be/repository"
)

type BlogsUsecase interface {
	GetAll() (interface{}, error)
}

type blogsUsecase struct {
	blogsRepository		repository.BlogsRepository
}

func NewBlogsUsecase(blogsRepo repository.BlogsRepository) *blogsUsecase {
	return &blogsUsecase{
		blogsRepository: blogsRepo,
	}
}

func (s *blogsUsecase) GetAll() (interface{}, error) {
	blogss, err := s.blogsRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return blogss, nil
}
