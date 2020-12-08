package dbhelper

import (
	"errors"
	"fmt"
	"sub/app/helpers/confighelper"
	"sub/app/helpers/loghelper"
	"sub/app/models"

	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

var dbConnections map[string]*gorm.DB // map to store db connections
var defaultHost string                // to store default db hostname

func init() {
	dbConnections = make(map[string]*gorm.DB)
}

// InitDatabases - create database connections and store in the dbConnections map
func InitDatabases() error {

	var dbErr error
	once.Do(func() {
		dbConfigs := models.DbConfigs{}

		// mapping the DB configs from dbcongif.yaml
		err := confighelper.Init(models.DbConfigPath, &dbConfigs)
		if err != nil {
			loghelper.LogError("error reading DB configs:", err)
			dbErr = err
		}

		// establishing connections for all the DBs in dbcongif.yaml
		for hostName, config := range dbConfigs.Hosts {

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Server, config.Port, config.Database)

			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // opening DB connection
			if err != nil {
				loghelper.LogError("error while connecting to DB:", err)
				dbErr = err
			}

			dbConnections[hostName] = db

			// Schema migration for Models
			dbConnections[hostName].AutoMigrate(&models.Hotel{}, &models.Room{}, &models.RatePlan{})

			if config.IsDefault {
				defaultHost = hostName
			}
		}
	})

	return dbErr
}

// GetConnByHost - to get connection from map of connections
func GetConnByHost(hostName string) (*gorm.DB, error) {

	if hostName == "" {
		connection, found := dbConnections[defaultHost]
		if !found {
			loghelper.LogError("Connection not found!")
			return nil, errors.New("Connection not found")
		}
		return connection, nil
	}

	connection, found := dbConnections[hostName]
	if !found {
		loghelper.LogError("Connection not found for ", hostName)
		return nil, errors.New("Connection not found for " + hostName)
	}
	return connection, nil

}
