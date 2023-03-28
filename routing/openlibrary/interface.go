package openlibrary

import "context"

//go:generate mockgen -destination=../../mocks/mocks_routing/mocks_openlibrary/openlibrary.go -package=mocks_openlibrary -source=./interface.go

type LibraryService interface {
	FetchBook(ctx context.Context, req LibraryRequest) (LibraryResponse, error)
}
