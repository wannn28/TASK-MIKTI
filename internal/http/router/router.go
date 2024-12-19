package router

import (
	"net/http"

	"github.com/wannn28/TASK-MIKTI/internal/http/handler"
	"github.com/wannn28/TASK-MIKTI/pkg/route"
)

var (
	adminOnly = []string{"Administrator"}
	// userOnly  = []string{"user"}
	allRoles = []string{"Administrator", "User"}
)

func PublicRoutes(movieHandler handler.MovieHandler, userHandler handler.UserHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: userHandler.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/request-reset-password",
			Handler: userHandler.ResetPasswordRequest,
		},
		{
			Method:  http.MethodPost,
			Path:    "/reset-password/:token",
			Handler: userHandler.ResetPassword,
		},
		{
			Method:  http.MethodGet,
			Path:    "/verify-email/:token",
			Handler: userHandler.VerifyEmail,
		},
	}
}

func PrivateRoutes(movieHandler handler.MovieHandler, userHandler handler.UserHandler) []route.Route {
	return []route.Route{

		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.GetUsers,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: userHandler.GetUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "/movies",
			Handler: movieHandler.GetMovies,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodGet,
			Path:    "/movies/:id",
			Handler: movieHandler.GetMovie,
			Roles:   allRoles,
		},
		{
			Method:  http.MethodPost,
			Path:    "/movies",
			Handler: movieHandler.CreateMovie,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/movies/:id",
			Handler: movieHandler.UpdateMovie,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/movies/:id",
			Handler: movieHandler.DeleteMovie,
			Roles:   adminOnly,
		},
	}
}
