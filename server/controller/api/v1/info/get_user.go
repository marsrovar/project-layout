package configApi

import (
	"iserver/internal/user"
	"iserver/server/protocol"

	"github.com/gin-gonic/gin"
)

type ReqGetUser struct {
	Account string `form:"account"`
}

type RespGetUser struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	EMail   string `json:"email"`
}

var GetUserInfo = protocol.ControllerInfo{
	Method:   protocol.Method.Get,
	Url:      "/api/v1/info/user",
	Endpoint: "/user",
	Handler:  GetUser,
}

func GetUser(c *gin.Context) {
	req := ReqGetUser{}
	if err := protocol.ShouldBind(c, &req); err != nil {
		return
	}

	userData, err := user.GetUserInfo(req.Account)
	if err != nil {
		protocol.RespError(c, err)
		return
	}

	protocol.RespSuccess(c, RespGetUser{
		Account: userData.Account,
		Name:    userData.Name,
		EMail:   userData.EMail,
	})
}
