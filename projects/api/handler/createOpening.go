package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaszmendes/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	logger.Infof("request received %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// unwrap
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening, http.StatusCreated)
}

// TODO: Validar pq deu 200 ao remover campos do post no insomnia
