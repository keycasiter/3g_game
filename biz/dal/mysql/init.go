package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DataBase *gorm.DB

func InitMysql() {
	ctx := context.Background()
	config := conf.GetConfig()

	hlog.CtxInfof(ctx, "Mysql初始化...")
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.DbName,
	)
	var err error
	DataBase, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "Mysql初始化失败... err:%v", err)
		panic(err)
	} else {
		hlog.CtxInfof(ctx, "Mysql初始化成功...")
	}
}

func GetDataBase() *gorm.DB {
	return DataBase
}
