package protocol

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EmptyData struct{}

type RespData struct {
	Status
	Data interface{} `json:"data"`
}

type Status struct {
	Code      int    `json:"code"`
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

// RespError
func RespError(c *gin.Context, err error) {
	PrintLog(c, err)
	resp(c, EmptyData{}, 0, err)
}

func RespErrorWithUnauthorized(c *gin.Context, err error) {
	PrintLog(c, err)
	resp(c, EmptyData{}, http.StatusUnauthorized, err)
}

func RespErrorWithHttpCode(c *gin.Context, httpCode int, err error) {
	PrintLog(c, err)
	resp(c, EmptyData{}, httpCode, err)
}

// RespSuccess 成功
func RespSuccess(c *gin.Context, data ...interface{}) {
	var d interface{} = EmptyData{}

	if len(data) != 0 {
		d = data[0]
	}

	resp(c, d, 0, nil)
}

func resp(c *gin.Context, data interface{}, withHttpCode int, err error) {
	httpCode, msg, code := http.StatusBadRequest, err.Error(), 9999

	if withHttpCode == 0 {
		withHttpCode = httpCode
	}

	if data == nil {
		data = EmptyData{}
	}

	c.JSON(withHttpCode, RespData{
		Status: Status{
			Code:      code,
			Message:   msg,
			Timestamp: time.Now().Unix(),
		},
		Data: data,
	})
}
