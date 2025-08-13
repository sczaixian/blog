package system

import (
	"blog/server/models"
	common_response "blog/server/models/common/response"
	"blog/server/models/request"
	//"blog/server/models/response"
	"blog/server/utils"
	"fmt"
	"strconv"
	"errors"
	"github.com/gin-gonic/gin"
)

type ArticleApi struct{}

func (a *ArticleApi) CreateArticle(c *gin.Context) {
	var article request.EditArticle
	err := c.ShouldBind(&article)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(article, utils.ArticleCreate)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	art := &models.Article{Title: article.Title, Content: article.Content, Excerpt: article.Excerpt, CategoryID: article.CategoryID}

	if err = articleService.CreateArticle(art); err != nil {
		common_response.FailWithMessage("发布失败", c)
		return
	}
	common_response.OkWithMessage("发布成功", c)
}

func (a *ArticleApi) checkID(c *gin.Context) (uint, error) {
	artID := c.Param("id")
	if artID == "" {
		return 0, errors.New("文章ID不能为空")
	}
	id, err := strconv.ParseUint(artID, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("无效的文章ID: %v", err)
	}
	return uint(id), nil
}

func (a *ArticleApi) DeleteArticle(c *gin.Context) {
	id, err := a.checkID(c)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = articleService.DeleteArticle(uint(id))
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	common_response.OkWithMessage("删除成功", c)
}

func (a *ArticleApi) EditArticle(c *gin.Context) {
	var r request.EditArticle
	err := c.ShouldBindJSON(&r)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ArticleUpdate)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	article := &models.Article{Title: r.Title, Content: r.Content, Excerpt: r.Excerpt, CategoryID: r.CategoryID}

	id, err := a.checkID(c)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = articleService.EditArticle(article, id)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	common_response.OkWithMessage("修改成功", c)
}

func (a *ArticleApi) GetArticle(c *gin.Context) {
	id, err := a.checkID(c)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	art, err := articleService.GetArticle(id)
	// todo:
}

func (a *ArticleApi) ListArticle(c *gin.Context) {

}

func (a *ArticleApi) SaveArticle(c *gin.Context) {

}
