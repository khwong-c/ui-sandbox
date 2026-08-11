package session

import (
	"context"

	"github.com/khwong-c/dnd/backend/drivers/sql"
	"github.com/khwong-c/dnd/backend/features/session/repo"
	"github.com/khwong-c/dnd/backend/tooling/di"

	"github.com/google/uuid"
	"github.com/samber/do/v2"
	"github.com/samber/oops"
	"github.com/samber/ro"
	"gorm.io/gorm"
)

type Orchestrator struct {
	db       *gorm.DB
	repo     *repo.SessionRepo
	notifier ro.Subject[uuid.UUID]
}

func NewOrchestrator(i do.Injector) (*Orchestrator, error) {
	db := di.InvokeOrProvide(i, sql.NewInMemorySQLite)
	srepo := repo.NewSessionRepo(db)
	if err := srepo.Migrate(); err != nil {
		return nil, oops.Wrapf(err, "failed to migrate session repo")
	}
	return &Orchestrator{
		db:       db,
		repo:     srepo,
		notifier: ro.NewPublishSubject[uuid.UUID](),
	}, nil
}

func (o *Orchestrator) CreateSession(ctx context.Context) (uuid.UUID, error) {
	return o.repo.CreateSession(ctx)
}

func (o *Orchestrator) GetSession(ctx context.Context, id uuid.UUID) (*GameSession, error) {
	if exists, err := o.repo.IsSessionExist(ctx, id); err != nil || !exists {
		return nil, oops.Wrapf(err, "session not found")
	}
	return &GameSession{
		id:           id,
		db:           o.db,
		orchestrator: o,
	}, nil
}

func (o *Orchestrator) NotifySessionEvent(ctx context.Context, id uuid.UUID) error {
	o.notifier.NextWithContext(ctx, id)
	return nil
}

func (o *Orchestrator) GetNotifier() ro.Subject[uuid.UUID] {
	return o.notifier
}
