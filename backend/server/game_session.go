package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/samber/oops"
	"github.com/samber/ro"

	"github.com/khwong-c/dnd/backend/features/session"
	"github.com/khwong-c/dnd/backend/features/session/repo"
)

func (s *Server) HandleNewSession(w http.ResponseWriter, r *http.Request) {
	sid, _ := s.orchestrator.CreateSession(r.Context())
	payload := map[string]any{
		"id": sid.String(),
	}
	b, _ := json.Marshal(payload)
	w.Write(b)
	w.Write([]byte("\n"))
}

func (s *Server) HandleGetGameSession(w http.ResponseWriter, r *http.Request) {
	// Get Session ID from URL
	sessionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sess, err := s.orchestrator.GetSession(ctx, sessionID)
	if err != nil {
		errO, _ := oops.AsOops(err)
		http.Error(w, errO.Public(), http.StatusNotFound)
		return
	}

	// TODO: Error Handling
	stm, _ := sess.GetSessionStream(ctx)

	// Set headers to mimic SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}
	flusher.Flush()

	observer := ro.NewObserverWithContext(
		func(ctx context.Context, item session.EventPair) {
			id, event, blob := item.Unpack()
			var data map[string]any
			json.Unmarshal(blob, &data)
			payload := map[string]any{
				"id":      id.String(),
				"event":   event,
				"payload": data,
			}
			b, _ := json.Marshal(payload)
			w.Write([]byte("event: msg\n"))
			w.Write([]byte("data: "))
			w.Write(b)
			w.Write([]byte("\n\n"))
			flusher.Flush()
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte("event: close\n\n"))
			flusher.Flush()
			//cancel()
		},
		func(ctx context.Context, err error) {
			cancel()
		},
		func(ctx context.Context) {
			w.Write([]byte("event: close\n\n"))
			flusher.Flush()
			//cancel()
		},
	)

	sub := stm.SubscribeWithContext(r.Context(), observer)
	defer sub.Unsubscribe()
	<-ctx.Done()
	fmt.Println("done")
}

func (s *Server) HandlePostNewChat(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	sess, err := s.orchestrator.GetSession(ctx, sessionID)
	if err != nil {
		errO, _ := oops.AsOops(err)
		http.Error(w, errO.Public(), http.StatusNotFound)
		return
	}
	_ = sess.AddEvent(ctx, repo.EventChatToGoblin, map[string]any{
		"message": "Hello, Goblin!",
	})
}

func (s *Server) HandleGetUI(w http.ResponseWriter, r *http.Request) {
	sessionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}
	page := fmt.Sprintf(`
// client.js
<html><head><script>
const maxReconnectTries = 3

let reconnectAttempts = 0
const sse = new EventSource('/session/%s')
sse.onmessage = m => {
  const { type, data } = JSON.parse(m.data)
  if (type === 'close') sse.close()
  else console.log(data)
}
sse.onerror = () => {
console.log("err")
  if (reconnectAttempts > maxReconnectTries) {
    sse.close()
    alert("We have a baaad network error!")
  } else {
    reconnectAttempts++
  }
}
</script></head><body></body></html>
`, sessionID.String())
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(page))
}

func (s *Server) RegisterSessionEndpoints() {
	r := s.Server.Handler.(*chi.Mux)
	r.Post("/session", s.HandleNewSession)
	r.Get("/session/{id}", s.HandleGetGameSession)
	r.Post("/session/{id}/chat", s.HandlePostNewChat)
	r.Get("/session/{id}/ui", s.HandleGetUI)
}
