// Package helper provides primitive functions which will be used in the application.
package helper

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// WaitForShutdown listens for system signals (SIGINT, SIGTERM) and waits for
// the context to be cancelled.
func WaitForShutdown(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			// wait for system signal && initialize termination in that case
			log.Print("waiting for the context to be cancelled")
			return
		case <-ctx.Done():
			log.Print("received a cancellation from context: shutting down")
			return
		}
	}
}
