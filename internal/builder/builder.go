package builder

import (
	"github.com/wannn28/TASK-MIKTI/config"
	"github.com/wannn28/TASK-MIKTI/internal/http/handler"
	"github.com/wannn28/TASK-MIKTI/internal/http/router"
	"github.com/wannn28/TASK-MIKTI/internal/repository"
	"github.com/wannn28/TASK-MIKTI/internal/service"
	"github.com/wannn28/TASK-MIKTI/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	// repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	// end

	// service
	userService := service.NewUserService(cfg, userRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	movieService := service.NewMovieService(movieRepository)
	//end

	// handler
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	// end

	return router.PublicRoutes(movieHandler, userHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	// _ = repository.NewUserRepository(db)
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)

	userService := service.NewUserService(cfg, userRepository)
	movieService := service.NewMovieService(movieRepository)
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)

	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(tokenService, userService)

	return router.PrivateRoutes(movieHandler, userHandler)
}
