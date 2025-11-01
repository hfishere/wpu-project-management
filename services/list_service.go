package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/repositories"
	"github.com/hfishere/wpu-project-management/utils"
)

type ListWithOrder struct {
	Positions []uuid.UUID
	Lists     []models.List
}

type ListService interface {
	GetByBoardID(boardPublicID string) (*ListWithOrder, error)
	GetByID(id uint) (*models.List, error)
	GetByPublicID(id string) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, positions []uuid.UUID)
}

type listService struct {
	listRepo         repositories.ListRepository
	boardRepo        repositories.BoardRepository
	listPositionRepo repositories.ListPositionRepository
}

func NewListService(
	listRepo repositories.ListRepository,
	boardRepo repositories.BoardRepository,
	listPositionRepo repositories.ListPositionRepository) ListService {
	return &listService{listRepo, boardRepo, listPositionRepo}
}

func (s *listService) Create(list *models.List) error {
	panic("unimplemented")
}

func (s *listService) Delete(id uint) error {
	panic("unimplemented")
}

func (s *listService) GetByBoardID(boardPublicID string) (*ListWithOrder, error) {
	if _, err := s.boardRepo.FindByPublicID(boardPublicID); err != nil {
		return nil, errors.New("board tidak ditemukan")
	}

	positions, err := s.listPositionRepo.GetListOrder(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list order :" + err.Error())
	}

	lists, err := s.listRepo.FindByBoardID(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list :" + err.Error())
	}

	// sorting by position
	orderedList := utils.SortingListByPosition(lists, positions)

	return &ListWithOrder{
		Positions: positions,
		Lists:     orderedList,
	}, nil
}

func (s *listService) GetByID(id uint) (*models.List, error) {
	panic("unimplemented")
}

func (s *listService) GetByPublicID(id string) (*models.List, error) {
	panic("unimplemented")
}

func (s *listService) Update(list *models.List) error {
	panic("unimplemented")
}

func (s *listService) UpdatePosition(boardPublicID string, positions []uuid.UUID) {
	panic("unimplemented")
}
