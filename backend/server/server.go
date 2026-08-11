package server

import (
	"net/http"

	"github.com/khwong-c/dnd/backend/features/session"
	"github.com/khwong-c/dnd/backend/tooling/di"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	dochi "github.com/samber/do/http/chi/v2"
	"github.com/samber/do/v2"
)

type Server struct {
	*http.Server
	orchestrator *session.Orchestrator
}

func NewServer(i do.Injector) (*Server, error) {
	r := chi.NewRouter()
	r.Use(middleware.NoCache)

	dochi.Use(r, "/debug/di", i)
	newServer := &Server{
		Server: &http.Server{
			Addr:    ":7086",
			Handler: r,
		},
		orchestrator: di.InvokeOrProvide(i, session.NewOrchestrator),
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})
	newServer.RegisterSessionEndpoints()

	return newServer, nil
}
