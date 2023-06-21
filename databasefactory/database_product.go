package databasefactory

type InterfaceDatabaseProduct interface {
	CreateDatabaseConnection()
	PingDatabase()
	CloseDatabaseConnection() error
}
