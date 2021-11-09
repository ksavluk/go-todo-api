package http

import (
	"context"
	"github.com/ksavluk/go-todo-api/pkg/plan"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ksavluk/go-todo-api/pkg/task"
)

type PlanService interface {
	GetAll(ctx context.Context) ([]plan.Plan, error)
	Get(ctx context.Context, planID uint) (*plan.Plan, error)
	Create(ctx context.Context, plan plan.Plan) (*plan.Plan, error)
	Update(ctx context.Context, plan plan.Plan) (*plan.Plan, error)
	Delete(ctx context.Context, planID uint) error
}

type TaskService interface {
	GetAll(ctx context.Context, planID uint) ([]task.Task, error)
	Get(ctx context.Context, planID, taskID uint) (*task.Task, error)
	Create(ctx context.Context, planID uint, task task.Task) (*task.Task, error)
	Update(ctx context.Context, planID uint, task task.Task) (*task.Task, error)
	Delete(ctx context.Context, taskID uint) error
	Done(ctx context.Context, planID, taskID uint) (*task.Task, error)
	Undo(ctx context.Context, planID, taskID uint) (*task.Task, error)
}

type handlers struct {
	echo        *echo.Echo
	planService PlanService
	taskService TaskService
}

func newHandlers(planService PlanService, taskService TaskService) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	r := &handlers{
		echo:        e,
		planService: planService,
		taskService: taskService,
	}

	r.initRoutes()

	return e
}

func (h *handlers) initRoutes() {
	api := h.echo.Group("/api/v1")
	api.GET("/plans", h.getAllPlans)
	api.POST("/plans", h.createPlan)
	api.GET("/plans/:planID", h.getPlan)
	api.PUT("/plans/:planID", h.updatePlan)
	api.DELETE("/plans/:planID", h.deletePlan)

	api.GET("/plans/:planID/tasks", h.getAllTasks)
	api.POST("/plans/:planID/tasks", h.createTask)
	api.GET("/plans/:planID/tasks/:taskID", h.getTask)
	api.PUT("/plans/:planID/tasks/:taskID", h.updateTask)
	api.DELETE("/plans/:planID/tasks/:taskID", h.deleteTask)
	api.POST("/plans/:planID/tasks/:taskID/done", h.doneTask)
	api.DELETE("/plans/:planID/tasks/:taskID/done", h.undoTask)
}
