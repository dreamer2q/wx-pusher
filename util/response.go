package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errType int

const (
	OK errType = iota
	ErrBadPayload
	ErrMissingParam
	ErrTokenInvalid
	ErrInternal
)

var errMsg = map[errType]string{
	OK:              "success",
	ErrBadPayload:   "bad payload",
	ErrMissingParam: "missing param",
	ErrTokenInvalid: "token invalid",
	ErrInternal:     "internal error",
}

func (e errType) String() string {
	msg, ok := errMsg[e]
	if ok {
		return msg
	}
	return "unknown error"
}

type Response struct {
	ErrCode errType     `json:"err"`
	ErrMsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func doResponse(c *gin.Context, code errType, data interface{}) {
	var r = Response{
		ErrCode: code,
		ErrMsg:  code.String(),
		Data:    data,
	}
	c.JSON(http.StatusOK, r)
}

func MakeSuccess(c *gin.Context, obj ...interface{}) {
	if len(obj) == 1 {
		doResponse(c, OK, obj[0])
	} else {
		doResponse(c, OK, obj)
	}
}

func MakeFailure(c *gin.Context, code errType, obj ...interface{}) {
	if len(obj) == 1 {
		doResponse(c, code, obj[0])
	} else {
		doResponse(c, code, obj)
	}
}
