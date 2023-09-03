package handlers

import "github.com/gin-gonic/gin"

func Build(r *gin.Engine) {
	NewUserHandler().Build(r)
}
