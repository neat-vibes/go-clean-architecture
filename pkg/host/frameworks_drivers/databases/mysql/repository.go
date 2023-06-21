package mysql

import (
	"context"

	"github.com/neat-vibes/go-clean-architecture/databasefactory"
	"github.com/neat-vibes/go-clean-architecture/entities"
)

type mysqlHostRepo struct {
	Conn *databasefactory.InterfaceDatabaseProduct
}

func NewMysqlHostRepo(conn *databasefactory.InterfaceDatabaseProduct) entities.HostRepository {
	return &mysqlHostRepo{
		Conn: conn,
	}
}

func (m *mysqlHostRepo) Fetch(ctx context.Context, cursor string, num int64) ([]entities.Host, string, error) {

	return nil, "", nil
}
