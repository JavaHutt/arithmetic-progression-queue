package action

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/http"
)

type shutdownConfig interface {
	HTTPShutdownTimeout() time.Duration
	ServiceShutdownTimeout() time.Duration
}

func GracefulShutdown(
	cfg shutdownConfig,
	errorChannel chan error,
	httpServer http.Server,
	doneChannel chan bool,
) {
	// Capture interrupts.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s %v", <-c, time.Now())
	}()

	if err := <-errorChannel; err != nil {
		fmt.Println(err)
		close(doneChannel)

		httpServerShutdown(httpServer, cfg.HTTPShutdownTimeout())

		fmt.Println("app stopped", time.Now())
	}
	os.Exit(1)
}

func httpServerShutdown(httpServer http.Server, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := httpServer.Close(ctx); err != nil {
		fmt.Printf("could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("http server stopped", time.Now())
}
