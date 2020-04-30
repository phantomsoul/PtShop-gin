package service

import (
	"database/sql"
	goRedis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"pt-gin/config"
	db "pt-gin/conn/mysql"
	"pt-gin/conn/redis"
	"pt-gin/dao"
)

type Service struct {
	c       *config.Config
	d       *dao.Dao
	db      *sql.DB
	orm		*gorm.DB
	goRedis *goRedis.Client //初始化一个go-redis单连接
}

func New(c *config.Config) (s *Service) {
	orm := db.NewDB(c.MysqlConf)
	s = &Service{
		c:       c,
		d:       dao.New(c),
		orm:	 orm,
		goRedis: redis.NewConn(c.RedisConf),
	}
	return s
}
