package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"dnd/backend/drivers/sql"
)

type Session struct {
	ID     uuid.UUID      `gorm:"primaryKey;autoIncrement:false"`
	Events []SessionEvent `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE"`

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SessionEvent struct {
	ID           uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	SessionID    uuid.UUID `gorm:"index:session_with_end,priority:1"`
	EndOfSession bool      `gorm:"index:session_with_end,priority:2"`
	Session      Session

	Type    EventType
	Payload sql.JSONType

	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (s *Session) BeforeCreate(*gorm.DB) (err error) {
	s.ID, err = uuid.NewV7()
	return
}

func (e *SessionEvent) BeforeCreate(*gorm.DB) (err error) {
	e.ID, err = uuid.NewV7()
	return
}
