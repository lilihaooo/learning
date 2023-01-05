package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"learning2.0/database"
	"learning2.0/routers"
	"learning2.0/settings"
)

var (
	cfg = pflag.StringP("config", "d", "conf/config_develop.yaml", "input the config")
)

func init() {
	pflag.Parse()
	//加载配置文件
	if err := settings.InitConfig(*cfg); err != nil {
		panic(err)
	}
	//初始化数据库
	database.InitMysql()
	//database.InitRedis()
}

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	r.Run(viper.GetString("addr"))
}
