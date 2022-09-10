package svc

import (
	"time"

	"github.com/zboyco/notify/notify/internal/config"
	"github.com/zboyco/notify/notify/internal/middleware"
	"github.com/zboyco/notify/notify/internal/task"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	Auth rest.Middleware //  auth middleware

	DB            *gorm.DB            // 数据库
	CronJobRunner *task.CronJobRunner // 定时任务
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ConnectDB(c)
	return &ServiceContext{
		Config:        c,
		Auth:          middleware.NewAuthMiddleware(c.Auth.Token).Handle,
		DB:            db,
		CronJobRunner: task.NewCronJobRunner(db),
	}
}

// Connect to the database
func ConnectDB(c config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.Postgres.Endpoint), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return db
}
