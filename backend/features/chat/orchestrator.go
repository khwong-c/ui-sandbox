package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/ro"
)

type Orchestrator struct {
	notifier ro.Subject[uuid.UUID]
}

type OrchestratorI interface {
	CreateSession(ctx context.Context) uuid.UUID
	GetSession(ctx context.Context, sessionID uuid.UUID) GameSessionI
	NotifySessionEvent(ctx context.Context, sessionID uuid.UUID)
}
