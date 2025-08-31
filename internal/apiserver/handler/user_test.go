package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/service"
	svcmocks "github.com/jwcen/mars/internal/apiserver/service/mocks"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Signup(t *testing.T) {
	testCase := []struct {
		name     string
		mock     func(*gomock.Controller) service.UserAndService
		body     string
		wantCode int
		wantBody string
		wantErr  error
	}{
		{
			name: "注册成功",
			mock: func(c *gomock.Controller) service.UserAndService {
				userSvc := svcmocks.NewMockUserAndService(c)
				userSvc.EXPECT().Signup(gomock.Any(), &domain.User{
					Email:    "123@qq.com",
					Password: "admin123",
				}).Return(nil)
				return userSvc
			},
			body: `{
				"email": "123@qq.com",
				"password": "admin123",
				"confirmPassword": "admin123"
			}`,
			wantCode: http.StatusOK,
			wantBody: "注册成功！",
			wantErr:  nil,
		},
	}

	gin.SetMode(gin.ReleaseMode)
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := gin.Default()
			uh := NewUserHandler(tc.mock(ctrl), nil)
			uh.RegisterRoutes(r)

			req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(tc.body)))
			require.NoError(t, err)

			// 设置请求头
			req.Header.Set("Content-Type", "application/json")
			// http请求的记录
			resp := httptest.NewRecorder()

			// HTTP 请求进入 GIN 框架的入口
			// 调用此方法时，Gin 会处理这个请求，将响应写回 resp 里
			r.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			assert.Equal(t, tc.wantBody, resp.Body.String())

		})
	}
}
