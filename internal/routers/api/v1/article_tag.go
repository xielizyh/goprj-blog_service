package v1

import (
	"context"

	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/service"
)

type ArticleTag struct{}

// NewArticle 新建文章标签
func NewArticleTag() ArticleTag {
	return ArticleTag{}
}

// func (at ArticleTag) Create(c *gin.Context) {
// 	param := service.CreateArticleTagRequest{}
// 	response := app.NewResponse(c)
// 	// valid, errs := app.BindAndValid(c, &param)
// 	// if !valid {
// 	// 	global.Logger.Errorf("app.BindAndValid error: %v", errs)
// 	// 	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 	// 	return
// 	// }
// 	svc := service.New(c.Request.Context())
// 	err := svc.CreateArticleTag(&param)
// 	if err != nil {
// 		global.Logger.Errorf("app.CreateArticleTag error: %v", err)
// 		response.ToErrorResponse(errcode.ErrorCreateArticleTagFail)
// 	}

// 	response.ToResponse(gin.H{})
// }

func (at ArticleTag) Create(tagID, articleID uint32) error {
	param := service.CreateArticleTagRequest{TagID: tagID, ArticleID: articleID}
	svc := service.New(context.TODO())
	err := svc.CreateArticleTag(&param)
	if err != nil {
		global.Logger.Errorf("app.CreateArticleTag error: %v", err)
		return err
	}
	return nil
}
