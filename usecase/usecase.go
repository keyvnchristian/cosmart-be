package usecase

import (
	"cosmart/repository"
	"cosmart/routing/openlibrary"
	"cosmart/usecase/service/books"
	"cosmart/usecase/service/schedule"
	"time"
)

type Usecase interface {
	books.BooksUsecase
	schedule.ScheduleUsecase
}

func NewUsecase(deps Dependencies) Usecase {
	return &struct {
		books.BooksUsecase
		schedule.ScheduleUsecase
	}{
		books.NewBookUsecase(deps.OpenLibraryService, deps.ContextTimeout),
		schedule.NewScheduleUsecase(deps.Repository),
	}
}

type Dependencies struct {
	Repository         repository.Repository
	OpenLibraryService openlibrary.LibraryService
	ContextTimeout     time.Duration
}
