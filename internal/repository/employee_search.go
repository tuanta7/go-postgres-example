package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/internal/repository/db"
)

// Only fixed predicates are appended; user values are always bound parameters.
func buildEmployeeSearch(filters domain.Employee) (string, []any) {
	var query strings.Builder
	query.WriteString("SELECT id, birth_date, first_name, last_name, gender, hire_date FROM employees.employee WHERE TRUE")
	args := make([]any, 0, 6)
	add := func(predicate string, value any) {
		args = append(args, value)
		fmt.Fprintf(&query, " AND %s $%d", predicate, len(args))
	}
	if filters.ID != 0 {
		add("id =", filters.ID)
	}
	if filters.FirstName != "" {
		add("first_name =", filters.FirstName)
	}
	if filters.LastName != "" {
		add("last_name =", filters.LastName)
	}
	if filters.Gender != "" {
		add("gender =", string(filters.Gender))
	}
	if !filters.BirthDate.IsZero() {
		add("birth_date =", filters.BirthDate)
	}
	if !filters.HireDate.IsZero() {
		add("hire_date =", filters.HireDate)
	}
	return query.String(), args
}

func (r *EmployeesRepository) SearchEmployees(ctx context.Context, filters domain.Employee) ([]domain.Employee, error) {
	query, args := buildEmployeeSearch(filters)
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Employee, error) {
		var employee db.EmployeesEmployee
		err := row.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate)
		return employeeToDomain(employee), err
	})
}

func (r *EmployeesRepository) ExplainSearchEmployees(ctx context.Context, filters domain.Employee) ([]string, error) {
	query, args := buildEmployeeSearch(filters)
	rows, err := r.db.Query(ctx, "EXPLAIN "+query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[string])
}

func (r *EmployeesRepository) ExplainAnalyzeSearchEmployees(ctx context.Context, filters domain.Employee) ([]string, error) {
	query, args := buildEmployeeSearch(filters)
	rows, err := r.db.Query(ctx, "EXPLAIN (ANALYZE, BUFFERS) "+query, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[string])
}
