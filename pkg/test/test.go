package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInfo(id int) (gin.H, error) {
	if id > 10 {
		return gin.H{"message": "test success"}, nil
	} else {
		return nil, fmt.Errorf("id must be greater than 10")
	}
}
