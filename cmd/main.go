package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/neat-vibes/go-clean-architecture/api/route"
	"github.com/neat-vibes/go-clean-architecture/config"
	"github.com/neat-vibes/go-clean-architecture/logger"

	"github.com/neat-vibes/go-clean-architecture/databasefactory"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line arguments
	configFile := flag.String("config", "config.json", "Configuration file")
	flag.Parse()

	// Load configuration
	config, err := config.NewConfig(*configFile)
	if err != nil {
		panic(err)
	}
	config.PrintConfig()

	// Create logger
	logger, err := logger.NewLogger(config.GetLogLevel(), config.GetLogPath())
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// create database factory
	factoryDB := databasefactory.DatabaseFactory{}
	productDB, err := factoryDB.CreateDatabase(config, logger)
	if err != nil {
		panic(err)
	}
	productDB.CreateDatabaseConnection()
	productDB.PingDatabase()
	defer func() {
		err := productDB.CloseDatabaseConnection()
		if err != nil {
			panic(err)
		}
	}()

	r := gin.New()
	r.Use(gin.Logger())

	// Setup route
	route.SetupRoute(config, time.Duration(config.GetContextTimeout()), &productDB, r)

	serverEntryPoint := fmt.Sprintf("%s:%s", config.GetServerAddress(), config.GetServerPort())
	r.Run(serverEntryPoint)
}
