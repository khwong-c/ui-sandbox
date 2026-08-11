package repo

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/samber/oops"
	"gorm.io/gorm"

	"github.com/khwong-c/dnd/backend/tooling"
)

type SessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (r *SessionRepo) Migrate() error {
	return r.db.AutoMigrate(
		&Session{},
		&SessionEvent{},
	)
}

func (r *SessionRepo) CreateSession(ctx context.Context) (uuid.UUID, error) {
	newSession := Session{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := gorm.G[Session](r.db).Create(ctx, &newSession); err != nil {
			return oops.Wrapf(err, "failed to create session")
		}
		sessID := newSession.ID
		err := r.AddEvent(ctx, sessID, EventChatInit, nil)
		return oops.Wrapf(err, "failed to add init event")
	})
	return newSession.ID, err
}

func (r *SessionRepo) AddEvent(ctx context.Context, sess uuid.UUID, eventType EventType, payload any) error {
	newEvent := SessionEvent{
		SessionID:    sess,
		Type:         eventType,
		Payload:      tooling.Must(json.Marshal(payload)),
		EndOfSession: eventType == EventEnd,
	}
	return gorm.G[SessionEvent](r.db).Create(ctx, &newEvent)
}

func (r *SessionRepo) GetSessionEvents(ctx context.Context, sess uuid.UUID) ([]SessionEvent, error) {
	session, err := gorm.G[Session](r.db).
		Preload("Events", nil).
		Where(sess).
		First(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to get session events")
	}
	return session.Events, nil
}

func (r *SessionRepo) GetSessionEventsAfter(ctx context.Context, sess uuid.UUID, lastEvent uuid.UUID) ([]SessionEvent, error) {
	events, err := gorm.G[SessionEvent](r.db).
		Where("session_id = ? AND id > ?", sess, lastEvent).
		Find(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to poll events")
	}
	return events, nil
}

func (r *SessionRepo) IsSessionEnded(ctx context.Context, sess uuid.UUID) (bool, error) {
	cnt, err := gorm.G[SessionEvent](r.db).
		Where(SessionEvent{
			SessionID:    sess,
			EndOfSession: true,
		}).
		Count(ctx, "*")
	if err != nil {
		return false, oops.Wrapf(err, "failed to Count")
	}
	return cnt > 0, nil
}

func (r *SessionRepo) IsSessionExist(ctx context.Context, sess uuid.UUID) (bool, error) {
	cnt, err := gorm.G[Session](r.db).
		Where(sess).
		Count(ctx, "*")
	if err != nil {
		return false, oops.Wrapf(err, "failed to Count")
	}
	return cnt > 0, nil
}

func (r *SessionRepo) GetLatestEventType(ctx context.Context, sess uuid.UUID) (EventType, error) {
	events, err := gorm.G[SessionEvent](r.db).
		Where(SessionEvent{
			SessionID: sess,
		}).
		Select("Type").
		Order("id DESC").
		Limit(1).
		Find(ctx)
	if err != nil || len(events) == 0 {
		return "", oops.Join(
			oops.New("session events not found"),
			oops.Wrapf(err, "Failed to query Data Storage"),
		)
	}
	return events[0].Type, nil
}
