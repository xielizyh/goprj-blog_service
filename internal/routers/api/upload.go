package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/internal/service"
	"github.com/xielizyh/goprj-blog_service/pkg/app"
	"github.com/xielizyh/goprj-blog_service/pkg/convert"
	"github.com/xielizyh/goprj-blog_service/pkg/errcode"
	"github.com/xielizyh/goprj-blog_service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 读取file字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// 确定上传文件类型信息
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	// 保存文件
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	// 返回文件的展示地址
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
