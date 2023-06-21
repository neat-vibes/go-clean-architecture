package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/neat-vibes/go-clean-architecture/logger"
)

type MySQLProduct struct {
	// logger
	L *logger.Logger

	// database connection
	DBConn *sql.DB
	Err    error

	// connect string
	Address  string
	Port     string
	User     string
	Password string
	DBName   string
	// connect other fields

}

func (p *MySQLProduct) CreateDatabaseConnection() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", p.User, p.Password, p.Address, p.Port, p.DBName)
	p.L.Debug("Connecting to database: " + connection)

	p.DBConn, p.Err = sql.Open("mysql", connection)
	if p.Err != nil {
		p.L.Panic("Error connecting to database: " + p.Err.Error())
		panic(p.Err.Error())
	}
}

func (p MySQLProduct) PingDatabase() {
	p.Err = p.DBConn.Ping()
	if p.Err != nil {
		p.L.Panic("Error pinging database: " + p.Err.Error())
		panic(p.Err.Error())
	}
}

func (p MySQLProduct) CloseDatabaseConnection() error {
	err := p.DBConn.Close()
	if err != nil {
		p.L.Panic("Error closing database connection: " + err.Error())
		return err
	}
	return nil
}
