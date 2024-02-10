package server

import (
	"iserver/server/protocol"
	"iserver/server/router"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func StartServer(port string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(stats.RequestStats())
	r.Use(gin.Recovery())

	corsConf := cors.DefaultConfig()
	corsConf.AllowHeaders = []string{"*"}
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	r.GET("/health", func(c *gin.Context) {
		protocol.RespSuccess(c)
	})

	router.Api(r)

	log.Println("server start :", port)
	if err := r.Run(port); err != nil {
		log.Println("server fail :", port)
	}
}
