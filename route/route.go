package route

import (
	"hearable/backend/internal/auth"
	"net/http"

	"context"

	"github.com/go-chi/chi/v5"

	"os"
	"os/signal"
	"syscall"

	"log/slog"

	"time"
)

func Start() {
	r := chi.NewRouter()
	authHandler := auth.NewHandler()
	r.Get("/test", authHandler.Test)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	apiServer := &http.Server{Handler: r, Addr: ":8080"}

	serverCtx, serverCtxCancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		logger.Info("Shutting down server..")

		shutdownCtx, shutdownCtxCancel := context.WithTimeout(serverCtx, 15*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Error("Graceful shut down timed out.. forcing exit.")
				os.Exit(1)
			}
			shutdownCtxCancel()
		}()

		err := apiServer.Shutdown(shutdownCtx)
		if err != nil {
			logger.Error("Error shutting down server", "error", err)
		}

		logger.Info("Server gracefully stopped")

		serverCtxCancel()
	}()

	go func() {
		logger.Info("Starting server on port 8080")
		err := apiServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("Error starting server", "error", err)
			serverCtxCancel()
		}
	}()

	<-serverCtx.Done()
}
