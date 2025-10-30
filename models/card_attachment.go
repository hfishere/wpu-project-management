package models

import (
	"time"

	"github.com/google/uuid"
)

type CardAttachment struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id"`
	File       string    `json:"file" db:"file"`
	UserID     int64     `json:"user_internal_id" db:"user_internal_id"`
	CardID     int64     `json:"card_internal_id" db:"card_internal_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
