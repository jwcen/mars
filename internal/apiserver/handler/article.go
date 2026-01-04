package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jwcen/mars/internal/apiserver/domain"
	ijwt "github.com/jwcen/mars/internal/apiserver/handler/jwt"
	"github.com/jwcen/mars/internal/apiserver/service"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func NewArticleHandler(svc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
	}
}

func (h *ArticleHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	c, _ := ctx.Get("claims")
	claim, ok := c.(*ijwt.UserClaims)
	if !ok {
		// 你可以考虑监控住这里
		log.Println("op=Edit||err=user claims not found")
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}

	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	article := domain.Article{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			Id: claim.Id,
		},
	}

	if article.Title == "" || article.Content == "" {
		log.Printf("op=Edit||err=title or content is empty")
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "标题和内容不能为空",
		})
		return
	}

	articleId, err := h.svc.Save(ctx, article)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		log.Printf("op=Edit||err=%v", err)
		return
	}

	ctx.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "Success",
		Data: articleId,
	})
}
