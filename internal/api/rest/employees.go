package rest

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/timeout"
	"github.com/jackc/pgx/v5"
	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/internal/scenarios/index"
)

type EmployeeHandler struct {
	singleIndex *index.SingleIndexScenario
}

func NewEmployeeHandler(singleIndex *index.SingleIndexScenario) *EmployeeHandler {
	return &EmployeeHandler{
		singleIndex: singleIndex,
	}
}

func (h *EmployeeHandler) RegisterRoutes(g fiber.Router) {
	h.RegisterReadRoutes(g)
}

func (h *EmployeeHandler) RegisterReadRoutes(g fiber.Router) {
	group := g.Group("/v1/employees")
	cfg := timeout.Config{
		Timeout: 30 * time.Second,
		OnTimeout: func(c fiber.Ctx) error {
			return fiber.NewError(fiber.StatusGatewayTimeout, "request timed out")
		},
	}

	group.Get("/search", timeout.New(h.SearchEmployees, cfg))
	group.Get("/search/explain", timeout.New(h.ExplainSearchEmployees, cfg))
	group.Get("/search/explain-analyze", timeout.New(h.ExplainAnalyzeSearchEmployees, cfg))

	group.Get("/:id", timeout.New(h.GetEmployeeByID, cfg))
	group.Get("/:id/count", timeout.New(h.CountEmployeesByID, cfg))
	group.Get("/:id/explain", timeout.New(h.ExplainGetEmployeeByID, cfg))
	group.Get("/:id/explain-analyze", timeout.New(h.ExplainAnalyzeGetEmployeeByID, cfg))
}

func (h *EmployeeHandler) GetEmployeeByID(c fiber.Ctx) error {
	id, err := employeeID(c)
	if err != nil {
		return err
	}

	employee, err := h.singleIndex.GetEmployeeByID(c.Context(), id)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(employeeResponse(employee))
}

func (h *EmployeeHandler) CountEmployeesByID(c fiber.Ctx) error {
	id, err := employeeID(c)
	if err != nil {
		return err
	}

	count, err := h.singleIndex.CountEmployeesByID(c.Context(), id)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(fiber.Map{"employee_id": id, "count": count})
}

func (h *EmployeeHandler) ExplainGetEmployeeByID(c fiber.Ctx) error {
	id, err := employeeID(c)
	if err != nil {
		return err
	}

	var plan []string
	plan, err = h.singleIndex.ExplainGetEmployeeByID(c.Context(), id)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(fiber.Map{"employee_id": id, "plan": plan})
}

func (h *EmployeeHandler) ExplainAnalyzeGetEmployeeByID(c fiber.Ctx) error {
	id, err := employeeID(c)
	if err != nil {
		return err
	}

	var plan []string
	plan, err = h.singleIndex.ExplainAnalyzeGetEmployeeByID(c.Context(), id)
	if err != nil {
		return employeeError(c, err)
	}
	return c.JSON(fiber.Map{"employee_id": id, "plan": plan})
}

type employeeIDRequest struct {
	ID int64 `uri:"id" validate:"gt=0"`
}

func employeeID(c fiber.Ctx) (int64, error) {
	var request employeeIDRequest
	if err := c.Bind().URI(&request); err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "employee id must be a positive integer")
	}
	return request.ID, nil
}

func employeeResponse(employee domain.Employee) fiber.Map {
	return fiber.Map{
		"id":         employee.ID,
		"first_name": employee.FirstName,
		"last_name":  employee.LastName,
		"birth_date": employee.BirthDate.Format("2006-01-02"),
		"hire_date":  employee.HireDate.Format("2006-01-02"),
		"gender":     employee.Gender,
	}
}

func employeeError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "employee not found"})
	case errors.Is(err, context.DeadlineExceeded):
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{"error": "database query timed out"})
	default:
		// The scenario logs the original error; keep database details out of responses.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "database query failed"})
	}
}
