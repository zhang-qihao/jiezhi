package company

import (
	"Demo/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Manageapply(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		cid := context.PostForm("cid")
		fmt.Println("cid = ", cid)
		var personPost []conf.PersonPost
		DB.Where("cid=?", cid).Find(&personPost)
		var length = len(personPost)
		fmt.Println("所有数据长度为:", length)

		Post1 := make([]conf.PersonPost, 0)
		Post2 := make([]conf.PersonPost, 0)
		Post3 := make([]conf.PersonPost, 0)
		var status int
		for i := 0; i < length; i++ {
			status = personPost[i].Status
			switch status {
			case 3:
				Post1 = append(Post1, personPost[i])
			case 1:
				Post2 = append(Post2, personPost[i])
			case 2:
				Post3 = append(Post3, personPost[i])
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"rank0": personPost,
			"rank1": Post1,
			"rank2": Post2,
			"rank3": Post3,
		})

	}
}
