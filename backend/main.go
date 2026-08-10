package main

import (
	"context"
	"syscall"

	"github.com/samber/do/v2"
	"github.com/samber/ro"

	"github.com/khwong-c/dnd/backend/drivers/sql"
	"github.com/khwong-c/dnd/backend/features/chat"
	"github.com/khwong-c/dnd/backend/features/chat/repo"
	"github.com/khwong-c/dnd/backend/server"
	"github.com/khwong-c/dnd/backend/tooling/di"
)

func main() {
	i := do.New()
	s := di.InvokeOrProvide(i, server.NewServer)

	go s.ListenAndServe()

	// Shutdown Gracefully
	_, _ = i.ShutdownOnSignals(
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL,
	)
}

func main1() {
	//orch := chat.NewOrchestrator()
	//uuid1, _ := uuid.NewV7()
	//uuid2, _ := uuid.NewV7()
	//orch.AddChat(uuid1, "Fugu Fish")
	//orch.AddChat(uuid2, "Fish 2")
	//orch.AddChat(uuid1, "End")
	//
	//sub1a := orch.SubscribeChat(uuid1).Subscribe(ro.PrintObserver[string]())
	//defer sub1a.Unsubscribe()
	//sub2 := orch.SubscribeChat(uuid2).Subscribe(ro.PrintObserver[string]())
	//defer sub2.Unsubscribe()
	//
	//orch.AddChat(uuid1, "Fish 1")
	//
	//sub1b := orch.SubscribeChat(uuid1).Subscribe(ro.PrintObserver[string]())
	//defer sub1b.Unsubscribe()
	//orch.AddChat(uuid1, "Fish 1, Second")
}

//func main2() {
//	ctx := context.Background()
//	db := di.InvokeOrProvide(nil, sql.NewInMemorySQLite)
//	sessRepo := repo.NewSessionRepo(db)
//	_ = sessRepo.Migrate()
//	sess, _ := sessRepo.CreateSession(ctx)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventChatFromGoblin, nil)
//	_ = sessRepo.AddEvent(ctx, sess, repo.EventEnd, nil)
//	r, _ := sessRepo.IsSessionEnded(ctx, sess)
//	fmt.Println(r)
//}

func main3() {
	ctx := context.Background()

	db := di.InvokeOrProvide(nil, sql.NewInMemorySQLite)
	r := repo.NewSessionRepo(db)
	_ = r.Migrate()

	orch := di.InvokeOrProvide(nil, chat.NewOrchestrator)
	id, _ := orch.CreateSession(ctx)
	sess, _ := orch.GetSession(ctx, id)
	strm, err := sess.GetSessionStream(ctx)
	if err != nil {
		panic(err)
	}

	strm.Subscribe(ro.PrintObserver[chat.EventPair]())
	sess.AddEvent(ctx, repo.EventChatFromGoblin, nil)
	sess.AddEvent(ctx, repo.EventChatFromGoblin, nil)
	//sess.AddEvent(ctx, repo.EventEnd, nil)

	strm2, err := sess.GetSessionStream(ctx)
	if err != nil {
		panic(err)
	}
	strm2.Subscribe(ro.PrintObserver[chat.EventPair]())
	sess.AddEvent(ctx, repo.EventAskTeam, nil)
	sess.AddEvent(ctx, repo.EventChatFromTeam, nil)
	sess.AddEvent(ctx, repo.EventEnd, nil)
}
