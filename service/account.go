package service

import (
	"account_service/model"
	"log"

	"gorm.io/gorm"
)

func GetUserAppID(tx *gorm.DB, uid string) (appID string, err error) {
	usr := model.User{}
	tx.Table("t_b_user").
		Select("app_id").
		Where("uid", uid).
		Find(&usr)
	appID = usr.AppID
	err = tx.Error
	return
}

func UpdateUser(tx *gorm.DB, uid string, appID string) (af int64, err error) {
	usr := model.User{
		AppID: appID,
	}
	log.Println(26, uid, appID, usr.AppID, usr.Table())
	tx.Table("t_b_user").
		Where("uid", uid).
		Updates(usr)
	af = tx.RowsAffected
	err = tx.Error
	return
}
