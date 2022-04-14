package company

import (
	"Demo/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Companyhomepage(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		exit := context.PostForm("exit")
		if exit != "" {
			context.SetCookie("authorized", "false", 3600, "/", "127.0.0.1:8080", false, true)
		}
		cid := context.PostForm("cid")
		fmt.Println("cid = ", cid)
		var company conf.CompanyInfo
		DB.Where("cid = ?", cid).Find(&company)
		context.JSON(http.StatusOK, gin.H{
			"fullname":    company.FullName,
			"companytype": company.CompanyType,
			"shortname":   company.ShortName,
			"scale":       company.Scale,
			"address":     company.Address,
			"brief":       company.Brief,
		})
	}
}
