package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbarbosaalves/go-api/schemas"
)

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
