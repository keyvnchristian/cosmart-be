package repository

import (
	"cosmart/model"
)

//go:generate mockgen -destination=../mocks/mocks_repository/repository.go -package=mocks_repository -source=./repository.go

type Repository interface {
	StoreBook(req model.Schedule) ([]model.Schedule, error)
}

type repository struct {
	Schedules *[]model.Schedule
}

func NewRepository(Schedules *[]model.Schedule) Repository {
	return &repository{
		Schedules: Schedules,
	}
}

func (r *repository) StoreBook(req model.Schedule) ([]model.Schedule, error) {
	*r.Schedules = append(*r.Schedules, req)
	return *r.Schedules, nil
}
