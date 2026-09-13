package handler

import (
	"net/http"

	"github.com/Walter-Alipio/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			"error linting openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
