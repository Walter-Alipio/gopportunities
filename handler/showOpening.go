package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Walter-Alipio/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, parsedId).Error; err != nil {
		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
