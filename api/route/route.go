package route

import (
	"time"

	"github.com/neat-vibes/go-clean-architecture/databasefactory"

	"github.com/neat-vibes/go-clean-architecture/config"

	"github.com/gin-gonic/gin"
)

func SetupRoute(config *config.Config, timeout time.Duration, db *databasefactory.InterfaceDatabaseProduct, gin *gin.Engine) {
	v1 := gin.Group("v1")
	{
		NewHostRouter(config, timeout, db, v1)
	}
}
