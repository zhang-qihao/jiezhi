package Login

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

//查询、更新数据
func SetPwd(usertel, userpassword, usercode string, c *gin.Context, R redis.Conn, DB *gorm.DB) {
		//切换db：4
		_, err := R.Do("select", 4)
		if err != nil {
			panic(err)
		}

		//查找前段获取到的手机号对应的redis键值存入find中
		find, _ := redis.String(R.Do("get", usertel))

		//判断手机号是否存在在redis中
		//存在：比对code->正确：userpwd写入redis和MySQL中
		//				错误：返回提示
		//不存在：返回提示
		if find == "" {
			//手机号不在redis中
			c.JSON(http.StatusOK, gin.H{"msg": "手机号码不存在"})
		}else {
			//手机号已经存在
			_, err = R.Do("select", 3)		//切换db：3
			if err != nil {
				panic(err)
			}
			//查找前端获取到的手机号对应的redis键值即存放的验证码存入code中
			code, _ := redis.String(R.Do("get", usertel))
			//比对code
			if usercode == code {
				WriteRedis(R, usertel, userpassword, 4, 2)
				UpdateMySQL(usertel, userpassword, "", DB, c, 1)
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": "验证码错误"})
				log.Println("验证码错误")
			}
		}


}
