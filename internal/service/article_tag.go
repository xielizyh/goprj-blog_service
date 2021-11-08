package service

type CreateArticleTagRequest struct {
	TagID     uint32 `form:"tag_id" binding:"required,gte=1"`
	ArticleID uint32 `form:"tag_id" binding:"required,gte=1"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
}

func (svc *Service) CreateArticleTag(param *CreateArticleTagRequest) error {
	return svc.dao.CreateArticleTag(
		param.TagID,
		param.ArticleID,
		param.CreatedBy,
	)
}
