package system

import (
	"blog/server/global"
	"blog/server/models"
	"blog/server/models/common/request"
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
	if id == 0 {
		return errors.New("article ID is required")
	}

	// 确保 art.ID 和传入的 id 一致，避免混淆
	art.ID = id

	// 使用结构体更新，更清晰
	result := global.GVA_DB.Model(&models.Article{}).
		Where("id = ?", id).
		Updates(models.Article{
			Title:      art.Title,
			Content:    art.Content,
			Excerpt:    art.Excerpt,
			CategoryID: art.CategoryID,
			Status:     art.Status,
			CoverImage: art.CoverImage,
		})

	if result.Error != nil {
		return result.Error
	}

	// 如果没找到记录，返回错误，而不是自动创建
	if result.RowsAffected == 0 {
		return errors.New("article not found")
	}

	return nil
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

func (as *ArticleService) ListArticle(uID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	db := global.GVA_DB.Model(models.Article{}).Where("user_id = ? and Status = 1", uID)
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var articles []models.Article
	if err = db.Order("Created_At desc").Scopes(info.Paginate()).Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, err
}
