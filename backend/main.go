package main

import (
	"context"
	"dnd/backend/drivers/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/ro"

	chat "dnd/backend/features/chat/backup"
	"dnd/backend/features/chat/repo"
	"dnd/backend/tooling/di"
)

func main() {
	main1()
}

func main1() {
	orch := chat.NewOrchestrator()
	uuid1, _ := uuid.NewV7()
	uuid2, _ := uuid.NewV7()
	orch.AddChat(uuid1, "Fugu Fish")
	orch.AddChat(uuid2, "Fish 2")
	orch.AddChat(uuid1, "End")

	sub1a := orch.SubscribeChat(uuid1).Subscribe(ro.PrintObserver[string]())
	defer sub1a.Unsubscribe()
	sub2 := orch.SubscribeChat(uuid2).Subscribe(ro.PrintObserver[string]())
	defer sub2.Unsubscribe()

	orch.AddChat(uuid1, "Fish 1")

	sub1b := orch.SubscribeChat(uuid1).Subscribe(ro.PrintObserver[string]())
	defer sub1b.Unsubscribe()
	orch.AddChat(uuid1, "Fish 1, Second")
}

func main2() {
	ctx := context.Background()
	db := di.InvokeOrProvide(nil, sql.NewInMemorySQLite)
	sessRepo := repo.NewSessionRepo(db)
	_ = sessRepo.Migrate()
	sess, _ := sessRepo.CreateSession(ctx)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventChatFromGoblin, nil)
	_ = sessRepo.AddEvents(ctx, sess, repo.EventEnd, nil)
	r, _ := sessRepo.IsSessionEnded(ctx, sess)
	fmt.Println(r)
}
