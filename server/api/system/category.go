package system

import (
	"blog/server/models"
	common_response "blog/server/models/common/response"
	"blog/server/models/request"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type CategoryApi struct{}

func (ca *CategoryApi) CreateCategory(c *gin.Context) {
	var category request.CreateCategory
	err := c.ShouldBind(&category)
	if err != nil {
		common_response.FailWithMessage("", c)
		return
	}
	err = utils.Verify(category, utils.CategoryVerify) // category 结构体不是指针
	if err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	cate := models.Category{Name: category.Name, Description: category.Description}
	if err = categoryService.CreateCategory(&cate); err != nil {
		common_response.FailWithMessage(err.Error(), c)
		return
	}
	common_response.OkWithMessage("创建成功", c)
}
