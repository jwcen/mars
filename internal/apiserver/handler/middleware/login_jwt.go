package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	ijwt "github.com/jwcen/mars/internal/apiserver/handler/jwt"
)

// LoginJWTMiddlewareBuilder JWT 登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	// 用 Go 的方式编码解码
	return func(ctx *gin.Context) {
		// 不需要登录校验的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		// 我现在用 JWT 来校验
		tokenHeader := ctx.GetHeader("Authorization")
		log.Printf("[JWT Middleware] 请求路径: %s, Authorization header: %s", ctx.Request.URL.Path, tokenHeader)
		if tokenHeader == "" {
			// 没登录
			log.Printf("[JWT Middleware] 错误: Authorization header 为空")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//segs := strings.SplitN(tokenHeader, " ", 2)
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			// 没登录，有人瞎搞
			log.Printf("[JWT Middleware] 错误: Authorization header 格式不正确: %s", tokenHeader)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		claims := &ijwt.UserClaims{}
		// ParseWithClaims 里面，一定要传入指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), nil
		})
		if err != nil {
			// 没登录
			log.Printf("[JWT Middleware] 错误: JWT 解析失败: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//claims.ExpiresAt.Time.Before(time.Now()) {
		//	// 过期了
		//}
		// err 为 nil，token 不为 nil
		if token == nil || !token.Valid || claims.Id == 0 {
			// 没登录
			log.Printf("[JWT Middleware] 错误: Token 无效 - token==nil: %v, Valid: %v, Id: %d", token == nil, token != nil && token.Valid, claims.Id)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Printf("[JWT Middleware] Token 验证通过 - Id: %d, UserAgent: %s, 请求UserAgent: %s", claims.Id, claims.UserAgent, ctx.Request.UserAgent())
		if claims.UserAgent != ctx.Request.UserAgent() {
			// 严重的安全问题
			// 你是要监控
			log.Printf("[JWT Middleware] 错误: UserAgent 不匹配: token中的=%s, 请求中的=%s", claims.UserAgent, ctx.Request.UserAgent())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		// 每十秒钟刷新一次
		if claims.ExpiresAt != nil && claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err = token.SignedString([]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"))
			if err != nil {
				// 记录日志
				log.Println("jwt 续约失败", err)
			}
			ctx.Header("x-jwt-token", tokenStr)
		}
		ctx.Set("claims", claims)
		//ctx.Set("userId", claims.Uid)
	}
}
