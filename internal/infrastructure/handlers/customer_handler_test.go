package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/witthawin0/go-hexagon/internal/domain"
)

// Mock the CustomerService interface
type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) CreateCustomer(customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerService) GetCustomerByID(id string) (*domain.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerService) GetAllCustomers() ([]*domain.Customer, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Customer), args.Error(1)
}

func (m *MockCustomerService) UpdateCustomer(id string, customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerService) DeleteCustomer(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetCustomerByID(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := NewCustomerHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		customerID   int64
		mockReturn   *domain.Customer
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Customer exists",
			customerID:   1,
			mockReturn:   &domain.Customer{ID: "1", Name: "John Doe"},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"name":"John Doe"}`,
		},
		{
			name:         "Customer not found",
			customerID:   2,
			mockReturn:   nil,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "customer not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetCustomerByID", mock.Anything, tt.customerID).Return(tt.mockReturn, tt.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/customers?id="+strconv.FormatInt(tt.customerID, 10), nil)
			rr := httptest.NewRecorder()

			handler.GetCustomerByID(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestCreateCustomer(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := NewCustomerHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		inputPayload string
		mockReturn   *domain.Customer
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid request",
			inputPayload: `{"name":"John Doe"}`,
			mockReturn:   &domain.Customer{ID: "1", Name: "John Doe"},
			mockError:    nil,
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":1,"name":"John Doe"}`,
		},
		{
			name:         "Invalid request payload",
			inputPayload: `{"name":123}`,
			mockReturn:   nil,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request payload\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturn != nil {
				mockService.On("CreateCustomer", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(tt.mockReturn, tt.mockError)
			}

			req, _ := http.NewRequest(http.MethodPost, "/customers", strings.NewReader(tt.inputPayload))
			rr := httptest.NewRecorder()

			handler.CreateCustomer(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestGetAllCustomers(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := NewCustomerHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		mockReturn   []*domain.Customer
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Customers found",
			mockReturn:   []*domain.Customer{{ID: "1", Name: "John Doe"}, {ID: "2", Name: "Jane Doe"}},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `[{"id":1,"name":"John Doe"},{"id":2,"name":"Jane Doe"}]`,
		},
		{
			name:         "No customers found",
			mockReturn:   []*domain.Customer{},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name:         "Service error",
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetAllCustomers", mock.Anything).Return(tt.mockReturn, tt.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/customers", nil)
			rr := httptest.NewRecorder()

			handler.GetAllCustomers(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := NewCustomerHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		inputPayload string
		mockReturn   *domain.Customer
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid request",
			inputPayload: `{"id":1,"name":"John Doe Updated"}`,
			mockReturn:   &domain.Customer{ID: "1", Name: "John Doe Updated"},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"name":"John Doe Updated"}`,
		},
		{
			name:         "Invalid request payload",
			inputPayload: `{"id":"abc","name":"John Doe Updated"}`,
			mockReturn:   nil,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request payload\n",
		},
		{
			name:         "Service error",
			inputPayload: `{"id":1,"name":"John Doe Updated"}`,
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturn != nil {
				mockService.On("UpdateCustomer", mock.Anything, mock.AnythingOfType("*domain.Customer")).Return(tt.mockReturn, tt.mockError)
			}

			req, _ := http.NewRequest(http.MethodPut, "/customers", strings.NewReader(tt.inputPayload))
			rr := httptest.NewRecorder()

			handler.UpdateCustomer(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := NewCustomerHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		customerID   int64
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Customer deleted successfully",
			customerID:   1,
			mockError:    nil,
			expectedCode: http.StatusNoContent,
			expectedBody: "",
		},
		{
			name:         "Customer not found",
			customerID:   2,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "customer not found\n",
		},
		{
			name:         "Service error",
			customerID:   3,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("DeleteCustomer", mock.Anything, tt.customerID).Return(tt.mockError)

			req, _ := http.NewRequest(http.MethodDelete, "/customers?id="+strconv.FormatInt(tt.customerID, 10), nil)
			rr := httptest.NewRecorder()

			handler.DeleteCustomer(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
