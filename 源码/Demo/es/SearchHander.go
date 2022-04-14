package es

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// step 1: 获取搜索的参数
		jobType := context.PostForm("index")
		var nature = "应届"
		if jobType == "littleday" {
			nature = "实习"
		}
		message := context.PostForm("message")
		pid := context.PostForm("pid")
		fmt.Printf("pid :=%s,head_select:=%s,message:=%s\n", pid, jobType, message)
		index := "job"
		typ := "job_info"
		// step 2: 调用ES引擎
		JobList := NatureMultiSearch(index, typ, nature, message)
		// step 3: 对返回的JobList数组做处理,得到非空的数据个数
		var count int
		for i := 0; i < len(JobList); i++ {
			if JobList[i].Jid != "" {
				count = count + 1
			}
		}
		context.JSON(http.StatusOK, gin.H{"joblist": JobList[0:count]})
	}

}
