package index

import (
	"context"
	"time"

	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/pkg/otelx"
	"go.uber.org/zap"
)

type EmployeesRepository interface {
	SearchEmployees(context.Context, domain.Employee) ([]domain.Employee, error)
	ExplainAnalyzeSearchEmployees(context.Context, domain.Employee) ([]string, error)
	ExplainSearchEmployees(context.Context, domain.Employee) ([]string, error)

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
