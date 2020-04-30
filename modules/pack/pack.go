package pack

import (
	"encoding/json"
	"pt-gin/modules/cf"
	"pt-gin/modules/ecode"
	"pt-gin/modules/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code 	  int			`json:"errno"`
	Msg		  string		`json:"msg"`
	Data      interface{}  	`json:"data"`
	TraceId   interface{}  	`json:"trace_id"`
}

func RespError(c *gin.Context, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*cf.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}
	eCode := ecode.Cause(err)
	resp := &Response{Code: eCode.Code(), Msg: eCode.Message(), Data: "", TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func RespSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*cf.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: 0, Msg: "", Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func Log(url string, req, res interface{}) {
	log.Debug("Request/Response --> ", zap.String("url", url), zap.Any("req", req), zap.Any("res", res))
}
