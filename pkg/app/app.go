package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xielizyh/goprj-blog_service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

// NewResponse 新建响应
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

// ToResponse 正确响应
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToErrorResponse 错误响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
