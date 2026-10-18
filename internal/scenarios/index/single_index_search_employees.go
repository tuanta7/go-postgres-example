package index

import (
	"context"
	"time"

	"github.com/tuanta7/go-postgres-example/internal/domain"
	"go.uber.org/zap"
)

func (s *SingleIndexScenario) SearchEmployees(ctx context.Context, filters domain.Employee) ([]domain.Employee, error) {
	start := time.Now()
	employees, err := s.employees.SearchEmployees(ctx, filters)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "SearchEmployees"),
		zap.Any("filters", filters),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("SearchEmployees failed", append(fields, zap.Error(err))...)
		return nil, err
	}
	s.logger.Debug("SearchEmployees completed", fields...)
	return employees, nil
}

func (s *SingleIndexScenario) ExplainSearchEmployees(ctx context.Context, filters domain.Employee) ([]string, error) {
	start := time.Now()
	plan, err := s.employees.ExplainSearchEmployees(ctx, filters)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "ExplainSearchEmployees"),
		zap.Any("filters", filters),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("ExplainSearchEmployees failed", append(fields, zap.Error(err))...)
		return nil, err
	}
	s.logger.Debug("ExplainSearchEmployees completed", fields...)
	return plan, nil
}

func (s *SingleIndexScenario) ExplainAnalyzeSearchEmployees(ctx context.Context, filters domain.Employee) ([]string, error) {
	start := time.Now()
	plan, err := s.employees.ExplainAnalyzeSearchEmployees(ctx, filters)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "ExplainAnalyzeSearchEmployees"),
		zap.Any("filters", filters),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("ExplainAnalyzeSearchEmployees failed", append(fields, zap.Error(err))...)
		return nil, err
	}
	s.logger.Debug("ExplainAnalyzeSearchEmployees completed", fields...)
	return plan, nil
}
