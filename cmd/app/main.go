// Command server runs the personal finance tracker HTTP service.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"atad-project/internal/api"
	"atad-project/internal/store"
	"atad-project/internal/version"
)

func main() {
	for _, a := range os.Args[1:] {
		if a == "--version" || a == "-v" {
			if err := version.WriteJSON(os.Stdout); err != nil {
				log.Fatalf("version: %v", err)
			}
			return
		}
	}

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           api.NewServer(store.New()),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		v := version.Current()
		log.Printf("%s %s (%s) listening on %s", v.App, v.Version, v.Commit, addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()

	<-ctx.Done()
	log.Print("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
