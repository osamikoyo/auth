package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"auth/api"
)

type Server struct {
	*echo.Echo
}

func New() Server {
	return Server{echo.New()}
}

func (s Server) Run() {
	s.Use(middleware.Logger())

	s.POST("/login", api.Login)
	s.POST("/profile", api.Profile)

	s.Logger.Panic(s.Start(":2020"))
}
