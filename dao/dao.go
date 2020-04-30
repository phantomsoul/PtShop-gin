package dao

import (
	"github.com/aliyun/aliyun-sts-go-sdk/sts"
	_ "github.com/go-sql-driver/mysql"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"pt-gin/config"
	db "pt-gin/conn/mysql"
	"pt-gin/conn/redis"
)

type Dao struct {
	c         *config.Config
	orm 	  *gorm.DB
	redis     *redigo.Pool
	taskRedis *redigo.Pool
	stsClient *sts.Client
}

func New(c *config.Config) (d *Dao) {
	orm := db.NewDB(c.MysqlConf)
	d = &Dao{
		c:         c,
		orm:       orm,
		redis:     redis.NewPool(c.RedisConf),
		taskRedis: redis.NewTaskPool(c.TaskRedisConf),
		stsClient: sts.NewClient(c.OSSConf.AccessKeyID, c.OSSConf.AccessKeySecret, c.OSSConf.RoleArn, c.OSSConf.SessionName),
	}
	return d
}

func (d *Dao) Close() {
	_ = d.redis.Close()
	_ = d.taskRedis.Close()
}
