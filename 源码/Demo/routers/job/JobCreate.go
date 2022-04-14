package job

import (
	"Demo/conf"
	"Demo/routers/Login"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func JobCreate(DB *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		cid := c.PostForm("cid")
		println("[JobCreate] cid:=", cid)
		var company conf.CompanyInfo
		DB.Where("cid=?", cid).Find(&company)
		//获取按钮的类型，来定义所触发的函数
		ButtonType := c.PostForm("ButtonType")
		//获取ajax中所传的值
		c.JSON(200, gin.H{"msg": "上传成功"})
		title := c.PostForm("positionName")
		address := c.PostForm("officeLocation")
		nature := c.PostForm("jobType")
		var jid = Login.CreateID(title + address + nature)
		if ButtonType == "SaveJob" {
			var job = conf.JobInfo{
				Jid:        jid,
				Cid:        cid,
				Title:      title,
				Province:   address,
				Company:    company.ShortName,
				Address:    c.PostForm("address"),
				Salary:     c.PostForm("lowest") + "-" + c.PostForm("highest"),
				Experience: c.PostForm("workExperience"),
				Education:  c.PostForm("academicRequirements"),
				Nature:     nature,
				Require:    c.PostForm("responsibilities"),
				Describe:   c.PostForm("requirements"),
				CreatTime:  time.Now(),
				ChangeTime: time.Now(),
			}
			DB.Create(&job)
		}
	}
	//RegisteredCompany := &conf.CompanyInfo{}
}
