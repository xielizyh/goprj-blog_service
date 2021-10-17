package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 新建配置
func NewSetting() (*Setting, error) {
	// 创建viper
	vp := viper.New()
	// 设置Config名
	vp.SetConfigName("config")
	// 添加Config路径
	vp.AddConfigPath("configs/")
	// 设置Config类型
	vp.SetConfigType("yaml")
	// 或者直接配置
	// vp.SetConfigFile("configs/config.yaml")
	// 读取配置文件
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
