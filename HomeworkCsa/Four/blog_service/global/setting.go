package global

import (
	"blog_service.com/m/pkg/logger"
	"blog_service.com/m/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
	JWTSetting *setting.JWTSettingS
	EmailSetting *setting.EmailSettingS
)
