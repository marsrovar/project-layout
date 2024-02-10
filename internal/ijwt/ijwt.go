package ijwt

import (
	"errors"
	"iserver/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	jwtToken *jwt.Token
	jwtData  jwt.MapClaims
	exp      time.Time
}

type AuthData struct {
	//
}

func JwtTokenParse(token string) (*Auth, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("JwtTokenParse: bad signed method received")
		}

		return GetJWTSecret(), nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}

	jwtData, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	t, err := jwtData.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &Auth{jwtToken, jwtData, t.Time}, nil
}

func (auth *Auth) IsExpired() bool {
	if auth.exp.Before(time.Now()) {
		return true
	}

	return false
}

func (auth *Auth) GetAccount() string {
	account := auth.jwtData["account"]

	if v, ok := account.(string); ok {
		return v
	}

	return ""
}

func (auth *Auth) GetAuthData() (AuthData, error) {
	data := auth.jwtData["data"]

	authData := AuthData{}

	if err := utils.Json.Unmarshal([]byte(data.(string)), &authData); err != nil {
		return AuthData{}, err
	}

	return authData, nil
}

func BindAuth(account string) (string, error) {
	b, _ := utils.Json.Marshal(AuthData{})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account": account,
		"data":    string(b),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(GetJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
