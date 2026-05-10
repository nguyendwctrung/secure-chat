package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	SenderID uuid.UUID `gorm:"type:uuid"`
	ReceiverID uuid.UUID `gorm:"type:uuid"`

	Content string `gorm:"type:text"`

	CreatedAt time.Time
}