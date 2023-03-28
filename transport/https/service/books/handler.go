package books

import (
	"cosmart/model"
	"cosmart/usecase"
	"cosmart/usecase/service/books"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookHandler struct {
	BookUsecase books.BooksUsecase
}

func NewBookHandler(g *echo.Group, uc usecase.Usecase) {
	handler := &BookHandler{
		BookUsecase: uc,
	}

	g.GET("/books/:genre", handler.GetBooksByGenre)
}

func (h *BookHandler) GetBooksByGenre(c echo.Context) error {
	genre := c.Param("genre")

	resp, err := h.BookUsecase.FetchBooks(c.Request().Context(), model.BookRequest{
		Genre: genre,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.CommonResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.CommonResponse{
		Message: "Successfully Retrieved Books",
		Data:    resp,
	})
}
