package app

import (
	"Kitchenizer/api"
	"Kitchenizer/repository"
	"Kitchenizer/service"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		defer cancel()
		<-signalChan
	}()

	// Initialize dependencies and service
	s := service.Service{Repo: repository.NewFake()}

	errChan := make(chan error, 1)
	defer close(errChan)

	// Start API
	errChan <- api.Start(s)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
