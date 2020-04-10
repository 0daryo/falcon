package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	route "github.com/0daryo/falcon/internal/server/routes"
	"github.com/apex/log"
	"github.com/go-chi/chi"
)

var (
	revision string
	addr     = ":8080"
)

func Start() {

	baseCtx := context.Background()
	r := chi.NewRouter()
	route.Routing(r)
	// server
	server := http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(baseCtx, r),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("main: closed server")
		}
	}()
	log.Info(fmt.Sprintf("main: start server. listening on port%s", addr))
	// graceful shuttdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("SIGNAL %d received, so server shutting down now...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Error("failed to gracefully shutdown")
	}

	log.Info("main: server shutdown completed")
}
