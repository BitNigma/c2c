package app

import (
	"fmt"
	"os"
	"os/signal"
	"p2p/config"
	"p2p/internal/server"
	"p2p/pkg/logers"
	"syscall"
)

var err error

func Run(cfg *config.Config) error {

	l := logers.New(cfg.Level)

	//Server
	server := server.NewServer(cfg)
	err = server.Start()
	if err != nil {
		l.Fatal("can't start server", err)
		return err
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	return nil
}
