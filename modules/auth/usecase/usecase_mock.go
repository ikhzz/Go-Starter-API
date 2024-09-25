package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	cmmd "starterapi/common/models"
	"starterapi/modules/auth/models"
)

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) SignIn(c *gin.Context, req *models.ReqPostSignin) *cmmd.GeneralResponse {
	args := m.Called(req)

	// Return the mock GeneralResponse struct from args
	if response, ok := args.Get(0).(*cmmd.GeneralResponse); ok {
		return response
	}
	return nil
}

func (m *MockAuthUsecase) GetProfile(g *gin.Context, req *models.ReqGetProfile) *cmmd.GeneralResponse {
	args := m.Called(req)

	// Return the mock GeneralResponse struct from args
	if response, ok := args.Get(0).(*cmmd.GeneralResponse); ok {
		return response
	}

	return nil
}
