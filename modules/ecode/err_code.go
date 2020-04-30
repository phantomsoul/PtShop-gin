package ecode

var (
	OK		  = New(0, "OK")
	ServerErr = New(500, "server error")
)

var (
	SignCheckErr  = New(10000, "签名校验错误")
	RequestArgErr = New(10001, "参数有误")
)