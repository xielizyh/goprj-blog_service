package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/service"
	"github.com/xielizyh/goprj-blog_service/pkg/app"
	"github.com/xielizyh/goprj-blog_service/pkg/convert"
	"github.com/xielizyh/goprj-blog_service/pkg/errcode"
)

type Tag struct{}

// NewTag 新建标签
func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.TagListRequest{}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 正常回应
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag error: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList error: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.CreateTagRequest{}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag error: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag error: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag error: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
}
