package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Cfg struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

var cfg = Cfg{
	Host:     "rm-2vc208dd40g94459wqo.mysql.cn-chengdu.rds.aliyuncs.com",
	User:     "lihaoselfhigh",
	Password: "lihao@123",
	Port:     3306,
	DB:       "learning2.0",
}

func main() {
	// 连接数据库
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../query",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)
	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)
	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)
	g.Execute()
}
