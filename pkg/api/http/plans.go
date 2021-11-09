package http

import (
	"github.com/ksavluk/go-todo-api/pkg/plan"
	"github.com/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const ErrPlanNotFound = "plan_not_found"

func (h *handlers) getAllPlans(c echo.Context) error {
	plans, err := h.planService.GetAll(c.Request().Context())
	if err != nil {
		return errors.Wrap(err, "get_all_plan")
	}

	return c.JSON(http.StatusOK, plans)
}

func (h *handlers) getPlan(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	plan, err := h.planService.Get(c.Request().Context(), planID)
	if err != nil {
		return errors.Wrap(err, "get_plan")
	}
	if plan == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrPlanNotFound)
	}

	return c.JSON(http.StatusOK, plan)
}

func (h *handlers) createPlan(c echo.Context) error {
	reqPlan := new(plan.Plan)
	if err := c.Bind(reqPlan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resPlan, err := h.planService.Create(c.Request().Context(), *reqPlan)
	if err != nil {
		return errors.Wrap(err, "create_plan")
	}
	if resPlan == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrPlanNotFound)
	}

	return c.JSON(http.StatusOK, resPlan)
}

func (h *handlers) updatePlan(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqPlan := new(plan.Plan)
	if err = c.Bind(reqPlan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	reqPlan.ID = planID

	resPlan, err := h.planService.Update(c.Request().Context(), *reqPlan)
	if err != nil {
		return errors.Wrap(err, "update_plan")
	}
	if resPlan == nil {
		return echo.NewHTTPError(http.StatusNotFound, ErrPlanNotFound)
	}

	return c.JSON(http.StatusOK, resPlan)
}

func (h *handlers) deletePlan(c echo.Context) error {
	planID, err := getParamUint(c, ParamPlanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.planService.Delete(c.Request().Context(), planID)
	if err != nil {
		return errors.Wrap(err, "delete_plan")
	}

	return c.NoContent(http.StatusOK)
}
