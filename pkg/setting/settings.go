package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// NewSetting 新建配置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigFile("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
