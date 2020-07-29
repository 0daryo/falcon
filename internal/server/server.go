package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0daryo/falcon/internal/server/handler"
	route "github.com/0daryo/falcon/internal/server/routes"
	pb "github.com/0daryo/falcon/pb/server"
	"github.com/apex/log"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
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

func GRPCStart() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// サーバ起動
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handler.User{})
	log.Info(fmt.Sprintf("Listening on %v", addr))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
