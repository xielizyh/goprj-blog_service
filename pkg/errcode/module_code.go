package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorGetArticleTagFail    = NewError(20030001, "获取文章标签列表失败")
	ErrorCreateArticleTagFail = NewError(20030002, "创建文章标签失败")
	ErrorUpdateArticleTagFail = NewError(20030003, "更新文章标签失败")
	ErrorDeleteArticleTagFail = NewError(20030004, "删除文章标签失败")
	ErrorCountArticleTagFail  = NewError(20030005, "统计文章标签失败")
)
