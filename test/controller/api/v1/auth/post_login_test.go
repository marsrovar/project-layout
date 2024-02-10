package auth

import (
	"iserver/internal/ijwt"
	"iserver/repository/cache"
	authApi "iserver/server/controller/api/v1/auth"
	"iserver/test/iitest"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostLoginSuccess(t *testing.T) {
	iitest.TestingNewAccountDo(func() {
		caseTest := []struct {
			Account  string
			Password string
		}{
			{"admin", "adminXXX"},
			{"user", "userXXX"},
		}

		for _, v := range caseTest {
			resp := &authApi.RespPostLogin{}

			code, status, err := iitest.NewIClient().SetBody(authApi.ReqPostLogin{
				Account:  v.Account,
				Password: v.Password,
			}).Do(resp, authApi.PostLoginInfo.Method.String(), authApi.PostLoginInfo.Url)
			assert.Nil(t, err)
			assert.Equal(t, status, http.StatusOK)
			assert.Equal(t, code, 0)

			ijwtData, ierr := ijwt.JwtTokenParse(resp.Token)
			assert.Nil(t, ierr)
			assert.Equal(t, ijwtData.GetAccount(), v.Account)

			cacheToken, ierr := cache.GetLastLoginAccountJwtToken(v.Account)
			assert.Nil(t, ierr)
			assert.Equal(t, cacheToken, resp.Token)
		}
	})
}

func TestPostLoginFail(t *testing.T) {
	iitest.TestingNewAccountDo(func() {
		caseTest := []struct {
			Account  string
			Password string
		}{
			{"", ""},
			{"admi", "adminXXX"},
			{"admin", "admnXXX"},
			{"user1", "userXXX"},
		}

		for _, v := range caseTest {
			code, status, err := iitest.NewIClient().SetBody(authApi.ReqPostLogin{
				Account:  v.Account,
				Password: v.Password,
			}).Do(nil, authApi.PostLoginInfo.Method.String(), authApi.PostLoginInfo.Url)
			assert.Nil(t, err)
			assert.Equal(t, status, http.StatusBadRequest)
			assert.NotEqual(t, code, 0)
		}
	})
}
