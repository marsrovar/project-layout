package user

import (
	"iserver/obj"
	"iserver/repository/cache"
	"iserver/repository/models"
)

func GetUserInfo(account string) (obj.UserData, error) {
	if d, err := cache.GetUser(account); err == nil {
		return d, nil
	}

	userData, err := models.GetUser(account)
	if err != nil {
		return obj.UserData{}, err
	}

	cache.SetUser(account, userData)

	return userData, nil
}
