package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/wannn28/TASK-MIKTI/config"
	"github.com/wannn28/TASK-MIKTI/internal/builder"
	"github.com/wannn28/TASK-MIKTI/pkg/database"
	"github.com/wannn28/TASK-MIKTI/pkg/server"
)

func main() {
	// load configuration via env
	cfg, err := config.NewConfig(".env")
	checkError(err)
	// init  & start database
	db, err := database.InitDatabase(cfg.MySQLConfig)
	// RBAC
	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)
	// init & start server
	srv := server.NewServer(cfg, publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
