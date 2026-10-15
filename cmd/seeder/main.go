package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tuanta7/go-postgres-example/internal/repository/seed"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Seed employees:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	db, err := seed.NewDB(ctx, os.Getenv("NEON_POSTGRES_DSN"))
	if err != nil {
		return err
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = db.Close(closeCtx)
	}()
	return db.ImportEmployeesData(ctx)
}
