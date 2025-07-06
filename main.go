package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/config"
	"golang-agnostic-template/src/pkg/database"
	"golang-agnostic-template/src/pkg/logger"
	"golang-agnostic-template/src/pkg/webserver"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := config.LoadConfig(ctx)

	if err != nil {
		fmt.Printf("cannot load %s: %v", utils.ENVIRONMENT, err)
		return
	}

	log, err := logger.NewLogger()
	if err != nil {
		fmt.Printf("cannot load %s: %v", utils.LOGGER, err)
		return
	}
	defer log.Sync()

	db := database.NewSurrealDBConnection()
	defer db.Close()

	ctx, server := webserver.NewServer(ctx)
	server.Routes(ctx, log)
	err = server.Run(ctx)
	if err != nil {
		log.Fatal("cannot initialize web server", logger.LoggerField{Key: utils.WEB_SERVER, Value: err})
		return
	}
}
