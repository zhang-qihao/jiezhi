package company

import (
	"Demo/conf"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func CreateID(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func JobEditor(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

		cid := context.PostForm("cid")
		jid := context.PostForm("jid")
		fmt.Println("cid:=", cid, "jid:=", jid)

		var jobinfonew conf.JobInfo
		DB.Where("cid = ? AND jid = ?", cid, jid).Find(&jobinfonew)

		context.JSON(http.StatusOK, gin.H{
			"title":       jobinfonew.Title,
			"salary":      jobinfonew.Salary,
			"address":     jobinfonew.Province,
			"experience":  jobinfonew.Experience,
			"education":   jobinfonew.Education,
			"nature":      jobinfonew.Nature,
			"fulladdress": jobinfonew.Address,
			"describe":    jobinfonew.Describe,
			"request":     jobinfonew.Require,
		})

		//获取按钮的类型，来定义所触发的函数
		ButtonType := context.PostForm("ButtonType")
		//获取ajax中所传的值
		positionName := context.PostForm("positionName")
		officeLocation := context.PostForm("officeLocation")
		address := context.PostForm("address")
		lowest := context.PostForm("lowest")
		highest := context.PostForm("highest")
		workExperience := context.PostForm("workExperience")
		academicRequirements := context.PostForm("academicRequirements")
		jobType := context.PostForm("jobType")

		sum := jobinfonew.Salary
		if lowest != "" || highest != "" {
			sum = lowest + "-" + highest
		}

		fmt.Println(ButtonType, positionName)

		jobinfonew.Title = positionName
		jobinfonew.Province = officeLocation
		jobinfonew.Address = address
		jobinfonew.Salary = sum
		jobinfonew.Experience = workExperience
		jobinfonew.Education = academicRequirements
		jobinfonew.Nature = jobType
		jobinfonew.ChangeTime = time.Now()

		DB.Model(&jobinfonew).Where("cid = ? AND jid = ?", cid, jid).Update(&jobinfonew)

		ButtonType1 := context.PostForm("ButtonType")
		fmt.Println(ButtonType1)
		responsibilities := context.PostForm("responsibilities")
		jobinfonew.Describe = responsibilities

		DB.Model(&jobinfonew).Where("cid = ? AND jid = ?", cid, jid).Update(&jobinfonew)

		ButtonType2 := context.PostForm("ButtonType")
		fmt.Println(ButtonType2)
		requirements := context.PostForm("requirements")
		jobinfonew.Require = requirements

		DB.Model(&jobinfonew).Where("cid = ? AND jid = ?", cid, jid).Update(&jobinfonew)

	}
}
