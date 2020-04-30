package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

var (
	errMap map[int]string
)

func New(code int, msg string) Ecode {
	if errMap == nil {
		errMap = make(map[int]string)
	}
	if _, ok := errMap[code]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", code))
	}
	errMap[code] = msg
	return Ecode(code)
}

type Codes interface {
	Error() string
	Code()	int
	Message() string
}

type Ecode int

func (e Ecode) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Ecode) Code() int {
	return int(e)
}

func (e Ecode) Message() string {
	if m, ok := errMap[e.Code()]; ok {
		return m
	}
	return e.Error()
}

func String(e string) Ecode {
	if e == "" {
		return OK
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Ecode(i)
}

func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}