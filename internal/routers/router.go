package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xielizyh/goprj-blog_service/internal/routers/api/v1"
)

// NewRouter 新建路由
func NewRouter() *gin.Engine {
	// 创建gin引擎实例
	r := gin.New()
	// 使用Logger中间件
	r.Use(gin.Logger())
	// 使用Recovery中间件
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()
	// 创建api/v1路由组
	apiv1 := r.Group("/api/v1")
	// 注册Handler到对应的路由规则
	{
		// 增
		apiv1.POST("/tags", tag.Create)
		// 删
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 改
		apiv1.PUT("/tags/:id", tag.Update)
		// 改
		apiv1.PATCH("/tags/:id/state", tag.Update)
		// 查
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.List)
		apiv1.GET("/articles", article.List)
	}

	return r
}