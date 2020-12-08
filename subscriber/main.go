package main

import (
	"fmt"
	"sub/app"
	"sub/app/helpers/confighelper"
	"sub/app/helpers/dbhelper"
	"sub/app/helpers/loghelper"
	"sub/app/models"
	"sub/app/utils"

	"github.com/streadway/amqp"
)

// init global things...
func init() {

	var err error

	// to set application configs from appConfig.yaml file
	err = setconfig()
	if err != nil {
		panic(err)
	}

	// Initialize logger
	initLogger()

	// Initialize database
	err = dbhelper.InitDatabases()
	utils.HandelException(err, true)
}

func main() {
	loghelper.LogInfo("Starting Subscriber")

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", models.AppConfig.AmqpConfig.Username, models.AppConfig.AmqpConfig.Password, models.AppConfig.AmqpConfig.Host, models.AppConfig.AmqpConfig.Port)
	conn, err := amqp.Dial(amqpURL) // connecting RabbitMQ
	utils.HandelException(err, true)

	defer conn.Close()

	// opening new channel
	ch, err := conn.Channel()
	utils.HandelException(err, true)

	defer ch.Close()

	// Initialize message queue consumer
	app.Init(ch)

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
