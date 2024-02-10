package authApi

import (
	"iserver/internal/ijwt"
	"iserver/repository/cache"
	"iserver/server/protocol"

	"github.com/gin-gonic/gin"
)

// request 定義
type ReqPostLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// response 定義
type RespPostLogin struct {
	Token string `json:"token"`
}

// router 設定
var PostLoginInfo = protocol.ControllerInfo{
	Method:   protocol.Method.Post,
	Url:      "/api/v1/auth/login",
	Endpoint: "/login",
	Handler:  PostLogin,
}

func PostLogin(c *gin.Context) {
	req := ReqPostLogin{}
	if err := protocol.ShouldBind(c, &req); err != nil {
		return
	}

	data, err := postLogin(req)
	if err != nil {
		protocol.RespError(c, err)
		return
	}

	protocol.RespSuccess(c, data)
}

func postLogin(req ReqPostLogin) (RespPostLogin, error) {
	// check and get user password

	// if user.Password != req.Password {
	// 	return RespPostLogin{}, ErrUnauthorized
	// }

	token, err := ijwt.BindAuth(req.Account)
	if err != nil {
		return RespPostLogin{}, err
	}

	if err := cache.SetLastLoginAccountJwtToken(req.Account, token); err != nil {
		return RespPostLogin{}, err
	}

	return RespPostLogin{token}, nil
}
