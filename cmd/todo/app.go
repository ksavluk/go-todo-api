package main

import (
	"context"
	"github.com/ksavluk/go-todo-api/pkg/plan"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ksavluk/go-todo-api/pkg/api/http"
	"github.com/ksavluk/go-todo-api/pkg/config"
	"github.com/ksavluk/go-todo-api/pkg/storage/postgres"
	"github.com/ksavluk/go-todo-api/pkg/task"
)

type server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type app struct {
	ctx        context.Context
	httpServer server
}

func newApp() *app {
	return &app{}
}

func (a *app) run(args []string) error {
	var ctxCancel context.CancelFunc
	a.ctx, ctxCancel = context.WithCancel(context.Background())

	defer func() {
		ctxCancel()

		if err := a.shutdown(); err != nil {
			log.Printf("[ERROR] app shutdown failed: %v\n", err)
		}
	}()

	cfg := config.LoadFromEnv()

	storage, err := postgres.NewStorage(cfg.DSN)
	if err != nil {
		return errors.Wrap(err, "init_storage")
	}

	planService := plan.NewService(storage)
	taskService := task.NewService(storage)
	httpServer := http.NewServer(cfg.Address, planService, taskService)

	go func() {
		if err := httpServer.Start(); err != nil {
			ctxCancel()

			log.Printf("[ERROR] http server has been stopped: %v\n", err)
		}
	}()

	if err := a.listenForTermination(); err != nil {
		return errors.Wrap(err, "app_has_been_terminated")
	}

	return nil
}

func (a *app) listenForTermination() error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(
		sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case <-sigCh:
		return errors.New("app has been terminated")
	case <-a.ctx.Done():
		return errors.Wrap(a.ctx.Err(), "context_done")
	}
}

func (a *app) shutdown() error {
	if err := a.httpServer.Shutdown(a.ctx); err != nil {
		return errors.Wrap(err, "http_server_shutdown")
	}

	return nil
}
