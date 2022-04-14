package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
	*RedisConfig `ini:"redis"`
	*SmsCode  	 `ini:"SmsCode"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

// RedisConfig 数据库配置
type RedisConfig struct {
	DB   string `ini:"db"`
	Host string `ini:"host"`
	Port int    `ini:"port"`
}

type SmsCode struct {
	SignName 		string `ini:"signName"`
	TemplateCode 	string `ini:"templateCode"`
	AppKey  		string `ini:"appKey"`
	AppSecret 	 	string `ini:"appSecret"`
	RegionId		string `ini:"regionId"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
