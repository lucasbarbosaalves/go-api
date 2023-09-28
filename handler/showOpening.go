package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbarbosaalves/go-api/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show job opening
// @Tags Openings
// @Accept  json
// @Produce  json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not foud")
		return
	}
	sendSucess(ctx, "show-opening", opening)
}
