package global

import(
	"go_gin_blog/pkg/setting"
	"go_gin_blog/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)