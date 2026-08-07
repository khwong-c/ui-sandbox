package backup

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/ro"
)

type Orchestrator struct {
	storage  *chatStorage
	notifier ro.Subject[uuid.UUID]
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		storage:  createStorage(),
		notifier: ro.NewPublishSubject[uuid.UUID](),
	}
}

func (o *Orchestrator) SubscribeChat(session uuid.UUID) ro.Observable[string] {
	chatLogs, chatSize := o.storage.getFullChat(session)
	formatter := func(chat string, _ int) string {
		return fmt.Sprintf("Session: %s -> %s", session.String(), chat)
	}
	notTerminated := func(value string) bool {
		return !strings.HasSuffix(value, "-> End")
	}
	chatLogs = lo.Map(chatLogs, formatter)

	if chatSize > 0 && !lo.EveryBy(
		chatLogs, notTerminated,
	) {
		return ro.Just(chatLogs...)
	}

	chatStream := ro.Pipe3[
		uuid.UUID,
		uuid.UUID,
		string,
		string,
	](
		o.notifier,
		ro.Filter(func(input uuid.UUID) bool {
			return input == session
		}),
		ro.FlatMap(func(session uuid.UUID) ro.Observable[string] {
			latest, newSize := o.storage.getChatFrom(session, chatSize)
			chatSize = newSize
			return ro.Just[string](lo.Map(latest, formatter)...)
		}),
		ro.TakeWhile(notTerminated),
	)
	return ro.Pipe[string, string](
		chatStream,
		ro.StartWith(chatLogs...),
	)
}

func (o *Orchestrator) AddChat(session uuid.UUID, chat string) {
	o.storage.addChat(session, chat)
	o.notifier.Next(session)
}
