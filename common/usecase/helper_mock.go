package usecase

import (
	"io"
	"starterapi/common/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockHelper struct {
	mock.Mock
}

func (m *MockHelper) PanicCatcher(mw io.Writer) (r gin.HandlerFunc) {
	return
}
func (m *MockHelper) CustomLogger(mw io.Writer) (r gin.HandlerFunc) {
	return
}
func (m *MockHelper) CreateLog(param *models.LogModel) {
}
func (m *MockHelper) Validate(param interface{}) (e error, ss []string) {
	return
}
func (m *MockHelper) CreateToken(param models.JWTData) (s string) {
	return
}

func (m *MockHelper) JwtMiddleware(c *gin.Context) {

}
