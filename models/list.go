package models

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	InternalID      int64     `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID        uuid.UUID `json:"public_id" db:"public_id"`
	Title           string    `json:"title" db:"title"`
	Position        int64     `json:"position" db:"position"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	BoardPublicID   uuid.UUID `json:"board_public_id" db:"board_public_id" gorm:"column:board_public_id"`
	BoardInternalID int64     `json:"-" db:"board_internal_id"`
}
