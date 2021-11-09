package http

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const (
	ParamPlanID = "planID"
	ParamTaskID = "taskID"
)

func getParam(c echo.Context, key string) (string, error) {
	value := c.Param(key)
	if len(value) == 0 {
		return "", fmt.Errorf("path param '%s' is required", key)
	}
	return value, nil
}

func getParamUint(c echo.Context, key string) (uint, error) {
	str, err := getParam(c, key)
	if err != nil {
		return 0, err
	}

	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parse_uint")
	}
	return uint(u), nil
}

func getTaskParams(c echo.Context) (planID, taskID uint, err error) {
	planID, err = getParamUint(c, ParamPlanID)
	if err != nil {
		return
	}

	taskID, err = getParamUint(c, ParamTaskID)
	return
}
