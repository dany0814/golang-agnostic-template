package bootstrap

import (
	"context"
	"fmt"
	"golang-agnostic-template/src/pkg/config"
	"golang-agnostic-template/src/pkg/logger"
)

func RunApplication(param string, log logger.ILogger) error {
	ctx := context.Background()
	err := config.LoadConfig(ctx)
	if err != nil {
		return err
	}

	fmt.Println("<---------->")
	fmt.Println(param)
	switch param {
	case "web":
		RunWebServer(ctx)
	case "nats":
		RunNatsServer(ctx)
	default:
		fmt.Println("default")
	}
	return nil
}
