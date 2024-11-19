package server

import (
	"github.com/labstack/echo/v4"
	"github.com/wannn28/TASK-MIKTI/pkg/route"
)

type Server struct {
	*echo.Echo
}

func NewServer(publicRoutes, privateRoutes []route.Route) *Server {
	e := echo.New()

	v1 := e.Group("app/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		for _, route := range privateRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}

	return &Server{e}
}
