package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	ServerType string
	ServerChan chan string
}

var CliConfig = Config{
	ServerChan: make(chan string, 1),
}

var rootCmd = &cobra.Command{
	Use:   "go",
	Short: "This template is made with ❤️",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Comandos para iniciar servidores",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Muestra la versión de la plantilla",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}

func Init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	serveCmd.PersistentFlags().StringVarP(&CliConfig.ServerType, "type", "t", CliConfig.ServerType, "Tipo de servidor (web|nats)")
	serveCmd.MarkFlagRequired("type")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "ayuda para mi template")
	initCommands()
}

// Execute the CLI and return the channel
func ExecuteCLI() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error ejecutando comando: %v", err)
		os.Exit(1)
	}
}

// Serve command for the "web" server
var serveWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Inicia un la aplicación como servidor web",
	Run: func(cmd *cobra.Command, args []string) {
		CliConfig.ServerType = "web"
		CliConfig.ServerChan <- "web" // Envia el tipo de servidor al canal
	},
}

var serveNatsCmd = &cobra.Command{
	Use:   "nats",
	Short: "Inicia el servidor NATS",
	Run: func(cmd *cobra.Command, args []string) {
		CliConfig.ServerType = "nats"
		CliConfig.ServerChan <- "nats" // Envia el tipo de servidor al canal
		fmt.Println("Iniciando servidor NATS...")
	},
}

func initServeCommands() {
	serveCmd.AddCommand(serveWebCmd)
	serveCmd.AddCommand(serveNatsCmd)
}

func initCommands() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
	initServeCommands()
}
