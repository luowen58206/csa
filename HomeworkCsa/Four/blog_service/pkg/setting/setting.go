package setting

import "github.com/spf13/viper"

/*
	·针对读取配置文件的行为进行封装、便应用程序的使用
 */

type Setting struct {
	vp *viper.Viper
}

/*
	·定义了NewSetting方法、用于初始化本地的项目的配置的基础属性
	·viper 时允许设置多个配置路径的、这样可以尽可能的尝试解决路径查找的问题
	·也就是可以不断的调用 AddConfigPath方法
 */

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil,err
	}
	return &Setting{vp},nil
}
