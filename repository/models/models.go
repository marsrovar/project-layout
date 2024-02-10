package models

import (
	"database/sql"
	"iserver/obj"
)

var IDB *sql.DB

func GetUser(account string) (obj.UserData, error) {
	data := obj.UserData{}

	if err := IDB.QueryRow("SELECT  id, account, name, email FROM users WHERE account=?", account).
		Scan(&data.Id, &data.Account, &data.Name, &data.EMail); err != nil {
		return obj.UserData{}, err
	}

	return data, nil
}
