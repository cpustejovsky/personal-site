package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cpustejovsky/personal-site/handlers"
)

func main() {

	// // Configuration
	// network := flag.String("n", "tcp", "network to listen on")
	// address := flag.String("a", ":8080", "address to listen on")
	// flag.Parse()

	port := os.Getenv("PORT")
	network := os.Getenv("NETWORK")
	if len(network) == 0 {
		network = "tcp"
	}
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	l, err := net.Listen(network, port)
	if err != nil {
		slog.Error("failed to listen on network",
			"network", network, "address", port, "error message", err.Error())
		os.Exit(1)
	}

	h, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	svr := http.Server{
		Handler: h,
	}

	// run server in a goroutine so we can multiplex between signal and error
	// handling below.
	errCh := make(chan error, 1)
	go func() {
		slog.Info("Server Started", "network", network, "address", port)
		errCh <- svr.Serve(l)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	select {
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	case <-ctx.Done():
		slog.Error("server shutting down", "error", ctx.Err())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := svr.Shutdown(ctx)
		if err != nil {
			slog.Error("failed to shutdown server, exiting anyway", "error", err)
			os.Exit(1)

		}
		slog.Info("Server shut down successfully")

	}

}
