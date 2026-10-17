package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/internal/repository/db"
)

type EmployeesRepository struct {
	q  *db.Queries
	db db.DBTX
}

func NewEmployeesRepository(conn db.DBTX) *EmployeesRepository {
	return &EmployeesRepository{
		q:  db.New(conn),
		db: conn,
	}
}

func (r *EmployeesRepository) CountEmployeesByID(ctx context.Context, id int64) (int64, error) {
	return r.q.CountEmployeesByID(ctx, id)
}

func (r *EmployeesRepository) GetEmployeeByID(ctx context.Context, id int64) (domain.Employee, error) {
	employee, err := r.q.GetEmployeeByID(ctx, id)
	if err != nil {
		return domain.Employee{}, err
	}
	return domain.Employee{
		ID:        employee.ID,
		BirthDate: employee.BirthDate.Time,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Gender:    domain.EmployeeGender(employee.Gender),
		HireDate:  employee.HireDate.Time,
	}, nil
}

func (r *EmployeesRepository) ExplainGetEmployeeByID(ctx context.Context, id int64) ([]string, error) {
	rows, err := r.db.Query(ctx, `EXPLAIN SELECT * FROM employees.employee WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[string])
}

// ExplainAnalyzeEmployeeByID executes the lookup and returns actual timings and buffer usage.
// Repeated calls can warm the database cache.
func (r *EmployeesRepository) ExplainAnalyzeGetEmployeeByID(ctx context.Context, id int64) ([]string, error) {
	rows, err := r.db.Query(ctx, `EXPLAIN (ANALYZE, BUFFERS) SELECT * FROM employees.employee WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[string])
}
