package routes

import (
	structure "mes/models"

	"github.com/gin-gonic/gin"
)

func Error404Handler(context *gin.Context) {
	context.JSON(404, gin.H{"services": []structure.Service{}, "message": "Wrong routes used. Please read the docs on https://github.com/mervin16/Mauritius-Emergency-Services-Api-Go", "success": false})
}
