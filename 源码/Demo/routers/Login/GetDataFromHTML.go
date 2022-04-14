package Login

import (
	"github.com/gin-gonic/gin"
	"log"
)

// GetDataFromHTML button的值最好当成参数传入函数中来，第一期就不做要求，2021.8.1
func GetDataFromHTML(c *gin.Context) (string, string, string) {
	log.SetPrefix("<<<GET DATA>>>")
	//获取前端信息
	inputphone := c.PostForm("inputphone")        		//获取前端输入的电话号码
	inputcode := c.PostForm("inputcode")              	//获取前端输入的验证码
	inputpassword := c.PostForm("inputpassword")      		//获取前端输入的密码
	//终端输出前端获log信息
	log.Printf("前端输入信息：usertel=%s, usercode=%s\n", inputphone, inputcode)
	return inputphone, inputcode, inputpassword 				  	//返回获取到的信息
}

//获取前端
func CompanyHeompageGetType(c *gin.Context) string {
	ButtonType := c.PostForm("ButtonType")
	return ButtonType
}

func CompanyGetDataFromHTML(c *gin.Context)(string,string,string,string,string,string,string,string,string,)  {
	fullName := c.PostForm("fullName")      //获取公司全名
	shortName := c.PostForm("shortName")	//获取公司简称
	companyType := c.PostForm("companyType")	//获取公司类型
	scale:= c.PostForm("scale")	//获取规模
	province:= c.PostForm("province")	//获取公司所在省
	city := c.PostForm("city")	//获取公司所在市
	specificAddress := c.PostForm("specificAddress")	//获取公司的具体地址
	logoUrl := c.PostForm("logUrl")	//获取公司logo的url
	brief := c.PostForm("brief")	//获取公司的简介

	log.SetPrefix("COMPLETE COMPANY INFORMATION")
	log.Println("fullName=",fullName)
	log.Println("shortName=",shortName)
	log.Println("companyType=",companyType)
	log.Println("scale=",scale)
	log.Println("province=",province)
	log.Println("city=",city)
	log.Println("specificAddress=",specificAddress)
	log.Println("logoUrl=",logoUrl)
	log.Println("brief=",brief)

	return fullName,shortName,companyType,scale,province,city,specificAddress,logoUrl,brief
}
