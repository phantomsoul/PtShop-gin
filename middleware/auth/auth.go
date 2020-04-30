package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
	"net/http"
	"pt-gin/conn/redis"
	"pt-gin/dto"
	"pt-gin/middleware/auth/drivers"
	"pt-gin/modules/cf"
	"pt-gin/modules/log"
	"pt-gin/modules/util"
)

const (
	JwtAuthDriverKey = "jwt"
	AclDriverKey     = "acl"
)

var driverList = map[string]Auth{
	JwtAuthDriverKey: drivers.NewJwtAuthDriver(),
}

var r *redigo.Pool

type Acl struct {
	c     *cf.RedisConfig
	redis *redigo.Pool
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) dto.JwtUserInfo
	CreateToken(userInfo map[string]interface{}) interface{}
}

func New(conf *cf.RedisConfig) *Acl {
	ret := &Acl{
		c:     conf,
		redis: redis.NewPool(conf),
	}
	r = ret.redis
	return ret
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, GenerateAuthDriver(authKey))
		c.Next()
	}
}

func Middleware(authKey string, aclFlag bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !GenerateAuthDriver(authKey).Check(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "认证失败，请重新登录",
			})
			c.Abort()
			return
		}

		// 开始权限检查
		if aclFlag {
			tInfo := drivers.NewJwtAuthDriver().User(c)

			// 检查Token信息是否完整
			if tInfo.RoleID == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "认证失败，Token信息不完整",
				})
				c.Abort()
				return
			}

			// 管理员跳过检查
			if tInfo.IsAdmin == 1 {
				c.Next()
			} else {
				hasPower := CheckRole(tInfo.CtmID, tInfo.RoleID, util.GetPowerStr(fmt.Sprintf("%s", c.Request.URL)))
				if hasPower {
					c.Next()
				} else {
					log.Warn(fmt.Sprintf("***** User %s has no power for %s *****", tInfo.UserName, c.Request.URL))
					c.JSON(http.StatusForbidden, gin.H{
						"code": 403,
						"msg":  "无权限访问",
					})
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) Auth {
	return driverList[string]
}

func CheckRole(ctmId int64, roleId string, aclName string) bool {
	key := fmt.Sprintf("set:acl:role:menu:%d:%s", ctmId, roleId)
	c := r.Get()
	res, err := redigo.Bool(c.Do("SISMEMBER", key, aclName))
	defer c.Close()
	if err != nil {
		return false
	}
	return res
}
