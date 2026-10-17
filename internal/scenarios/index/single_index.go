package index

import (
	"context"
	"time"

	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/pkg/otelx"
	"go.uber.org/zap"
)

type EmployeesRepository interface {
	CountEmployeesByID(context.Context, int64) (int64, error)
	GetEmployeeByID(context.Context, int64) (domain.Employee, error)
	ExplainGetEmployeeByID(context.Context, int64) ([]string, error)
	ExplainAnalyzeGetEmployeeByID(context.Context, int64) ([]string, error)
}

type SingleIndexScenario struct {
	employees EmployeesRepository
	logger    *otelx.Logger
}

func NewSingleIndexScenario(employees EmployeesRepository, logger *otelx.Logger) *SingleIndexScenario {
	return &SingleIndexScenario{
		employees: employees,
		logger:    logger,
	}
}

func (s *SingleIndexScenario) CountEmployeesByID(ctx context.Context, id int64) (int64, error) {
	start := time.Now()
	count, err := s.employees.CountEmployeesByID(ctx, id)
	fields := []zap.Field{
		zap.String("scenario", "single_index"),
		zap.String("operation", "CountEmployeesByID"),
		zap.Int64("employee_id", id),
		zap.Duration("duration", time.Since(start)),
	}
	if err != nil {
		s.logger.Error("CountEmployeesByID failed", append(fields, zap.Error(err))...)
		return 0, err
	}

	s.logger.Debug("CountEmployeesByID completed", fields...)
	return count, nil
}

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
