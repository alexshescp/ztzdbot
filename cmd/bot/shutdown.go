// cmd/bot/shutdown.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown() {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	s := <-sig
	log.Printf("Received shutdown signal: %s", s)
}
