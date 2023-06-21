package databasefactory

import (
	"fmt"

	"github.com/neat-vibes/go-clean-architecture/databasefactory/databases/mongodb"
	"github.com/neat-vibes/go-clean-architecture/databasefactory/databases/mysql"

	"github.com/neat-vibes/go-clean-architecture/config"
	"github.com/neat-vibes/go-clean-architecture/logger"
)

type InterfaceDatabaseFactory interface {
	CreateDatabase(config *config.Config) (InterfaceDatabaseProduct, error)
}

type DatabaseFactory struct {
}

func (factory DatabaseFactory) CreateDatabase(config *config.Config, l *logger.Logger) (InterfaceDatabaseProduct, error) {
	switch config.Database.Type {
	case "mysql":
		return &mysql.MySQLProduct{
			L:        l,
			DBConn:   nil,
			Err:      nil,
			Address:  config.Database.MySQL.Host,
			Port:     config.Database.MySQL.Port,
			User:     config.Database.MySQL.User,
			Password: config.Database.MySQL.Password,
			DBName:   config.Database.MySQL.DBName,
		}, nil
	case "mongodb":
		return &mongodb.MongodbProduct{
			DBConn: nil,
			Err:    nil,
		}, nil
	default:
		return nil, fmt.Errorf("unknown database type: %s", config.Database.Type)
	}
}
