package main


import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"pt-gin/config"
	"pt-gin/middleware"
	"pt-gin/middleware/auth"
	"pt-gin/routes"
)

func initRouter(c *config.Config) {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*") // html模板

	if config.GetEnv().Debug {
		pprof.Register(router) // 性能分析工具
	}

	// router.Use(gin.Logger())
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义格式
		return fmt.Sprintf("[%s] - %s - %s --> %s \"%s %d %s\" %s \"%s\"\n",
			//param.TimeStamp.Format(time.RFC1123),
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(middleware.RegisterCache(c.CacheRedisConf))       // 全局cache
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth")) // 全局auth jwt

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
	})

	routes.RegisterApiRouter(c, router)

	go func() {
		router.Run(c.HTTPSrvConf.Addr)
	}()
}
