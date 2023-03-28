package schedule

import (
	"cosmart/model"
	"cosmart/transport/helper"
	"cosmart/usecase"
	"cosmart/usecase/service/schedule"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScheduleHandler struct {
	ScheduleUsecase schedule.ScheduleUsecase
}

func NewScheduleHandler(g *echo.Group, uc usecase.Usecase) {
	handler := &ScheduleHandler{
		ScheduleUsecase: uc,
	}

	g.POST("/schedule", handler.CreateSchedule)
}

func (h *ScheduleHandler) CreateSchedule(c echo.Context) error {
	var req model.ScheduleRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if ok, err := helper.IsRequestValid(&req); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	schedules, err := h.ScheduleUsecase.CreateSchedule(model.ScheduleRequest{
		UserID:         req.UserID,
		Title:          req.Title,
		Author:         req.Author,
		Edition:        req.Edition,
		PickupSchedule: req.PickupSchedule,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.CommonResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.CommonResponse{
		Message: "Successfully Add Schedule",
		Data:    schedules,
	})
}
