package Login

import (
	"Demo/conf"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

//WriteMySQL 信息写入数据库
func WriteMySQL(usertel, userpassword, usertype, ID string, DB *gorm.DB, c *gin.Context) { //mode用来判断是更新用户信息还是创建新用户

	if usertype == "private" {
		log.Println("注册新个人用户...") //终端提示注册新个人用户

		//把抓取的数据放入定义的结构体类型infor中
		infor := conf.PrivateAccount{
			Pid:       ID,
			Telephone: usertel,
		}

		DB.Save(&infor)                                          //信息存入MySQL数据库并返回响应
		log.Println("新用户注册成功，写入MySQL，table:private_account")     //终端提示注册成功
		c.JSON(http.StatusOK, gin.H{"msg": "新用户注册成功", "ID": ID}) //给前端传递JSON数据："msg":"新用户注册成功"
	} else {
		log.Println("注册新企业用户...") //终端提示注册新企业用户

		//把抓取的数据放入定义的结构体类型infor中
		infor := conf.BusinessAccount{
			Cid:       ID,
			Telephone: usertel,
			Password:  userpassword,
		}

		DB.Save(&infor)                                          //信息存入MySQL数据库并返回响应
		log.Println("新用户注册成功，写入MySQL，table:business_account")    //终端输出提示
		c.JSON(http.StatusOK, gin.H{"msg": "新用户注册成功", "ID": ID}) //给前端传递JSON数据："msg":"新用户注册成功"
	}
}

//UpdateMySQL 更新数据库
func UpdateMySQL(usertel, userpassword, Cid string, DB *gorm.DB, c *gin.Context, mode int) {
	log.Println("数据库企业信息更新...") //终端提示数据库企业信息更新

	//把抓取的数据放入定义的结构体类型infor中
	if mode == 1 {
		infor := conf.BusinessAccount{
			Password: userpassword,
		}
		//查询business_account表中telephone = usertel的字段，并更新密码，返回错误信息
		err := DB.Find(&conf.BusinessAccount{}).Where("telephone = ?", usertel).Update(&infor).Error
		if err != nil {
			panic(err)
		} else {
			log.Println("信息更新成功")                         //终端提示信息更改成功
			c.JSON(http.StatusOK, gin.H{"msg": "信息更新成功"}) //给前端传递JSON数据："msg":"企业信息更新成功"
		}
	} else {
		infor := conf.BusinessAccount{
			Cid: Cid,
		}
		//查询business_account表中telephone = usertel的字段，并更新密码，返回错误信息
		err := DB.Find(&conf.BusinessAccount{}).Where("telephone = ?", usertel).Update(&infor).Error
		if err != nil {
			panic(err)
		} else {
			log.Println("信息更新成功") //终端提示信息更改成功
			log.Println("cid=", Cid)
			c.JSON(http.StatusOK, gin.H{"msg": "信息更新成功", "ID": Cid}) //给前端传递JSON数据："msg":"企业信息更新成功"
		}
	}
}

//WriteRedis 信息写入redis
func WriteRedis(R redis.Conn, key, value string, db, mode int) {
	log.SetPrefix("<<<WRITE REDIS>>>")
	_, err := R.Do("select", db) //切换db
	if err != nil {
		panic(err)
	}
	_, err = R.Do("set", key, value) //用户信息写入redis
	if err != nil {
		panic(err)
	} else {
		//终端输出提示
		if mode == 1 {
			log.Printf("写入Redis value = %v - key = %v   成功 databasse: %v\n", key, value, mode)
		} else {
			log.Printf("更新Redis value = %v - key = %v   成功 databasse: %v\n", key, value, mode)
		}
	}
}

func CompanyWriteToDatabase(Cid, fullName, shortName, companyType, scale, province, city, specificAddress, logoUrl, brief string, DB *gorm.DB, c *gin.Context) {
	currentTime := time.Now()
	companyInfo := conf.CompanyInfo{
		Cid:         Cid,
		ShortName:   shortName,
		FullName:    fullName,
		CompanyType: companyType,
		Address:     province + "省" + city + specificAddress,
		Industry:    "NULL",
		Scale:       scale,
		LogoUrl:     logoUrl,
		Brief:       brief,
		CreatTime:   currentTime,
		ChangeTime:  currentTime,
	}
	company := new(conf.CompanyInfo)
	DB.Where("full_name=?", fullName).Find(&company)
	log.Println("RegisteredCompany=", company.FullName)
	if company.FullName != "" {
		c.JSON(200, gin.H{"msg": "该公司已被注册"})
	} else {
		usertel, _, _ := GetDataFromHTML(c)
		DB.Save(&companyInfo)
		UpdateMySQL(usertel, "", Cid, DB, c, 2)
	}
}
