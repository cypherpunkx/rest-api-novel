package service

import (
	"enigmacamp.com/rest-api-novel/domain"
	"enigmacamp.com/rest-api-novel/repository"
	"enigmacamp.com/rest-api-novel/utils/exception"
)

type NovelService interface {
	GetAllNovelList(page, perPage int, code, title, publisher, year, author string) ([]domain.Novel, int, error)
	GetNovelById(id string) (domain.Novel, error)
	CreateNovel(payload domain.Novel) error
	UpdateNovelByID(id string, payload domain.Novel) (*domain.Novel, error)
	DeleteNovelByID(id string) error
}
type novelService struct {
	repository repository.NovelRepository
}

func (n *novelService) GetAllNovelList(page, perPage int, code, title, publisher, year, author string) ([]domain.Novel, int, error) {
	novels := []domain.Novel{}

	totalData, err := n.repository.Count(code, title, publisher, year, author)
	if err != nil {
		return novels, 0, err
	}

	novels, err = n.repository.ListPaginate(page, perPage, code, title, publisher, year, author)
	if err != nil {
		return novels, 0, err
	}

	return novels, totalData, nil
}

func (n *novelService) GetNovelById(id string) (domain.Novel, error) {
	novel, err := n.repository.Get(id)
	if err != nil {
		return novel, err
	}

	return novel, nil
}

func (n *novelService) CreateNovel(payload domain.Novel) error {
	novels, err := n.repository.List()
	if err != nil {
		return exception.ErrFailedCreate
	}

	for _, novel := range novels {
		if novel.Code == payload.Code {
			return exception.ErrCodeAlreadyExist
		}

		if novel.Title == payload.Title {
			return exception.ErrTitleAlreadyExist
		}
	}

	err = n.repository.Create(payload)
	if err != nil {
		return exception.ErrFailedCreate
	}

	return nil
}

func (n *novelService) UpdateNovelByID(id string, payload domain.Novel) (*domain.Novel, error) {
	novel, err := n.repository.Get(id)
	if err != nil {
		return &novel, err
	}

	novels, err := n.repository.List()
	if err != nil {
		return &novel, exception.ErrFailedCreate
	}

	for _, novel := range novels {
		if novel.Code == payload.Code {
			return &novel, exception.ErrCodeAlreadyExist
		}

		if novel.Title == payload.Title {
			return &novel, exception.ErrTitleAlreadyExist
		}
	}

	data, err := n.repository.Update(id, payload)
	if err != nil {
		return &novel, exception.ErrFailedUpdate
	}

	return data, nil
}

func (n *novelService) DeleteNovelByID(id string) error {
	novel, err := n.repository.Get(id)
	if err != nil {
		return err
	}

	err = n.repository.Delete(novel.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewNovelService(novelRepo repository.NovelRepository) NovelService {
	return &novelService{
		repository: novelRepo,
	}
}
