package gorm

import (
	"context"
	"github.com/golang/protobuf/proto"
	mysql "github.com/oldbai555/driver-mysql"
	"github.com/oldbai555/gorm"
	"github.com/oldbai555/gorm/logger"
	"github.com/oldbai555/lbtool/log"
	"time"
)

const (
	autoMigrateOptKey   = "gorm:table_options"
	autoMigrateOptValue = "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"
)

var modelList []interface{}
var masterOrm *gorm.DB

func RegisterModel(vs ...interface{}) {
	modelList = append(modelList, vs...)
}

func InitGorm(dsn string) error {
	var err error
	masterOrm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: gorm.NamingStrategy{
			SingularTable: true,  // 是否单表，命名是否复数
			NoLowerCase:   false, // 是否关闭驼峰命名
		},

		NowFunc: func() int32 {
			return int32(time.Now().Unix())
		},

		PrepareStmt: true, // 预编译 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率

		Logger: NewOrmLog( //  日志配制
			log.GetLogger(),
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // 日志级别
			}),
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := masterOrm.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if len(modelList) > 0 {
		err := masterOrm.Set(autoMigrateOptKey, autoMigrateOptValue).AutoMigrate(modelList...)
		if err != nil {
			log.Errorf("err:%v", err)
			panic(err)
		}
	}

	return nil
}

func NewScope(ctx context.Context, model proto.Message) *Scope {
	if masterOrm == nil {
		panic("master orm is nil")
	}
	return &Scope{
		db: masterOrm.WithContext(ctx).Model(model),
	}
}
