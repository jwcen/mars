package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	ijwt "github.com/jwcen/mars/internal/apiserver/handler/jwt"
	"github.com/jwcen/mars/internal/apiserver/integration/startup"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func (s *ArticleTestSuite) SetupSuite() {
	s.server = gin.Default()
	s.server.Use(func(ctx *gin.Context) {
		ctx.Set("claims", &ijwt.UserClaims{
			Id: 123,
		})
	})
	s.db = startup.InitTestDB()
	artHdl := startup.InitArticleHandler()
	artHdl.RegisterRoutes(s.server)
}

func (s *ArticleTestSuite) TearDownTest() {
	s.db.Exec("TRUNCATE TABLE article")
}

func (s *ArticleTestSuite) Test_EditArticle() {
	t := s.T()
	testCases := []struct {
		name string

		art Article

		// 集成测试准备数据
		before func(t *testing.T)
		// 集成测试验证数据
		after func(t *testing.T)

		// http code
		wantCode int

		wantResult Result[int64]
	}{
		{
			name: "新建文章成功",
			art: Article{
				Title:   "test",
				Content: "test",
			},
			wantCode: http.StatusOK,
			wantResult: Result[int64]{
				Code: 0,
				Msg:  "Success",
				Data: 1,
			},
			before: func(t *testing.T) {},
			after: func(t *testing.T) {
				// check article in db
				var dbArt model.ArticleM
				err := s.db.Where("id = ?", 1).First(&dbArt).Error
				assert.NoError(t, err)
				assert.True(t, dbArt.UpdatedAt != time.Time{})
				assert.True(t, dbArt.CreatedAt != time.Time{})

				dbArt.CreatedAt = time.Time{}
				dbArt.UpdatedAt = time.Time{}

				assert.Equal(t, model.ArticleM{
					Id:       1,
					Title:    "test",
					Content:  "test",
					AuthorId: 123,
				}, dbArt)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			reqBody, err := json.Marshal(tc.art)
			assert.NoError(t, err)
			httpReq, err := http.NewRequest(http.MethodPost, "/api/v1/articles/edit", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			httpReq.Header.Set("Content-Type", "application/json")
			httpResp := httptest.NewRecorder()

			s.server.ServeHTTP(httpResp, httpReq)
			if httpResp.Code != 200 {
				t.Fatalf("httpResp.Code = %d, want %d", httpResp.Code, 200)
			}
			var result Result[int64]
			err = json.NewDecoder(httpResp.Body).Decode(&result)
			assert.NoError(t, err)

			assert.Equal(t, tc.wantCode, httpResp.Code)
			assert.Equal(t, tc.wantResult, result)
			tc.after(t)
		})
	}
}

func Test_ArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Result[T any] struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data T      `json:"data"`
}
