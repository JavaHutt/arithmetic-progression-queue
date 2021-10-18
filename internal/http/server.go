package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Server will perform operations over http.
type Server interface {
	// Open will setup a tcp listener and serve the http requests.
	Open() error

	// Close will close the socket if it's open.
	Close(ctx context.Context) error

	// Handler returns a http handler with all routes in place.
	Handler() http.Handler
}

type taskService interface {
	AddTask(task model.Task) error
	GetTasks() []model.TaskInfo
}

// Server represents an HTTP server.
type server struct {
	log         logrus.Logger
	serv        *http.Server
	encoder     *encoder
	port        string
	taskService taskService
}

func NewServer(log logrus.Logger, port string, taskService taskService) Server {
	return &server{
		log:         log,
		encoder:     newEncoder(),
		port:        port,
		taskService: taskService,
	}
}

// Open will setup a tcp listener and serve the http requests.
func (s *server) Open() error {
	s.serv = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.port),
		Handler: s.Handler(),
	}
	s.log.Info("server listening")
	return s.serv.ListenAndServe()
}

// Close will close the socket if it's open.
func (s *server) Close(ctx context.Context) error {
	if s.serv != nil {
		if err := s.serv.Shutdown(ctx); err != nil {
			return err
		}
		s.serv = nil
	}
	return nil
}

func (s *server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/tasks", newTasksHandler(s.encoder, s.taskService).Routes)
	})

	return r
}
