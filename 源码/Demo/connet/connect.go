package connet

import (
	"Demo/conf"
	"Demo/setting"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB       *gorm.DB
	PoolConn redis.Conn
)

func InitMySQL(cfg *setting.MySQLConfig) (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//模型Model绑定数据库的private_account、business_account表
	DB.AutoMigrate(
		&conf.PrivateAccount{},
		&conf.BusinessAccount{},
		&conf.PersonInfoStr{},
		&conf.JobInfo{},
		&conf.CompanyInfo{},
		&conf.PersonPost{},
	)
	return DB, DB.DB().Ping()
}

func InitRedis(cfg *setting.RedisConfig) (PoolConn redis.Conn, err error) {
	dsn := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	PoolConn, err = redis.Dial("tcp", dsn)
	return
}
