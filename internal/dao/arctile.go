package dao

import "github.com/xielizyh/goprj-blog_service/internal/model"

func (d *Dao) CreateArticle(title, desc, content, coverImageUrl, createdBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		Model:         &model.Model{CreatedBy: createdBy},
	}

	return article.Create(d.engine)
}
