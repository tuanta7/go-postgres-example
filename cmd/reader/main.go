package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tuanta7/go-postgres-example/internal/api/rest"
	"github.com/tuanta7/go-postgres-example/internal/repository"
	"github.com/tuanta7/go-postgres-example/internal/scenarios/index"
	"github.com/tuanta7/go-postgres-example/pkg/otelx"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger := otelx.NewLogger()
	defer logger.Sync()

	startupCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	pool, err := pgxpool.New(startupCtx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("configure database pool: %w", err)
	}
	defer pool.Close()
	if err := pool.Ping(startupCtx); err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	cancel()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code, message := fiber.StatusInternalServerError, "internal server error"
			var httpErr *fiber.Error
			if errors.As(err, &httpErr) {
				code, message = httpErr.Code, httpErr.Message
			}
			return c.Status(code).JSON(fiber.Map{"error": message})
		},
	})
	repo := repository.NewEmployeesRepository(pool)
	scenario := index.NewSingleIndexScenario(repo, logger)
	rest.NewEmployeeHandler(scenario).RegisterRoutes(app)

	logger.Info("starting reader API", zap.String("address", cfg.ServerAddr))
	return app.Listen(cfg.ServerAddr, fiber.ListenConfig{
		GracefulContext:       ctx,
		ShutdownTimeout:       35 * time.Second,
		DisableStartupMessage: true,
	})
}
