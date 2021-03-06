package global

import (
	"github.com/aajonas/blog-service/pkg/logger"
	"github.com/aajonas/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)