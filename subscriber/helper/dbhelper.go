package helper

import (
	"encoding/json"
	"os"
	"sub/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	db = Connect() // initiate db connection
}

// Connect - to establish connection with DB
func Connect() *gorm.DB {
	dsn := "pubsub:pubsub@tcp(127.0.0.1:3306)/pubsub?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Check(err, "Error connecting to db")
	log.WithFields(log.Fields{"Client": "Subscriber"}).Info("Connected to Database")

	err = db.AutoMigrate(&models.Hotel{}, &models.Room{}, &models.RatePlan{})
	Check(err, "Schema migration failed")

	return db
}

// SaveMessage - save message to DB
func SaveMessage(msg []byte) bool {
	var msgData models.MsgData
	err := json.Unmarshal(msg, &msgData)
	Check(err, "Error while unmarshalling")
	if len(msgData.Offers) > 0 {
		db.Create(msgData.Offers[0].Hotel)
		db.Create(msgData.Offers[0].Room)
		db.Create(msgData.Offers[0].RatePlan)
	} else {
		log.WithFields(log.Fields{"Client": "Subscriber"}).Info("No data to save")
		return false
	}
	return true
}
