package services

import (
	"errors"
	"sub/app/helpers/dbhelper"
	"sub/app/models"
)

// SaveMsg - to save hotel, room and rateplan object
func SaveMsg(msgData *models.MsgData) error {

	conn, err := dbhelper.GetConnByHost("")
	if err != nil {
		return err
	}

	if msgData == nil {
		return errors.New("No data received")
	}

	for _, offer := range msgData.Offers {

		err := conn.Create(&offer.Hotel).Error
		if err != nil {
			return err
		}
		err = conn.Create(&offer.Room).Error
		if err != nil {
			return err
		}

		err = conn.Create(&offer.RatePlan).Error
		if err != nil {
			return err
		}
	}

	return nil
}
