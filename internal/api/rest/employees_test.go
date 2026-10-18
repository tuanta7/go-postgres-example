package rest

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tuanta7/go-postgres-example/internal/domain"
	"github.com/tuanta7/go-postgres-example/internal/mocks"
	"github.com/tuanta7/go-postgres-example/internal/scenarios/index"
	"github.com/tuanta7/go-postgres-example/pkg/otelx"
)

type EmployeeHandlerSuite struct {
	suite.Suite
	repo           *mocks.MockEmployeesRepository
	app            *fiber.App
	requestTimeout time.Duration
}

func (s *EmployeeHandlerSuite) SetupTest() {
	s.repo = mocks.NewMockEmployeesRepository(s.T())
	s.app = fiber.New()
	s.requestTimeout = 0
	s.app.Use(func(c fiber.Ctx) error {
		if s.requestTimeout > 0 {
			ctx, cancel := context.WithTimeout(c.Context(), s.requestTimeout)
			defer cancel()
			c.SetContext(ctx)
		}
		return c.Next()
	})

	logger := otelx.NewNopLogger()
	singleIndex := index.NewSingleIndexScenario(s.repo, logger)
	handler := NewEmployeeHandler(singleIndex)
	handler.RegisterRoutes(s.app)
}

func TestEmployeeHandlerSuite(t *testing.T) {
	suite.Run(t, new(EmployeeHandlerSuite))
}

func (s *EmployeeHandlerSuite) TestContextHasThirtySecondDeadline() {
	var hasDeadline bool
	var deadline, receivedAt time.Time
	s.repo.EXPECT().GetEmployeeByID(mock.Anything, int64(42)).RunAndReturn(
		func(ctx context.Context, _ int64) (domain.Employee, error) {
			deadline, hasDeadline = ctx.Deadline()
			receivedAt = time.Now()
			return domain.Employee{}, nil
		},
	).Once()

	startedAt := time.Now()
	response, err := s.app.Test(httptest.NewRequest("GET", "/v1/employees/42", nil))
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Require().True(hasDeadline, "Fiber timeout middleware must supply the query deadline")
	// The middleware creates the deadline between request start and repository invocation.
	s.GreaterOrEqual(deadline.Sub(startedAt), 30*time.Second)
	s.LessOrEqual(deadline.Sub(receivedAt), 30*time.Second)
}

func (s *EmployeeHandlerSuite) TestTimedOutResponse() {
	// An earlier parent deadline exercises the route timeout without a 30-second wait.
	s.requestTimeout = 50 * time.Millisecond

	release := make(chan struct{})
	defer close(release)

	s.repo.EXPECT().GetEmployeeByID(mock.Anything, int64(42)).RunAndReturn(
		func(ctx context.Context, _ int64) (domain.Employee, error) {
			<-ctx.Done()
			<-release // Keep the query pending
			return domain.Employee{}, ctx.Err()
		},
	).Once()

	response, err := s.app.Test(httptest.NewRequest("GET", "/v1/employees/42", nil))
	s.Require().NoError(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	s.Require().NoError(err)
	s.Equal(fiber.StatusGatewayTimeout, response.StatusCode)
	s.Equal("request timed out", string(body))
}
