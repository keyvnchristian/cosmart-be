package openlibrary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type libraryService struct {
	client http.Client
}

var _ LibraryService = (*libraryService)(nil)

func New(client http.Client) LibraryService {
	return &libraryService{
		client: client,
	}
}

func (uc *libraryService) FetchBook(ctx context.Context, req LibraryRequest) (LibraryResponse, error) {
	var resp LibraryResponse
	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json", req.Name)

	if req.IsDetail {
		url += `?details=true`
	}

	request, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	res, err := uc.client.Do(request)
	if err != nil {
		return LibraryResponse{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return LibraryResponse{}, err
	}

	return resp, nil
}
