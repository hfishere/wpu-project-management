package utils

import (
	"github.com/google/uuid"
	"github.com/hfishere/wpu-project-management/models"
)

func SortingListByPosition(lists []models.List, order []uuid.UUID) []models.List {
	ordered := make([]models.List, 0, len(order))

	listMap := make(map[uuid.UUID]models.List)
	for _, l := range lists {
		listMap[l.PublicID] = l
	}

	// Urutkan sesuai order
	for _, id := range order {
		if list, ok := listMap[id]; ok {
			ordered = append(ordered, list)
		}
	}

	return ordered
}
