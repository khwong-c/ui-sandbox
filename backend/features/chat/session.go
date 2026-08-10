package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"github.com/samber/ro"
	"gorm.io/gorm"

	"github.com/khwong-c/dnd/backend/features/chat/repo"
)

type EventPair lo.Tuple3[uuid.UUID, repo.EventType, any]

func (p EventPair) Unpack() (uuid.UUID, repo.EventType, any) {
	return p.A, p.B, p.C
}

type GameSession struct {
	id           uuid.UUID
	db           *gorm.DB
	orchestrator *Orchestrator
}

func projectEntryToPair(e repo.SessionEvent, _ int) EventPair {
	return EventPair{e.ID, e.Type, e.Payload}
}

func (s *GameSession) GetSessionStream(ctx context.Context) (ro.Observable[EventPair], error) {
	// Retrieve all existing events
	r := repo.NewSessionRepo(s.db)
	eventEntries, err := r.GetSessionEvents(ctx, s.id)
	if err != nil {
		return nil, oops.Wrap(err)
	}

	prevEvents := lo.Map(eventEntries, projectEntryToPair)
	lastEventID := uuid.Nil

	liveStream := ro.Pipe6[
		uuid.UUID,
		uuid.UUID,
		[]EventPair,
		EventPair,
		EventPair,
		EventPair,
	](
		// Subscribe to the orchestrator's notifier, filter only events related to current session.
		s.orchestrator.GetNotifier(),
		ro.Filter(func(id uuid.UUID) bool {
			return id == s.id
		}),

		// Retrieve all new events from the storage.
		// Flatten the result into individual events.
		ro.MapErrWithContext[uuid.UUID, []EventPair](
			func(ctx context.Context, _ uuid.UUID) ([]EventPair, context.Context, error) {
				newEntries, err := r.GetSessionEventsAfter(ctx, s.id, lastEventID)
				return lo.Map(newEntries, projectEntryToPair), ctx, oops.Wrap(err)
			}),
		ro.Flatten[EventPair](),

		// Append events started before subscribing the stream.
		ro.StartWith(prevEvents...),

		// Update the last event ID for each event.
		ro.TapOnNext(func(pair EventPair) {
			id, _, _ := pair.Unpack()
			lastEventID = id
		}),

		// End the Stream when we encounter the End event.
		ro.TakeWhile[EventPair](func(pair EventPair) bool {
			_, event, _ := pair.Unpack()
			return event != repo.EventEnd
		}),
	)

	return liveStream, nil
}

func (s *GameSession) AddEvent(ctx context.Context, event repo.EventType, data any) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		r := repo.NewSessionRepo(tx)
		lastState, err := r.GetLatestEventType(ctx, s.id)
		if err != nil {
			return oops.Wrap(err)
		}
		// TODO: Validate Transitions
		_ = lastState
		return r.AddEvent(ctx, s.id, event, data)
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
