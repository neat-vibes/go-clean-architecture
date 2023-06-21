package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/neat-vibes/go-clean-architecture/config"
	"github.com/neat-vibes/go-clean-architecture/entities"
)

type HostController struct {
	HostUsecase entities.HostUsecase
	Config      *config.Config
}

func (p *HostController) Host(c *gin.Context) {

}
