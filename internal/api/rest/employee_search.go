package rest

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/tuanta7/go-postgres-example/internal/domain"
)

func (h *EmployeeHandler) SearchEmployees(c fiber.Ctx) error {
	filters, err := employeeSearchFilters(c)
	if err != nil {
		return err
	}

	employees, err := h.singleIndex.SearchEmployees(c.Context(), filters)
	if err != nil {
		return employeeError(c, err)
	}

	result := make([]fiber.Map, 0, len(employees))
	for _, employee := range employees {
		result = append(result, employeeResponse(employee))
	}
	return c.JSON(result)
}

func (h *EmployeeHandler) ExplainSearchEmployees(c fiber.Ctx) error {
	filters, err := employeeSearchFilters(c)
	if err != nil {
		return err
	}

	plan, err := h.singleIndex.ExplainSearchEmployees(c.Context(), filters)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(fiber.Map{
		"plan": plan,
	})
}

func (h *EmployeeHandler) ExplainAnalyzeSearchEmployees(c fiber.Ctx) error {
	filters, err := employeeSearchFilters(c)
	if err != nil {
		return err
	}

	plan, err := h.singleIndex.ExplainAnalyzeSearchEmployees(c.Context(), filters)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(fiber.Map{
		"plan": plan,
	})
}

type employeeSearchRequest struct {
	ID        *int64  `query:"id" validate:"omitempty,gt=0"`
	FirstName *string `query:"first_name" validate:"omitempty,nonblank,max=14"`
	LastName  *string `query:"last_name" validate:"omitempty,nonblank,max=16"`
	Gender    *string `query:"gender" validate:"omitempty,oneof=M F"`
	BirthDate *string `query:"birth_date" validate:"omitempty,datetime=2006-01-02,ne=0001-01-01"`
	HireDate  *string `query:"hire_date" validate:"omitempty,datetime=2006-01-02,ne=0001-01-01"`
}

func employeeSearchFilters(c fiber.Ctx) (domain.Employee, error) {
	var request employeeSearchRequest
	var filters domain.Employee
	if err := c.Bind().Query(&request); err != nil {
		return filters, fiber.NewError(fiber.StatusBadRequest, "invalid search parameters")
	}

	if request.ID != nil {
		filters.ID = *request.ID
	}
	if request.FirstName != nil {
		filters.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		filters.LastName = *request.LastName
	}
	if request.Gender != nil {
		filters.Gender = domain.EmployeeGender(*request.Gender)
	}
	if request.BirthDate != nil {
		filters.BirthDate, _ = time.Parse("2006-01-02", *request.BirthDate)
	}
	if request.HireDate != nil {
		filters.HireDate, _ = time.Parse("2006-01-02", *request.HireDate)
	}

	return filters, nil
}
