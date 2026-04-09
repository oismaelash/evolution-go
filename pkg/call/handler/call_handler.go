package call_handler

import (
	"net/http"

	_ "github.com/EvolutionAPI/evolution-go/pkg/core"

	call_service "github.com/EvolutionAPI/evolution-go/pkg/call/service"
	instance_model "github.com/EvolutionAPI/evolution-go/pkg/instance/model"
	"github.com/gin-gonic/gin"
)

type CallHandler interface {
	RejectCall(ctx *gin.Context)
}

type callHandler struct {
	callService call_service.CallService
}

// Reject Call
// @Summary Reject Call
// @Description Reject Call
// @Tags Call
// @Accept json
// @Produce json
// @Param message body call_service.RejectCallStruct true "Call data"
// @Success 200 {object} core.CallRejectResponse
// @Failure 400 {object} core.Error400
// @Failure 401 {object} core.Error401
// @Failure 403 {object} core.Error403
// @Failure 404 {object} core.Error404
// @Failure 500 {object} core.Error500
// @Router /call/reject [post]
func (g *callHandler) RejectCall(ctx *gin.Context) {
	getInstance := ctx.MustGet("instance")

	instance, ok := getInstance.(*instance_model.Instance)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "instance not found"})
		return
	}

	var data *call_service.RejectCallStruct
	err := ctx.ShouldBindBodyWithJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = g.callService.RejectCall(data, instance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func NewCallHandler(
	callService call_service.CallService,
) CallHandler {
	return &callHandler{
		callService: callService,
	}
}
