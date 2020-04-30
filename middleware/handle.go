package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"pt-gin/modules/log"
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调试使用
		fmt.Println("-----Header-----")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		bodyStr, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("------Body------\r\n" + string(bodyStr))
		fmt.Println("----------------")

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyStr)) // 关键点:还原body内容
		defer func() {
			if err := recover(); err != nil {
				log.Sugar("Error->", err)
				var (
					errMsg		string
					mysqlError 	*mysql.MySQLError
					ok			bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError,
						gin.H{
						"code": 500,
						"msg":	"system error, " + errMsg,
						})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError,
						gin.H{
						"code": 500,
						"msg": 	"databases error, " + mysqlError.Error(),
						})
					return
				} else {
					c.JSON(http.StatusInternalServerError,
						gin.H{
						"code": 500,
						"msg":	"system error",
						})
					return
				}
			}
		}()
		c.Next()
	}
}
