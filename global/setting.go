package global

import (
	"github.com/xielizyh/goprj-blog_service/pkg/logger"
	"github.com/xielizyh/goprj-blog_service/pkg/setting"
)

// 全局区段配置
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
