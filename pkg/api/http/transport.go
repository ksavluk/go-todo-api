package http

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type server struct {
	echo    *echo.Echo
	address string
}

func NewServer(address string, planService PlanService, taskService TaskService) *server {
	return &server{
		echo:    newHandlers(planService, taskService),
		address: address,
	}
}

func (s *server) Start() error {
	return s.echo.Start(s.address)
}

func (s *server) Shutdown(ctx context.Context) error {
	ctx, ctxCancel := context.WithTimeout(ctx, 5*time.Second)
	defer ctxCancel()

	return s.echo.Shutdown(ctx)
}
