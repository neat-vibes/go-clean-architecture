package route

import (
	"time"

	"github.com/neat-vibes/go-clean-architecture/api/controller"
	"github.com/neat-vibes/go-clean-architecture/config"
	"github.com/neat-vibes/go-clean-architecture/databasefactory"
	"github.com/neat-vibes/go-clean-architecture/pkg/host/frameworks_drivers/databases/mysql"
	"github.com/neat-vibes/go-clean-architecture/pkg/host/usecases"

	"github.com/gin-gonic/gin"
)

func NewHostRouter(config *config.Config, timeout time.Duration, db *databasefactory.InterfaceDatabaseProduct, group *gin.RouterGroup) {

	hostRepo := mysql.NewMysqlHostRepo(db)
	hc := controller.HostController{
		HostUsecase: usecases.NewHostUsecase(hostRepo, timeout),
		Config:      config,
	}
	group.GET("/host", hc.Host)
}
