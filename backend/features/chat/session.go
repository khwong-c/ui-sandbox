package chat

import (
	"context"

	"github.com/samber/lo"
	"github.com/samber/ro"

	"dnd/backend/features/chat/repo"
)

type EventPair lo.Tuple2[repo.EventType, any]

type GameSessionI interface {
	GetSessionStream(ctx context.Context) ro.Observable[EventPair]
	AddEvent(ctx context.Context, event repo.EventType, data any)
	IsEnded(ctx context.Context) bool
}
