package system

import (
	"blog/server/models"
	common_request "blog/server/models/common/request"
	common_response "blog/server/models/common/response"
	"blog/server/models/request"
	"blog/server/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct{}

func (a *ArticleApi) CreateArticle(c *gin.Context) {
	var article request.ModifyArticleBase
	err := c.ShouldBindJSON(&article)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(article, utils.ArticleCreate)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	art := &models.Article{Title: article.Title, Content: article.Content, Excerpt: article.Excerpt, CategoryID: article.CategoryID, UserID: article.UserID, Status: 1}

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

func (a *ArticleApi) GetArticle(c *gin.Context) {
	id, err := a.checkID(c)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	art, err := articleService.GetArticle(id)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	common_response.OkWithDetailed(art, "success", c)
}

func (a *ArticleApi) ListArticle(c *gin.Context) {
	var pageInfo common_request.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := articleService.ListArticle(2, pageInfo)
	if err != nil {
		//log
		common_response.FailWithMessage("获取文章列表失败", c)
		return
	}
	common_response.OkWithDetailed(common_response.PageResult{
		List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize,
	}, "列表获取成功", c)
}

func (a *ArticleApi) EditArticle(c *gin.Context) {
	var r request.ModifyArticleBase
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
	var article = &models.Article{Title: r.Title, Content: r.Content, Excerpt: r.Excerpt, CategoryID: r.CategoryID, Status: r.Status, CoverImage: r.CoverImage, UserID: r.UserID}

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
