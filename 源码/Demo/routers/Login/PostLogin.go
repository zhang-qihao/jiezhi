package Login

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

//CaseJudge 获取Post请求 做出对应的判断
func CaseJudge(code string, DB *gorm.DB, R redis.Conn, mode int) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch mode {
		//个人用户登陆注册/企业注册
		case 1, 2:
			usertel, usercode, _ := GetDataFromHTML(c) 						//获取前端数据usertel,usercode
			_, err := R.Do("select", 3)					//切换db：3 获取对应手机号匹配的验证码
			if err != nil {
				panic(err)
			}
			code, _ = redis.String(R.Do("get", usertel))		//获得对应手机号的code值存入code
			if mode == 1 {
				log.SetPrefix("<<<INDIVIDUAL CUSTOMERS SIGN UP>>>")
				PersonLoginRegister(usertel, usercode, c, R, DB)			//查找此手机号注册状态并进行相应操作
				c.Cookie("usertel")
				c.SetCookie("usertel","cookie",300,"/","localhost:8080",true,true)
			} else {
				log.SetPrefix("<<<ENTERPRISE USER SIGN UP>>>")
				CompanyRegister(usertel, usercode, c, R, DB)			//查找此手机号注册状态并进行相应操作
			}
		//企业登陆
		case 3:
			usertel, usercode, userpassword := GetDataFromHTML(c) 			//获取前端数据usertel,usercode,userpassword
			if userpassword != "" {
				log.SetPrefix("<<<ENTERPRISE USER SIGN IN BY PASSWORD>>>")
				//进行密码登录
				_, _ = R.Do("select", 4)				//切换db：4 获取对应手机号匹配的密码
				CompanySecretLogin(usertel, userpassword, DB, c, R)				//查找用户是否已经注册并比对密码是否正确
			}else {
				log.SetPrefix("<<<ENTERPRISE USER SIGN IN BY MESSAGE>>>")
				//进行短信登陆
				_, err := R.Do("select", 3) 			//切换db：3 获取对应手机号匹配的验证码
				if err != nil {
					panic(err)
				}
				code, err = redis.String(R.Do("get", usertel))//获得对应手机号的code值存入code
				if err != nil {
					panic(err)
				}
				_, err = R.Do("select", "2") 			//切换db:2 用来查询手机号是否注册
				if err != nil {
					panic(err)
				}
				CompanyMessageLogin(usertel, usercode, code, DB, c, R)			//查找用户是否已经注册并比对验证码是否正确
			}
		//企业修改密码
		case 4:
			exit := c.PostForm("exit")
			if exit != "" {
				Setsess(c, "authorizedId", "false")
			}
			ButtonType := CompanyHeompageGetType(c)
			if (ButtonType == "SetPwd"){
				log.SetPrefix("<<<ENTERPRISE USER MODIFY OR SET PASSWORD>>>")
				usertel, usercode, userpassword := GetDataFromHTML(c)			//获取前端数据usertel,usercode,userpassword
				SetPwd(usertel, userpassword, usercode, c, R, DB)				//查找用户是否已经注册并比对验证码是否正确，正确则写入数据库
			}else if(ButtonType == "SaveInfo"){
				log.SetPrefix("<<<ENTERPRISE USER FILL IN COMPANY INFORMATION>>>")
				fullName,shortName,companyType,scale,province,city,specificAddress,logoUrl,brief := CompanyGetDataFromHTML(c)
				//生成唯一标识符ID
				CID := CreateID(fullName+specificAddress+logoUrl)
				CompanyWriteToDatabase(CID,fullName,shortName,companyType,scale,province,city,specificAddress,logoUrl,brief,DB,c)
				err := Setsess(c, "authorizedId", CID)
				if err != nil {
					return
				}
			}			//查找用户是否已经注册并比对验证码是否正确，正确则写入数据库
		}
	}
}