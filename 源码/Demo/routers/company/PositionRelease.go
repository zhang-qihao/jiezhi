package company

import (
	"Demo/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Position struct {
	Number string `json:"a"`
}

func PositionRelease(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		cid := context.PostForm("cid")
		fmt.Println(cid)
		var jobInfo []conf.JobInfo
		DB.Where("cid=?", cid).Find(&jobInfo)
		//fmt.Println("..........................................")
		//fmt.Println("企业id为：" + cid + "的全部岗位相关信息如下")
		//fmt.Println(rank0)
		context.JSON(http.StatusOK, gin.H{
			"number": len(jobInfo),
			"rank0":  jobInfo,
		})
	}
}
