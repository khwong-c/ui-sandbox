package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"github.com/samber/ro"
	"gorm.io/gorm"

	"dnd/backend/features/chat/repo"
)

type EventPair lo.Tuple2[repo.EventType, any]

type GameSession struct {
	id           uuid.UUID
	lastEventID  uuid.UUID // Managed by Operators in GetSessionStream
	db           *gorm.DB
	orchestrator *Orchestrator
}

func (s *GameSession) GetSessionStream(ctx context.Context) (ro.Observable[EventPair], error) {
	//TODO implement me
	panic("implement me")
}

func (s *GameSession) AddEvent(ctx context.Context, event repo.EventType, data any) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		//TODO implement me
		panic("implement me")
		return nil
	}); err != nil {
		return oops.Wrapf(
			err,
			"Failed to add Event",
		)
	}
	return oops.Wrapf(
		s.orchestrator.NotifySessionEvent(ctx, s.id),
		"Notification Error",
	)
}

func (s *GameSession) IsEnded(ctx context.Context) (bool, error) {
	return repo.NewSessionRepo(s.db).IsSessionEnded(ctx, s.id)
}
