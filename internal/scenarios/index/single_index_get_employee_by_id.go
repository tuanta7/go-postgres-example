package index

import (
	"context"
	"time"

	"github.com/tuanta7/go-postgres-example/internal/domain"
	"go.uber.org/zap"
)

func (s *SingleIndexScenario) GetEmployeeByID(ctx context.Context, id int64) (domain.Employee, error) {
	start := time.Now()
	employee, err := s.employees.GetEmployeeByID(ctx, id)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "GetEmployeeByID"),
		zap.Int64("employee_id", id),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("GetEmployeeByID failed", append(fields, zap.Error(err))...)
	} else {
		s.logger.Debug("GetEmployeeByID completed", fields...)
	}
	return employee, err
}

func (s *SingleIndexScenario) ExplainGetEmployeeByID(ctx context.Context, id int64) ([]string, error) {
	start := time.Now()
	plan, err := s.employees.ExplainGetEmployeeByID(ctx, id)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "ExplainGetEmployeeByID"),
		zap.Int64("employee_id", id),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("ExplainGetEmployeeByID failed", append(fields, zap.Error(err))...)
		return nil, err
	}
	s.logger.Debug("ExplainGetEmployeeByID completed", fields...)
	return plan, nil
}

func (s *SingleIndexScenario) ExplainAnalyzeGetEmployeeByID(ctx context.Context, id int64) ([]string, error) {
	start := time.Now()
	plan, err := s.employees.ExplainAnalyzeGetEmployeeByID(ctx, id)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "ExplainAnalyzeGetEmployeeByID"),
		zap.Int64("employee_id", id),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("ExplainAnalyzeGetEmployeeByID failed", append(fields, zap.Error(err))...)
		return nil, err
	}

	s.logger.Debug("ExplainAnalyzeGetEmployeeByID completed", fields...)
	return plan, nil
}
