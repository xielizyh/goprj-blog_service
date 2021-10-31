package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/service"
	"github.com/xielizyh/goprj-blog_service/pkg/app"
	"github.com/xielizyh/goprj-blog_service/pkg/errcode"
)

type Article struct{}

// NewArticle 新建文章
func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)  {}
func (a Article) List(c *gin.Context) {}

// @Summary 新增文章
// @Produce  json
// @Param tag_id body uint32 true "标签"
// @Param title body string true "标题" minlength(2) maxlength(100)
// @Param desc body string true "描述" minlength(2) maxlength(255)
// @Param content body string true "内容" minlength(2) maxlength(4294967295)
// @Param cover_image_url body string true "图片地址"
// @Param created_by body string true "创建者" minlength(2) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("app.CreateArticle error: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	// 更新文章标签
	tagID := param.TagID
	// TODO 查询数据库文章ID
	var articleID uint32 = 1
	at := NewArticleTag()
	err = at.Create(tagID, articleID)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticleTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
