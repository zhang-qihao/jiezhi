package person

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ModifyResume() gin.HandlerFunc {
	return func(context *gin.Context) {
		modified1 :=context.PostForm("modify_1")
		fmt.Println("modified_1: ",modified1)
	}
}
