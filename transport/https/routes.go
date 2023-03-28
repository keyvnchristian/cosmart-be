package https

import (
	"cosmart/transport/https/middlewares"
	"cosmart/transport/https/service/books"
	"cosmart/transport/https/service/schedule"
	"cosmart/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Routes  *echo.Echo
	Usecase usecase.Usecase
}

func NewServer(e *echo.Echo, usecase usecase.Usecase) Server {
	return Server{
		Routes:  e,
		Usecase: usecase,
	}
}

func (s *Server) Initialize() {
	middleware := middlewares.GoMiddleware{}

	public := s.Routes.Group("/v1", middleware.CORS)

	books.NewBookHandler(public, s.Usecase)
	schedule.NewScheduleHandler(public, s.Usecase)
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString(`server.port`)
	}

	s.Routes.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	fmt.Println("Running Server ...")
	log.Fatal(http.ListenAndServe(":"+port, s.Routes))
}
