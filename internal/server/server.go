package server

import (
	"p2p/config"
	"p2p/internal/router"
	"p2p/pkg/logers"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Addr   string
	logger *logers.Logger
	notify chan error
	r      *fiber.App
}

func NewServer(cfg *config.Config) *Server {

	server := &Server{
		Addr:   cfg.Port,
		logger: &logers.Logger{},
		notify: make(chan error, 1),
		r:      fiber.New(),
	}
	return server
}

func (s *Server) Start() error {

	s.r = fiber.New(fiber.Config{
		Concurrency: 20,
	})
	router.Setup(s.r)
	go func() {
		s.notify <- s.r.Listen(s.Addr)
		s.r.Shutdown()
		close(s.notify)
	}()

	return nil
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	return s.r.Shutdown()
}
