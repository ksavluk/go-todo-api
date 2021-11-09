package http

import (
	"github.com/ksavluk/go-todo-api/pkg/task"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const ErrTaskNotFound = "task_not_found"

func (h *handlers) getAllTasks(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tasks, err := h.taskService.GetAll(c.Request().Context(), planID)
	if err != nil {
		return errors.Wrap(err, "get_all_task")
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *handlers) getTask(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	taskID, err := getParamUint(c, ParamTaskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.Get(c.Request().Context(), planID, taskID)
	if err != nil {
		return errors.Wrap(err, "get_task")
	}
	if task == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrTaskNotFound)
	}

	return c.JSON(http.StatusOK, task)
}

func (h *handlers) createTask(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqTask := new(task.Task)
	if err = c.Bind(reqTask); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resTask, err := h.taskService.Create(c.Request().Context(), planID, *reqTask)
	if err != nil {
		return errors.Wrap(err, "create_task")
	}
	if resTask == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrTaskNotFound)
	}

	return c.JSON(http.StatusOK, resTask)
}

func (h *handlers) updateTask(c echo.Context) error {
	planID, taskID, err := getTaskParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqTask := new(task.Task)
	if err = c.Bind(reqTask); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	reqTask.ID = taskID

	resTask, err := h.taskService.Update(c.Request().Context(), planID, *reqTask)
	if err != nil {
		return errors.Wrap(err, "update_task")
	}
	if resTask == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrTaskNotFound)
	}

	return c.JSON(http.StatusOK, resTask)
}

func (h *handlers) deleteTask(c echo.Context) error {
	_, taskID, err := getTaskParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.taskService.Delete(c.Request().Context(), taskID)
	if err != nil {
		return errors.Wrap(err, "delete_task")
	}

	return c.NoContent(http.StatusOK)
}

func (h *handlers) doneTask(c echo.Context) error {
	planID, taskID, err := getTaskParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.Done(c.Request().Context(), planID, taskID)
	if err != nil {
		return errors.Wrap(err, "done_task")
	}
	if task == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrTaskNotFound)
	}

	return c.JSON(http.StatusOK, task)
}

func (h *handlers) undoTask(c echo.Context) error {
	planID, taskID, err := getTaskParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.Undo(c.Request().Context(), planID, taskID)
	if err != nil {
		return errors.Wrap(err, "done_task")
	}
	if task == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrTaskNotFound)
	}

	return c.JSON(http.StatusOK, task)
}
