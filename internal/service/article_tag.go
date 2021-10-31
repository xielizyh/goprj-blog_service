package service

type CreateArticleTagRequest struct {
	TagID     uint32 `form:"tag_id" binding:"required,gte=1"`
	ArticleID uint32 `form:"tag_id" binding:"required,gte=1"`
}

func (svc *Service) CreateArticleTag(param *CreateArticleTagRequest) error {
	return svc.dao.CreateArticleTag(
		param.TagID,
		param.ArticleID,
	)
}
