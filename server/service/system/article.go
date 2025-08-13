package system

import (
	"blog/server/global"
	"blog/server/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ArticleService struct{}

func (as *ArticleService) CreateArticle(art *models.Article) (err error) {
	err = global.GVA_DB.Create(art).Error
	return err
}

func (as *ArticleService) EditArticle(art *models.Article, id uint) (err error) {
	err = as.update(art, &id)
	return err
}

func (as *ArticleService) DeleteArticle(id uint) (err error) {
	result := global.GVA_DB.Delete(&models.Article{}, id)
	if result.Error != nil {
		err = fmt.Errorf("删除文章时出错")
		return
	}
	if result.RowsAffected == 0 {
		err = fmt.Errorf("文章不存在")
		return
	}
	return nil
}

func (as *ArticleService) update(art *models.Article, id *uint) (err error) {
	var articleID uint
	if id == nil {
		articleID = art.ID
	} else {
		articleID = *id
	}
	result := global.GVA_DB.Model(&art).Where("id = ?", articleID).Updates(map[string]interface{}{
		"title":       art.Title,
		"content":     art.Content,
		"except":      art.Excerpt,
		"category_id": art.CategoryID,
		"updated_at":  gorm.Expr("CURRENT_TIMESTAMP"),
		"status":      0,
	})
	if result.Error != nil {
		err = result.Error
	}
	if result.RowsAffected == 0 {
		err = global.GVA_DB.Create(&art).Error
		return err
	}
	return nil
}

func (as *ArticleService) SaveArticle(art *models.Article) (err error) {
	err = as.update(art, nil)
	return err
}

func (as *ArticleService) GetArticle(id uint) (art *models.Article, err error) {
	var article models.Article
	result := global.GVA_DB.First(&article, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("文章不存在")
		} else {
			err = fmt.Errorf("数据库错误")
		}
	}
	return &article, nil
}
