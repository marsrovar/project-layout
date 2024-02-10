package ijwt

import (
	"iserver/utils/consts"
)

func GetJWTSecret() []byte {
	return consts.JWTSecret
}
