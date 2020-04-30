package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"pt-gin/modules/cf"
	"pt-gin/modules/pack"
	"strings"
)

func IPAuthMiddleware() gin.HandlerFunc {
	var cfh *cf.HTTPServerConfig
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range GetAllowIpConf(cfh.AllowIp) {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			pack.RespError(c, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetAllowIpConf(key string) []string {
	keys := strings.Split(key, ",")
	if len(keys) < 2 {
		return nil
	}
	conf := keys
	return conf
}