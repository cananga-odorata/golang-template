package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	RegisterFunc func(req RegisterRequest) (*AuthResponse, error)
}

// implement the interface
func (m *MockService) Register(req RegisterRequest) (*AuthResponse, error) {
	return m.RegisterFunc(req) //call function by us injection before
}

func TestRegister(t *testing.T) {
	t.Run("Search: should return 201 when input is valid", func(t *testing.T) {
		// Setup Mock: service response success always
		mockSvc := &MockService{
			RegisterFunc: func(req RegisterRequest) (*AuthResponse, error) {
				return &AuthResponse{
					AccessToken: "test-token",
					User:        User{ID: "1", Email: req.Email},
				}, nil
			},
		}

		// Init Handler by Mock
		handler := NewHandler(mockSvc)

		// Create request and response recorder
		reqBody := []byte(`{"email":"test@example.com","password":"password123"}`)
		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
		rec := httptest.NewRecorder()

		// Call handler
		handler.Register(rec, req)

		// Check response status code
		if rec.Code != http.StatusCreated {
			t.Errorf("expected status 201, got %d", rec.Code)
		}

		var resp map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)
		if resp["success"] != true {
			t.Errorf("expected success true")
		}
	})

	t.Run("Fail: should return 400 when body is invalid", func(t *testing.T) {
		// Not setup mock because Validation fail before call Service
		handler := NewHandler(&MockService{})

		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer([]byte(`{"email":"test@example.com","password":"password123"`)))
		rec := httptest.NewRecorder()

		handler.Register(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", rec.Code)
		}
	})

	t.Run("Fail: should return 500 when service fails", func(t *testing.T) {
		// Setup Mock: define service Error Always
		mockSvc := &MockService{
			RegisterFunc: func(req RegisterRequest) (*AuthResponse, error) {
				return nil, errors.New("Database Erorr")
			},
		}

		handler := NewHandler(mockSvc)

		reqBody := []byte(`{"email": "test@example.com", "password": "password123"}`)
		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(reqBody))
		rec := httptest.NewRecorder()

		handler.Register(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500, got %d", rec.Code)
		}
	})
}
