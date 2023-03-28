package books

import (
	"context"
	"cosmart/model"
	"cosmart/routing/openlibrary"
	"fmt"
	"time"
)

type BooksUsecase interface {
	FetchBooks(ctx context.Context, req model.BookRequest) ([]model.BookResponse, error)
}

type bookUsecase struct {
	openLibraryService openlibrary.LibraryService
	contextTimeout     time.Duration
}

func NewBookUsecase(openLibraryService openlibrary.LibraryService, contextTimeout time.Duration) BooksUsecase {
	return &bookUsecase{
		openLibraryService: openLibraryService,
		contextTimeout:     contextTimeout,
	}
}

func (uc *bookUsecase) FetchBooks(c context.Context, req model.BookRequest) ([]model.BookResponse, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	resp, err := uc.openLibraryService.FetchBook(ctx, openlibrary.LibraryRequest{
		Name:     req.Genre,
		IsDetail: false,
	})
	if err != nil {
		fmt.Println("[FetchBooks] Failed", err)
		return nil, err
	}

	return uc.convertToBookResponse(resp), nil
}

func (*bookUsecase) convertToBookResponse(resp openlibrary.LibraryResponse) []model.BookResponse {
	bookResp := []model.BookResponse{}

	for _, v := range resp.Works {
		bookResp = append(bookResp, model.BookResponse{
			Title:   v.Title,
			Author:  v.Authors[0].Name,
			Edition: v.Edition,
		})
	}

	return bookResp
}
