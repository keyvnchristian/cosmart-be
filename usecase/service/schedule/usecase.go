package schedule

import (
	"cosmart/model"
	"cosmart/repository"
	"time"
)

type ScheduleUsecase interface {
	CreateSchedule(req model.ScheduleRequest) ([]model.Schedule, error)
}

type scheduleUsecase struct {
	ScheduleRepository repository.Repository
}

func NewScheduleUsecase(ScheduleRepository repository.Repository) ScheduleUsecase {
	return &scheduleUsecase{
		ScheduleRepository: ScheduleRepository,
	}
}

func (uc *scheduleUsecase) CreateSchedule(req model.ScheduleRequest) ([]model.Schedule, error) {
	date, _ := time.Parse(time.DateOnly, req.PickupSchedule)
	schedules, err := uc.ScheduleRepository.StoreBook(model.Schedule{
		UserID:     req.UserID,
		Title:      req.Title,
		Author:     req.Author,
		Edition:    req.Edition,
		PickUpTime: date,
	})
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
