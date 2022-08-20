package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/onemgvv/wb-l0/internal/domain"
	"github.com/onemgvv/wb-l0/internal/service"
	mockService "github.com/onemgvv/wb-l0/internal/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetByID(t *testing.T) {
	responseData := domain.OrderJSON{
		"id":   "f4bde549-ef88-4cff-a974-fd1abee0e598",
		"data": `{"orderID": "542", "sum": "1.300", "curr": "RUB"}`,
	}
	rdVal, _ := json.Marshal(responseData)

	responseBody := fmt.Sprintf(`{"statusCode":200,"message":"Order found","data":{"order":%s}}`, rdVal)

	type mockBehavior func(s *mockService.MockOrders, id string)

	testTable := []struct {
		name                 string
		paramName            string
		paramValue           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "OK",
			paramName:  "id",
			paramValue: "f4bde549-ef88-4cff-a974-fd1abee0e598",
			mockBehavior: func(s *mockService.MockOrders, id string) {
				s.EXPECT().GetById(id).Return(responseData, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: responseBody,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			order := mockService.NewMockOrders(c)
			testCase.mockBehavior(order, testCase.paramValue)

			services := &service.Service{Orders: order}
			handler := NewHandler(services)

			// Test server
			r := gin.New()
			r.GET("/order/:id", handler.GetByID)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/order/%s", testCase.paramValue), bytes.NewBufferString(""))

			// Perform request
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
