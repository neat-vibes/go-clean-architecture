package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neat-vibes/go-clean-architecture/entities"
)

type HostHandler struct {
	HostUsecase entities.HostUsecase
}

func (h *HostHandler) Fetch(c *gin.Context) {
	numStr := c.Query("num")
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		num = 0
	}
	cursor := c.Query("cursor")
	ctx := c.Request.Context()
	listAr, nextCursor, err := h.HostUsecase.Fetch(ctx, cursor, num)
	if err != nil {
		return
	}
	c.Request.Response.Header.Set("X-Cursor", nextCursor)
	return c.JSON(200, listAr)
}

func NewHostHandler(g *gin.Engine, usecaseHost entities.HostUsecase) {
	handler := &HostHandler{
		HostUsecase: usecaseHost,
	}

	v1 := g.Group("/v1")
	{
		v1.GET("/hosts", handler.Fetch)
	}
}
