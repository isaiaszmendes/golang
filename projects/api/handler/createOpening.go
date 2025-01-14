package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	logger.Infof("request received %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening %v", err.Error())
		ctx.JSON(500, gin.H{"error": "error creating opening"})
		return
	}
}

// TODO: Validar pq deu 200 ao remover campos do post no insomnia
