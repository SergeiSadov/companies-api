package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"companies-api/internal/configs"
	"companies-api/internal/definitions"
	"companies-api/internal/pkg/constants"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	moduleMain = "main"
)

var rootCmd = &cobra.Command{
	Use:   "companies",
	Short: "companies crud service",
	Run:   RunService,
}

func RunService(_ *cobra.Command, _ []string) {
	di, err := definitions.Build()
	if err != nil {
		log.Fatalf("failed to create di, error: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	server := di.Get(definitions.HttpDef).(*http.Server)
	cfg := di.Get(definitions.CfgDef).(configs.Config)
	logger := di.Get(definitions.LoggerDef).(*zap.Logger)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				zap.L().Error(err.Error(), zap.String(constants.FieldModule, moduleMain))
			}
			return
		}
	}()
	logger.Info(fmt.Sprintf("started listening port %d", cfg.App.Port), zap.String(constants.FieldModule, moduleMain))

	sig := <-sigChan
	logger.Info(fmt.Sprintf("got signal %v, starting shutdown", sig), zap.String(constants.FieldModule, moduleMain))
	cancel()

	if err = server.Shutdown(ctx); err != nil {
		if err != http.ErrServerClosed {
			logger.Error(err.Error(), zap.String(constants.FieldModule, moduleMain), zap.String(constants.FieldAction, "server_shutdown"))
		}
	}

	os.Exit(0)
}
