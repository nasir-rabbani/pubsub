package main

import (
	"fmt"
	"pub/app"
	"pub/app/helpers/confighelper"
	"pub/app/helpers/loghelper"
	"pub/app/models"
	"pub/app/utils"

	"github.com/streadway/amqp"
)

// init all stuffs required to project here...
func init() {

	var err error

	// to set application configs from appConfig.yaml file
	err = setconfig()
	if err != nil {
		panic(err)
	}

	// Initialize logger
	initLogger()
}

func main() {

	loghelper.LogInfo("Starting Publisher")
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", models.AppConfig.AmqpConfig.Username, models.AppConfig.AmqpConfig.Password, models.AppConfig.AmqpConfig.Host, models.AppConfig.AmqpConfig.Port)
	conn, err := amqp.Dial(amqpURL)
	utils.HandelException(err, true)

	defer conn.Close()

	ch, err := conn.Channel()
	utils.HandelException(err, true)

	defer ch.Close()

	// Initialize message queue publisher
	app.Publish(ch)

}

// setconfig - map project config provided in configs/appconfig.yaml file
func setconfig() error {
	err := confighelper.Init(models.ConfigFilePath, &models.AppConfig)
	if err != nil {
		return err
	}

	return nil
}

// Initialize Logger
func initLogger() {
	loghelper.Init()
}
