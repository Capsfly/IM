package utils

import (
	"awesomeProject/dao"
	"awesomeProject/entity"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	viper.SetConfigName("app") //读取配置文件
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("初始化配置失败,err=" + err.Error())
	}
}

func InitMysql() {
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	temp_db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{
		Logger: logger})
	if err != nil {
		panic("连接mysql 失败,err==" + err.Error())
	}
	dao.DB = temp_db
	dao.DB.AutoMigrate(&entity.User{})
}

func InitRedis() {
	dao.RDS = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})

}
