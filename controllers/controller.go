package controllers

import (
	"pt-gin/config"
	"pt-gin/middleware/auth"
	"pt-gin/service"
)

var (
	srv *service.Service
	acl *auth.Acl
)

func Init(c *config.Config) {
	// 连接依赖服务
	srv = service.New(c)
	acl = auth.New(c.RedisConf)
}