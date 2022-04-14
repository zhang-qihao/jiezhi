package job

import (
	"Demo/conf"
	"Demo/connet"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func InfoOfJob(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		jid := context.PostForm("jid")
		pid := context.PostForm("pid")
		submit := context.PostForm("submit")
		fmt.Printf("[InfoOfJob] jid=%s,pid=%s, submit=%s\n", jid, pid, submit)
		var job conf.JobInfo
		var com conf.CompanyInfo
		DB.Where("jid=?", jid).Find(&job)
		DB.Where("cid=?", job.Cid).Find(&com)
		if submit != "申请职位" {
			context.JSON(http.StatusOK, gin.H{
				"title":            job.Title,
				"salary":           job.Salary,
				"address":          job.Address,
				"experience":       job.Experience,
				"education":        job.Education,
				"nature":           job.Nature,
				"companyName":      com.FullName,
				"companyType":      com.CompanyType,
				"companyScale":     com.Scale,
				"companyBrief":     com.Brief,
				"responsibilities": job.Describe,
				"requirements":     job.Require,
				"companyAddress":   com.Address,
			})
		} else {
			// 将申请信息录入该人的PersonPost表中
			// step1: 确定他是否申请过该职位
			var personPost conf.PersonPost
			DB.Where("pid=? and jid=?", pid, jid).Find(&personPost)
			if personPost.Pid != "" && personPost.Jid != "" {
				DB.Model(&personPost).Update("changeTime", time.Now())
				context.JSON(http.StatusOK, gin.H{"msg": "update"})
			} else {
				// step2 : 获取该人的个人信息
				var personInfoStr conf.PersonInfoStr
				DB.Where("pid=?", pid).Find(&personInfoStr)
				mode := 8
				basicField := connet.GetFieldName(conf.StructMap[mode])
				basicInfoStr := connet.Split(personInfoStr.BasicInfo)
				BasicStruct := connet.SetValueToStruct(basicField, basicInfoStr, mode)
				BasicInfo := *BasicStruct.(*conf.BasicInfoElement)
				fmt.Println("[InfoOfJob]: BasicInfo:=", BasicInfo)
				// 拼合得到需要写入的数据格式
				personPost.Pid = pid
				personPost.Jid = jid
				personPost.Cid = job.Cid
				personPost.Status = 0
				personPost.Title = job.Title
				personPost.Address = job.Address
				personPost.Experience = job.Experience
				personPost.Education = job.Education
				personPost.Salary = job.Salary
				personPost.Name = BasicInfo.Name
				// songchuang TODO: calculate Age
				personPost.Age = 69
				personPost.CreatTime = time.Now()
				personPost.ChangeTime = time.Now()
				DB.Save(&personPost)
				context.JSON(http.StatusOK, gin.H{"msg": "create"})
			}
		}

	}
}
