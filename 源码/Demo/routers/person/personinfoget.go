package person

import (
	"Demo/conf"
	"Demo/connet"
	"Demo/routers/Login"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func GetPersonInfo(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//根据获取的前端姓名手机号生成Pid
		name := context.PostForm("userName")
		phone := context.PostForm("userTel")
		s := name + phone
		ID := Login.CreateID(s)
		// 个人信息的总表
		var personInfoStr conf.PersonInfoStr
		// 个人登录的第1，2，3个子页面需要填写的字段于one,two,threeForm绑定
		userName := context.PostForm("userName")
		gender := context.PostForm("gender")
		identity := context.PostForm("identity")
		userBirth := context.PostForm("userBirth")
		userPlace := context.PostForm("userPlace")
		userTel := context.PostForm("userTel")
		userEmail := context.PostForm("userEmail")
		expectJob := context.PostForm("expectJob")
		expectCity := context.PostForm("expectCity")
		expectSalary1 := context.PostForm("expectSalary1")
		expectSalary2 := context.PostForm("expectSalary2")
		nature_of_works := context.PostForm("nature_of_works")
		workTime1 := context.PostForm("workTime1")
		school := context.PostForm("school")
		topEducation := context.PostForm("topEducation")
		majorStudy := context.PostForm("majorStudy")
		schoolTime1 := context.PostForm("schoolTime1")
		schoolTime2 := context.PostForm("schoolTime2")
		//基础信息
		fmt.Println("===========nature_of_works==========", nature_of_works)
		var oneForm = conf.BasicInfoElement{
			userTel,
			userName,
			gender,
			identity,
			userBirth,
			userPlace,
			userEmail,
		}
		//期望职业
		var twoForm = conf.ExpectJobElement{
			expectJob,
			expectCity,
			expectSalary1,
			expectSalary2,
			nature_of_works,
			workTime1,
		}
		//教育经历
		var threeForm = conf.EducationElement{
			school,
			topEducation,
			majorStudy,
			schoolTime1,
			schoolTime2,
			"",
		}
		if threeForm.SchoolName != "" {
			//new结构体
			var infor = conf.PrivateAccount{
				Pid: ID,
			}
			// 说明有数值，需要写入数据库
			// step1:将数据进行拼接
			oneFormStr := connet.Merge(oneForm)
			twoFormStr := connet.Merge(twoForm)
			threeFormStr := connet.Merge(threeForm)
			// step2: 将数据赋值给personInfoStr
			personInfoStr.Pid = ID
			personInfoStr.BasicInfo = oneFormStr
			personInfoStr.ExpectJob1 = twoFormStr
			personInfoStr.Education1 = threeFormStr
			personInfoStr.ChangeTime = time.Now()
			personInfoStr.CreatTime = time.Now()
			// step3: 将数据写入给数据库
			DB.Find(&conf.PrivateAccount{}).Where("telephone = ?", phone).Update(&infor) //更新数据库
			DB.Save(&personInfoStr)                                                      //信息存入MySQL数据库并返回响应
			fmt.Println("写入MySQL成功")
		}
		//new结构体
		var infor conf.PersonInfoStr
		//查找对应手机号用户的Pid和姓名
		DB.Where("basic_info like ?", phone+"%").Find(&infor)
		context.JSON(http.StatusOK, gin.H{"msg": "新用户信息填写成功", "Name": oneForm.Name, "ID": infor.Pid})
		err := Login.Setsess(context, "authorizedId", infor.Pid)
		if err != nil {
			return
		}
	}
}
