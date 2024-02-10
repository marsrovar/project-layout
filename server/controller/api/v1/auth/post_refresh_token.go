package authApi

import (
	"iserver/internal/ijwt"
	"iserver/repository/cache"
	"iserver/server/protocol"

	"github.com/gin-gonic/gin"
)

type RespPostRefreshToken struct {
	Token string `json:"token"`
}

var PostRefreshTokenInfo = protocol.ControllerInfo{
	Method:   protocol.Method.Post,
	Url:      "/api/v1/auth/refreshToken",
	Endpoint: "/refreshToken",
	Handler:  PostRefreshToken,
}

func PostRefreshToken(c *gin.Context) {
	jwtData, err := protocol.GetJwtAuth(c)
	if err != nil {
		protocol.RespError(c, err)
		return
	}

	data, err := postRefreshToken(jwtData)
	if err != nil {
		protocol.RespError(c, err)
		return
	}

	protocol.RespSuccess(c, data)
}

func postRefreshToken(jwtData *ijwt.Auth) (RespPostRefreshToken, error) {
	token, err := ijwt.BindAuth(jwtData.GetAccount())
	if err != nil {
		return RespPostRefreshToken{}, err
	}

	if err := cache.SetLastLoginAccountJwtToken(jwtData.GetAccount(), token); err != nil {
		return RespPostRefreshToken{}, err
	}

	return RespPostRefreshToken{token}, nil
}
