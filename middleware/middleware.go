package middleware

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"pt-gin/modules/cf"
	"time"
)

func RegisterCache(c *cf.CacheRedisConfig) gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(c.Addr, c.Pwd, time.Minute)
	return cache.Cache(&cacheStore)
}