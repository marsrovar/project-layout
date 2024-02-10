package auth

import (
	"iserver/internal/ijwt"
	authApi "iserver/server/controller/api/v1/auth"
	"iserver/test/iitest"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func loginAndRefreshToken(isWithoutToken bool) (string, string, int, int, error) {
	respToken := &authApi.RespPostLogin{}

	code, status, er := iitest.NewIClient().SetBody(authApi.ReqPostLogin{
		Account:  "user",
		Password: "userXXX",
	}).Do(respToken, authApi.PostLoginInfo.Method.String(), authApi.PostLoginInfo.Url)
	if er != nil || status != http.StatusOK || code != 0 {
		panic(status)
	}

	time.Sleep(time.Second)

	nc := iitest.NewIClient()

	if !isWithoutToken {
		nc.SetHeader(map[string]string{
			"Authorization": respToken.Token,
		})
	}

	resp := &authApi.RespPostRefreshToken{}

	code, status, er = nc.Do(resp, authApi.PostRefreshTokenInfo.Method.String(), authApi.PostRefreshTokenInfo.Url)

	return respToken.Token, resp.Token, code, status, er
}

func TestPostRefreshTokenSuccess(t *testing.T) {
	iitest.TestingNewAccountDo(func() {
		respToken, resp, code, status, er := loginAndRefreshToken(false)
		assert.Nil(t, er)
		assert.Equal(t, status, http.StatusOK)
		assert.Equal(t, code, 0)

		user1, err := ijwt.JwtTokenParse(respToken)
		assert.Nil(t, err)

		user2, err := ijwt.JwtTokenParse(resp)
		assert.Nil(t, err)

		assert.Equal(t, user1.GetAccount(), user2.GetAccount())
		assert.NotEqual(t, respToken, resp)
	})
}

func TestPostRefreshTokenUnauthorized(t *testing.T) {
	iitest.TestingNewAccountDo(func() {
		_, _, code, status, er := loginAndRefreshToken(true)
		assert.Nil(t, er)
		assert.Equal(t, status, http.StatusUnauthorized)
		assert.NotEqual(t, code, 0)
	})
}
