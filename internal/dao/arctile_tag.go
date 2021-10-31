package dao

import "github.com/xielizyh/goprj-blog_service/internal/model"

func (d *Dao) CreateArticleTag(tagID, articleID uint32) error {
	articleTag := model.ArticleTag{
		TagID:     tagID,
		ArticleID: articleID,
	}

	return articleTag.Create(d.engine)
}
