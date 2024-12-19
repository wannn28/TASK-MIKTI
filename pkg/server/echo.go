package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/wannn28/TASK-MIKTI/config"
	"github.com/wannn28/TASK-MIKTI/internal/entity"
	"github.com/wannn28/TASK-MIKTI/pkg/response"
	"github.com/wannn28/TASK-MIKTI/pkg/route"
)

type Server struct {
	*echo.Echo
}

func NewServer(cfg *config.Config, publicRoutes, privateRoutes []route.Route) *Server {
	e := echo.New()
	v1 := e.Group("/api/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	if len(privateRoutes) > 0 {
		for _, route := range privateRoutes {
			v1.Add(route.Method, route.Path, route.Handler, JWTMiddleware(cfg.JWTConfig.SecretKey), RBACMiddleware(route.Roles))
		}
	}
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return &Server{e}
}

func JWTMiddleware(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTCustomClaims)
		}, SigningKey: []byte(secretKey),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Harus Login"))
		},
	})
}

func RBACMiddleware(Roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*entity.JWTCustomClaims)
			allowed := false
			for _, role := range Roles {
				if role == claims.Role {
					allowed = true
					break
				}
			}
			if !allowed {
				return c.JSON(http.StatusForbidden, response.ErrorResponse(http.StatusForbidden, "Anda tidak memiliki akses ke dalam sini"))
			}
			return next(c)
		}
	}
}
