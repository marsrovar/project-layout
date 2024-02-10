package router

import (
	authApi "iserver/server/controller/api/v1/auth"
	infoApi "iserver/server/controller/api/v1/info"
	"iserver/server/middleware"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")

	authGroup := apiGroup.Group("/auth")
	{
		authApi.PostLoginInfo.Router(authGroup)

		authGroup.Use(middleware.CheckJwtAuthorizationHeader)
		authApi.PostRefreshTokenInfo.Router(authGroup)
	}

	infoGroup := apiGroup.Group("/info")
	infoGroup.Use(middleware.CheckJwtAuthorizationHeader)
	{
		infoApi.GetUserInfo.Router(infoGroup)
	}
}
