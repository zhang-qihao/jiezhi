package application

import (
	"Demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func PersonPostPage(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		pid := context.PostForm("pid")
		println("[PersonPostPage]:", pid)
		personPost := make([]conf.PersonPost, 0)
		Post1 := make([]conf.PersonPost, 0)
		Post2 := make([]conf.PersonPost, 0)
		Post3 := make([]conf.PersonPost, 0)
		Post4 := make([]conf.PersonPost, 0)
		DB.Where("pid=?", pid).Find(&personPost)
		length := len(personPost)
		var status int
		for i := 0; i < length; i++ {
			status = personPost[i].Status
			switch status {
			// 1: 未查看
			case 0:
				Post1 = append(Post1, personPost[i])
			// 待定
			case 1:
				Post2 = append(Post2, personPost[i])
			// 邀请面试
			case 2:
				Post3 = append(Post3, personPost[i])
			// 拒绝
			case 3:
				Post4 = append(Post4, personPost[i])
			}
		}
		//fmt.Printf("[PersonPostPage],全部个数为:%d,未查看个数为:%d,待定个数为:%d,面试个数为:%d,拒绝个数为:%d\n", len(personPost), len(Post1), len(Post2), len(Post3), len(Post4))

		context.JSON(http.StatusOK, gin.H{
			"rank0": personPost,
			"rank1": Post1,
			"rank2": Post2,
			"rank3": Post3,
			"rank4": Post4,
		})
	}
}
