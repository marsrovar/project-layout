package middleware

import (
	"iserver/internal/ijwt"
	"iserver/repository/cache"
	"iserver/server/protocol"
	"iserver/utils/errs"

	"github.com/gin-gonic/gin"
)

func CheckJwtAuthorizationHeader(c *gin.Context) {
	jwtData, err := CheckJwtAuthorization(protocol.GetAuthorizationHeader(c))
	if err != nil {
		protocol.RespError(c, err)
		c.Abort()
		return
	}

	protocol.SetJwtAuth(c, jwtData)

	c.Next()
}

func CheckJwtAuthorization(token string) (*ijwt.Auth, error) {
	if token == "" {
		return nil, errs.ErrUnauthorized
	}

	jwtData, err := ijwt.JwtTokenParse(token)
	if err != nil {
		return nil, errs.ErrUnauthorized
	}

	if jwtData.IsExpired() {
		return nil, errs.ErrUnauthorized
	}

	lastToken, err := cache.GetLastLoginAccountJwtToken(jwtData.GetAccount())
	if err != nil {
		return nil, errs.ErrUnauthorized
	}

	if token != lastToken {
		return nil, errs.ErrUnauthorized
	}

	return jwtData, nil
}
