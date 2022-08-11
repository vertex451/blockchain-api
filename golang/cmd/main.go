package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"silverspase/blockchain-api/golang/internal/dep_container"
)

func main() {
	container, err := dep_container.New()
	if err != nil {
		panic(fmt.Sprintf("error initializing DI container: %+v", err))
	}

	go container.RunHTTPServer()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit

	if err := container.Delete(); err != nil {
		zap.L().Error("error deleting DI container", zap.Error(err))
	}
}
