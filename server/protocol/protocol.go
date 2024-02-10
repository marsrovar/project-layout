package protocol

import (
	"iserver/internal/ijwt"
	"iserver/utils/errs"
	"log"

	"github.com/gin-gonic/gin"
)

var Method = struct {
	Get  MethodType
	Post MethodType
	Put  MethodType
	Del  MethodType
}{
	Get:  "Get",
	Post: "Post",
	Put:  "Put",
	Del:  "Del",
}

type MethodType string

func (mt MethodType) String() string {
	return string(mt)
}

type ControllerInfo struct {
	Method   MethodType
	Url      string // 完整的 url
	Endpoint string // 最終節點
	Handler  gin.HandlerFunc
}

func (cInfo ControllerInfo) Router(r *gin.RouterGroup) {
	switch cInfo.Method {
	case Method.Get:
		r.GET(cInfo.Endpoint, cInfo.Handler)
	case Method.Post:
		r.POST(cInfo.Endpoint, cInfo.Handler)
	case Method.Put:
		r.PUT(cInfo.Endpoint, cInfo.Handler)
	case Method.Del:
		r.DELETE(cInfo.Endpoint, cInfo.Handler)
	}
}

func ShouldBind(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		RespError(c, errs.ErrInputInvalidParameter)
		return errs.ErrInputInvalidParameter
	}

	return nil
}

func PrintLog(c *gin.Context, err error) {
	log.Println(c.Request.Method + "," + c.Request.URL.Path + ": " + err.Error())
}

func GetAuthorizationHeader(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func SetJwtAuth(c *gin.Context, jwtAuth *ijwt.Auth) {
	c.Set("jwtAuth", jwtAuth)
}

func GetJwtAuth(c *gin.Context) (*ijwt.Auth, error) {
	data, ok := c.Get("jwtAuth")
	if !ok {
		return nil, errs.ErrSystem
	}

	return data.(*ijwt.Auth), nil
}
