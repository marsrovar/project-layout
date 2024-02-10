package iitest

import (
	"iserver/repository/models"
)

func insertAccount() error {
	_, err := models.IDB.Exec("INSERT INTO user ")

	return err
}
