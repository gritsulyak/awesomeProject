package v1

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/gritsulyak/awesomeProject/internal/model"
)

// 1. Define a pure, minimalist mock for the service
type mockSatelliteService struct {
	getSatelliteByNameFunc func(ctx context.Context, name string) (*model.Satellite, error)
}

func (m *mockSatelliteService) GetSatelliteByName(ctx context.Context, name string) (*model.Satellite, error) {
	return m.getSatelliteByNameFunc(ctx, name)
}

func TestController_Get(t *testing.T) {
	// Sample testing data
	targetSatellite := &model.Satellite{Name: "Hubble"}

	tests := []struct {
		name           string
		paramValue     string
		mockBehavior   func(ctx context.Context, name string) (*model.Satellite, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:       "Success - 200 OK",
			paramValue: "Hubble",
			mockBehavior: func(ctx context.Context, name string) (*model.Satellite, error) {
				return targetSatellite, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"name":"Hubble"}`, // Matches the fields inside model.Satellite
		},
		{
			name:       "Failure - 500 Internal Server Error",
			paramValue: "Hubble",
			mockBehavior: func(ctx context.Context, name string) (*model.Satellite, error) {
				return nil, errors.New("database breakdown")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"database breakdown"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange: Setup isolated Echo Context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Mock path parameters explicitly
			c.SetPath("/:name")
			c.SetParamNames("name")
			c.SetParamValues(tt.paramValue)

			mockSvc := &mockSatelliteService{
				getSatelliteByNameFunc: tt.mockBehavior,
			}
			ctl := &Controller{service: mockSvc}

			// Act: Execute only the controller handler function
			err := ctl.Get(c)

			// Assert: Verify results are exact
			if err != nil {
				t.Fatalf("unexpected error returned from handler: %v", err)
			}

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Clean strings from trailing spaces/newlines for comparison
			gotBody := strings.TrimSpace(rec.Body.String())
			if gotBody != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, gotBody)
			}
		})
	}
}
