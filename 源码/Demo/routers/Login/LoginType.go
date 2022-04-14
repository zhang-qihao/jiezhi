package Login

import (
	"Demo/conf"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

//PersonLoginRegister 个人用户短信登陆/注册
func PersonLoginRegister(usertel, usercode string, c *gin.Context, R redis.Conn, DB *gorm.DB) {
	_, err := R.Do("select", 1) //切换db：1 查询手机号是否注册
	if err != nil {
		panic(err)
	}
	find, _ := redis.String(R.Do("get", usertel)) //查找前段获取到的手机号对应的value存入find中，用来判断手机号是否注册
	_, err = R.Do("select", 3)                    //切换db：3 获取对应手机号匹配的验证码
	if err != nil {
		panic(err)
	}
	code, _ := redis.String(R.Do("get", usertel)) //查找前端获取到的手机号对应的value值存入code中

	//判断手机号是否存在在redis中
	//存在：比对code->正确：进入主页
	//				错误：返回提示
	//不存在：比对code->正确：手机号写入MySQL和redis db：private_account和 1
	//				  错误：返回提示
	if find == "" {
		log.Println("手机号不存在")
		//比对code
		if usercode == code {
			WriteMySQL(usertel, "", "private", "", DB, c) //手机号写入MySQL
			WriteRedis(R, usertel, "1", 1, 1)             //手机号写入Redis db:1
		} else {
			log.Println("验证码错误")
			c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"}) //给前端传递JSON数据："msg":"验证码错误"
		}
	} else {
		//比对code
		if usercode == code {
			var info = new(conf.PersonInfoStr)
			DB.Where("basic_info like ?", usertel+"%").Find(&info)
			if info.BasicInfo != "" {
				c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "ID": info.Pid}) //给前端传递JSON数据："msg":"登陆成功"
				log.Println("登陆成功")
				err = Setsess(c, "authorizedId", info.Pid)
				if err != nil {
					return
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "未填写个人简历"})
				log.Println("个人简历未填写")
			}
		} else {
			log.Println("验证码错误")
			c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"}) //给前端传递JSON数据："msg":"验证码错误"
		}
	}
}

//CompanyRegister 企业注册
func CompanyRegister(usertel, usercode string, c *gin.Context, R redis.Conn, DB *gorm.DB) {
	_, err := R.Do("select", 2) //切换db：2 查询手机号是否注册
	if err != nil {
		panic(err)
	}
	find, _ := redis.String(R.Do("get", usertel)) //查找前段获取到的手机号对应的value存入find中，用来判断手机号是否注册
	_, err = R.Do("select", 3)                    //切换db：3 获取对应手机号匹配的验证码
	if err != nil {
		panic(err)
	}
	code, _ := redis.String(R.Do("get", usertel)) //查找前端获取到的手机号对应的value值存入code中

	//判断手机号是否存在在redis中
	//存在：比对code->返回提示
	//不存在：比对code->正确：手机号写入MySQL和redis db：business_account和 2
	//				  错误：返回提示
	if find == "" {
		log.Println("手机号不存在") //终端提示手机号不在redis中
		//比对code
		if usercode == code {
			WriteMySQL(usertel, "", "business", "", DB, c) //手机号写入MySQL
			WriteRedis(R, usertel, "1", 2, 1)              //手机号写入Redis db:2
			WriteRedis(R, usertel, usertel, 4, 1)          //手机号写入Redis db:4
		} else {
			log.Println("验证码错误")                         //终端提示验证码错误
			c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"}) //给前端传递JSON数据："msg":"验证码错误"
		}
	} else {
		//比对code
		if usercode == code {
			log.Println("验证码正确")                           //终端提示验证码正确
			c.JSON(http.StatusOK, gin.H{"msg": "手机号码已注册"}) //给前端传递JSON数据："msg":"登陆成功"
			log.Println("手机号码已注册")                         //终端提示登陆成功
		} else {
			log.Println("验证码错误")                         //终端提示验证码错误
			c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"}) //给前端传递JSON数据："msg":"验证码错误"
		}
	}
}

//CompanySecretLogin 企业密码登陆
func CompanySecretLogin(usertel, userpassword string, DB *gorm.DB, c *gin.Context, R redis.Conn) {

	_, err := R.Do("select", "4") //切换db:4 判断手机号是否注册
	if err != nil {
		panic(err)
	}
	find, _ := redis.String(R.Do("get", usertel)) //查找前段获取到的手机号对应的redis键值存入find中

	//判断手机号是否存在在redis中
	//存在:比对password->正确：登陆成功
	//				    错误：登陆错误
	//不存在:->号码未注册
	if find == "" {
		//手机号不在redis中
		log.Println("手机号未注册")                           //终端提示手机号未注册
		c.JSON(http.StatusOK, gin.H{"error": "手机号未注册"}) //给前端传递JSON数据："msg":"手机号未注册"
	} else {
		log.Println("手机号已经存在") //手机号已经存在
		if find == userpassword {
			var info = new(conf.BusinessAccount) //查找对应的id
			DB.Where("telephone = ?", usertel).Find(&info)
			c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "ID": info.Cid}) //给前端传递JSON数据："msg":"登陆成功"
			log.Println("登陆成功")                                         //终端提示登陆成功
			err = Setsess(c, "authorizedId", info.Cid)
			if err != nil {
				return
			}
		} else {
			log.Println("密码错误")                         //终端提示密码错误
			c.JSON(http.StatusOK, gin.H{"msg": "密码错误"}) //给前端传递JSON数据："msg":"密码错误"
		}
	}
}

//企业短信登录
func CompanyMessageLogin(usertel, usercode, code string, DB *gorm.DB, c *gin.Context, R redis.Conn) {
	_, err := R.Do("select", 2) //切换db：2 用来查找手机号是否存在
	if err != nil {
		panic(err)
	}
	find, _ := redis.String(R.Do("get", usertel)) //查找前段获取到的手机号对应的value存入find中

	//判断手机号是否存在在redis中
	//存在：比对code->正确：返回提示
	//				错误：返回提示
	//不存在：返回提示
	if find == "" {
		//手机号不在redis中
		log.Println("手机号未注册") //终端提示手机号未注册
		c.JSON(http.StatusOK, gin.H{"msg": "手机号未注册"})
	} else {
		//手机号已经存在
		//比对code
		if usercode == code {
			var info = new(conf.BusinessAccount) //查找对应的id
			DB.Where("telephone = ?", usertel).Find(&info)
			c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "ID": info.Cid}) //给前端传递JSON数据："msg":"登陆成功"
			log.Println("登陆成功")                                         //终端提示登陆成功
			err = Setsess(c, "authorizedId", info.Cid)
			if err != nil {
				return
			}
		} else {
			log.Println("验证码错误")                         ////终端提示验证码错误
			c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"}) //给前端传递JSON数据："msg":"验证码错误"
		}
	}
}
