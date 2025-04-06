package main

import (
	"fmt"
	"golang-agnostic-template/bootstrap"
	"golang-agnostic-template/src/pkg/cli"
	"golang-agnostic-template/src/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Inicializar logger
	log, err := logger.InitLogger()
	if err != nil {
		log.Error("Error al iniciar el Logger")
	}
	defer log.Sync()
	// Crear un canal para manejar las señales de interrupción
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	serverChan := cli.CliConfig.ServerChan // Usamos el canal de comandos que ya tienes en cli

	cli.Init()
	cli.ExecuteCLI()

	for {
		select {
		case serverType := <-serverChan:
			if err := bootstrap.RunApplication(serverType, log); err != nil {
				log.Fatal("Error al levantar la aplicación: ")
			}
		case sig := <-signalChan:
			fmt.Println(sig)
			log.Info("Señal recibida: ")
			log.Info("Terminando la aplicación...")
			return // Salir cuando se reciba la señal
		}
	}
}
