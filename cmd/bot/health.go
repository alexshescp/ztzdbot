// cmd/bot/health.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func StartHealthServer(port int) {
	go func() {
		http.HandleFunc("healthz", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "ok")
		})

		addr := fmt.Sprintf(":%d", port)
		log.Printf("Health server started on %s", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Printf("Health server error: %v", err)
		}
	}()
}
