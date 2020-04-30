package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"pt-gin/config"
	"pt-gin/controllers"
	"pt-gin/modules/log"
	"pt-gin/modules/sig"
	"strings"

	// _ "pt-gin/modules/schedule" // 定时任务
	"runtime"
)

func main() {
	currPath := getCurrentDirectory() //获取当前绝对路径
	log.InitLog(currPath+"/log", "debug")

	// 加载ini配置文件
	if err := config.Load(); err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化服务
	controllers.Init(config.Conf)

	// 初始化路由
	initRouter(config.Conf)

	sig.Run(func() {
		log.Info("Go-Gin PtShop Server Exit ^_^")
	})
}


func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
