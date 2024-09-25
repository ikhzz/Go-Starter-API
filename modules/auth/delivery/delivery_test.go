package delivery_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"starterapi/common/models"

	hpus "starterapi/common/usecase"
	"starterapi/modules/auth/delivery"
	aumod "starterapi/modules/auth/models"
	"starterapi/modules/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock common helper
type MockCommonUsecase struct {
	mock.Mock
}

func TestSignin(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUsecase := new(usecase.MockAuthUsecase)
	mockHelper := new(hpus.MockHelper)

	mockResponse := &models.GeneralResponse{
		StatusCode: http.StatusOK,
		Status:     true,
		Message:    []string{"success"},
		Data: aumod.ResSignin{
			Token: "token",
		},
	}

	mockUsecase.On("SignIn", &aumod.ReqPostSignin{Email: "john@Mail.com", Password: "password123"}).Return(mockResponse)

	deliv := delivery.AuthDelivery{
		Usecase: mockUsecase,
		Helper:  mockHelper,
	}

	r.POST("/v1/auth/signin", deliv.Signin)
	// Create request payload
	payload := aumod.ReqPostSignin{
		Email:    "john@Mail.com",
		Password: "password123",
	}
	body, _ := json.Marshal(payload)

	// Create request to signin route
	req, _ := http.NewRequest("POST", "/v1/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Record response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")

	// Verify expectations
	mockUsecase.AssertExpectations(t)
	mockHelper.AssertExpectations(t)
}
