package server_handler

import (
	_ "github.com/EvolutionAPI/evolution-go/pkg/core"
	"github.com/gin-gonic/gin"
)

type ServerHandler interface {
	ServerOk(ctx *gin.Context)
}

type serverHandler struct {
}

// @Summary Check server status
// @Description Returns the server status to verify it is running
// @Tags Server
// @Produce json
// @Success 200 {object} core.ServerOkResponse
// @Failure 400 {object} core.Error400
// @Failure 401 {object} core.Error401
// @Failure 403 {object} core.Error403
// @Failure 404 {object} core.Error404
// @Failure 500 {object} core.Error500
// @Router /server/ok [get]
func (s *serverHandler) ServerOk(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func NewServerHandler() ServerHandler {
	return &serverHandler{}
}
