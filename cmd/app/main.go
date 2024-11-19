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

	_, err = database.InitDatabase(cfg.MySQLConfig)
	fmt.Println("Sudah masuk ke database")
	// init  & start database

	publicRoutes := builder.BuildPublicRoutes()
	privateRoutes := builder.BuildPrivateRoutes()
	srv := server.NewServer(publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)
	// init & start server
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
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
