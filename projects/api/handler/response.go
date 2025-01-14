package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// o gin.H é um map[string]interface{}
// o H vem de Hash e nao de Header

// gin.H{} equivale a map[string]interface{}
// posso inclusive usar o map[string]interface{} no lugar do gin.H{}

func sendError(ctx *gin.Context, code int, msg string) {
	// application/json; charset=utf-8
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(code, gin.H{
		"message":    msg,
		"isdetailed": false,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}, code ...int) {
	statusCode := http.StatusOK

	ctx.Header("Content-Type", "application/json; charset=utf-8")
	if len(code) > 0 {
		statusCode = code[0]
	}
	ctx.JSON(statusCode, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})

}
