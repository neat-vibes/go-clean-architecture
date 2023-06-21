package mongodb

import "database/sql"

type MongodbProduct struct {
	DBConn *sql.DB
	Err    error
}

func (p *MongodbProduct) CreateDatabaseConnection() {

}

func (p MongodbProduct) PingDatabase() {

}

func (p MongodbProduct) CloseDatabaseConnection() error {
	return nil
}
