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

func TestCreateOrder(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		inputPayload string
		mockReturn   *domain.Order
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid order creation",
			inputPayload: `{"customer_id":1,"items":[{"product_id":1,"quantity":2}]}`,
			mockReturn:   &domain.Order{ID: "1", CustomerID: "1", TotalAmount: 100.0},
			mockError:    nil,
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":1,"customer_id":1,"total_amount":100}`,
		},
		{
			name:         "Invalid request payload",
			inputPayload: `{"customer_id":1,"items":[]}`,
			mockReturn:   nil,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request payload\n",
		},
		{
			name:         "Service error",
			inputPayload: `{"customer_id":1,"items":[{"product_id":1,"quantity":2}]}`,
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturn != nil {
				mockService.On("CreateOrder", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(tt.mockReturn, tt.mockError)
			}

			req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(tt.inputPayload))
			rr := httptest.NewRecorder()

			handler.CreateOrder(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestGetOrderById(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		orderID      int64
		mockReturn   *domain.Order
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Order found",
			orderID:      1,
			mockReturn:   &domain.Order{ID: 1, CustomerID: 1, TotalAmount: 100.0},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"customer_id":1,"total_amount":100}`,
		},
		{
			name:         "Order not found",
			orderID:      2,
			mockReturn:   nil,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "order not found\n",
		},
		{
			name:         "Service error",
			orderID:      3,
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetOrderById", mock.Anything, tt.orderID).Return(tt.mockReturn, tt.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/orders?id="+strconv.FormatInt(tt.orderID, 10), nil)
			rr := httptest.NewRecorder()

			handler.GetOrderById(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestGetOrderById(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		orderID      int64
		mockReturn   *domain.Order
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Order found",
			orderID:      1,
			mockReturn:   &domain.Order{ID: "1", CustomerID: "1", TotalAmount: 100.0},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"customer_id":1,"total_amount":100}`,
		},
		{
			name:         "Order not found",
			orderID:      2,
			mockReturn:   nil,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "order not found\n",
		},
		{
			name:         "Service error",
			orderID:      3,
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetOrderById", mock.Anything, tt.orderID).Return(tt.mockReturn, tt.mockError)

			req, _ := http.NewRequest(http.MethodGet, "/orders?id="+strconv.FormatInt(tt.orderID, 10), nil)
			rr := httptest.NewRecorder()

			handler.GetOrderById(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		inputPayload string
		mockReturn   *domain.Order
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid order update",
			inputPayload: `{"id":1,"customer_id":1,"items":[{"product_id":1,"quantity":3}]}`,
			mockReturn:   &domain.Order{ID: 1, CustomerID: 1, TotalAmount: 150.0},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"customer_id":1,"total_amount":150}`,
		},
		{
			name:         "Order not found",
			inputPayload: `{"id":2,"customer_id":1,"items":[{"product_id":1,"quantity":3}]}`,
			mockReturn:   nil,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "order not found\n",
		},
		{
			name:         "Service error",
			inputPayload: `{"id":1,"customer_id":1,"items":[{"product_id":1,"quantity":3}]}`,
			mockReturn:   nil,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *Testing.T) {
			if tt.mockReturn != nil {
				mockService.On("UpdateOrder", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(tt.mockReturn, tt.mockError)
			}

			req, _ := http.NewRequest(http.MethodPut, "/orders", strings.NewReader(tt.inputPayload))
			rr := httptest.NewRecorder()

			handler.UpdateOrder(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	// Define test cases
	tests := []struct {
		name         string
		orderID      int64
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Order deleted successfully",
			orderID:      1,
			mockError:    nil,
			expectedCode: http.StatusNoContent,
			expectedBody: "",
		},
		{
			name:         "Order not found",
			orderID:      2,
			mockError:    ports.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: "order not found\n",
		},
		{
			name:         "Service error",
			orderID:      3,
			mockError:    someError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("DeleteOrder", mock.Anything, tt.orderID).Return(tt.mockError)

			req, _ := http.NewRequest(http.MethodDelete, "/orders?id="+strconv.FormatInt(tt.orderID, 10), nil)
			rr := httptest.NewRecorder()

			handler.DeleteOrder(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
