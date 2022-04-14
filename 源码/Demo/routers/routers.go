package routers

import (
	"Demo/es"
	"Demo/routers/Login"
	"Demo/routers/MainPage"
	"Demo/routers/application"
	"Demo/routers/company"
	"Demo/routers/job"
	"Demo/routers/person"
	"Demo/setting"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var SendCodeFlag = 0
var SessionSwitch = 0

//SetupRouter 控制路由
func SetupRouter(DB *gorm.DB, R redis.Conn) *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	Login.InitSession(r)
	//渲染界面
	r.LoadHTMLGlob("template/*")    //告诉gin html文件在templates下的任意位置
	r.Static("statics/", "statics") //告诉gin静态文件在statics下找

	var code string

	//v0组路由，主要用于获取验证码
	v0Group := r.Group("/v0")
	{
		v0Group.GET("/getcode", func(c *gin.Context) {})
		v0Group.POST("/getcode", func(c *gin.Context) {
			code = Login.SendCode(c, R, SendCodeFlag)
		})
	}

	// v1组路由,主要用于注册
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/cpregister", func(c *gin.Context) {
			Login.Setsess(c, "authorizedId", "false")
			c.HTML(http.StatusOK, "cpregister.html", "nil") //渲染企业用户注册界面
		})
		v1Group.POST("/cpregister", Login.CaseJudge(code, DB, R, 2))

		v1Group.GET("/personlogin", func(c *gin.Context) {
			Login.Setsess(c, "authorizedId", "false")
			c.HTML(http.StatusOK, "personlogin.html", "nil") //渲染个人用户注册/登录界面
		})
		v1Group.POST("/personlogin", Login.CaseJudge(code, DB, R, 1))

		v1Group.GET("/cplogin", func(c *gin.Context) {
			Login.Setsess(c, "authorizedId", "false")
			c.HTML(http.StatusOK, "cplogin.html", "nil") //渲染企业用户登陆界面
		})
		v1Group.POST("/cplogin", Login.CaseJudge(code, DB, R, 3))
	}

	// v2组路由,主要用于个人页面的编写
	v2Group := r.Group("/v2")
	{
		v2Group.GET("/recruit", func(context *gin.Context) {
			context.HTML(http.StatusOK, "resume.html", gin.H{"title": "hello"})
		})
		v2Group.POST("/recruit", person.GetPersonInfo(DB))
		//预览简历
		v2Group.GET("/resumePreview", func(c *gin.Context) {
			println("...........")
			c.HTML(http.StatusOK, "resumePreview.html", "nil") //渲染企业用户登陆界面
		})
		//v2Group.GET("/resumePreview", person.GenerateResume(DB))
		v2Group.POST("/resumePreview", person.GenerateResume(DB))

		//编辑简历'
		v2Group.GET("/resumeModify", person.GetPerson(DB))

		v2Group.POST("/resumeModify", person.EditPerson(DB))

		v2Group.GET("/application", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "myapplication.html", "nil")
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "myapplication.html", "nil")
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v2Group.POST("/application", application.PersonPostPage(DB))
		v2Group.GET("/jobinfo", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "jobinformation.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "jobinformation.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v2Group.POST("/jobinfo", job.InfoOfJob(DB))
	}

	//	v3组路由，主要用于企业页面
	v3Group := r.Group("/v3")
	{
		v3Group.GET("/company", func(context *gin.Context) {
			cid := context.Query("cid")
			if SessionSwitch == 0 {
				context.HTML(http.StatusOK, "companyHomePage.html", nil)
			} else {
				authorized := Login.Getsess(context, "authorizedId")
				if authorized == cid {
					context.HTML(http.StatusOK, "companyHomePage.html", nil)
				} else {
					context.String(http.StatusOK, "未登录")
				}
			}

		})
		v3Group.POST("/company", company.Companyhomepage(DB))

		v3Group.GET("/companyInfoEdit", func(context *gin.Context) {
			context.HTML(http.StatusOK, "companyInfoEdit.html", nil)
		})
		v3Group.POST("/companyInfoEdit", Login.CaseJudge(code, DB, R, 4))

		v3Group.GET("/positionRelease", func(c *gin.Context) {
			cid := c.Query("cid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "positionRelease.html", nil) //渲染企业用户登陆界面
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == cid {
					c.HTML(http.StatusOK, "positionRelease.html", nil) //渲染企业用户登陆界面
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v3Group.POST("/positionRelease", company.PositionRelease(DB))

		v3Group.GET("/jobCreate", func(c *gin.Context) {
			c.HTML(http.StatusOK, "jobCreate.html", "nil") //渲染增加职位界面
		})
		v3Group.POST("/jobCreate", job.JobCreate(DB))

		v3Group.GET("/jobEdit", func(c *gin.Context) {
			c.HTML(http.StatusOK, "jobEdit.html", nil) //渲染增加职位界面
		})
		v3Group.POST("/jobEdit", company.JobEditor(DB))

		v3Group.GET("/manageApply", func(c *gin.Context) {
			cid := c.Query("cid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "manageApply.html", "nil") //渲染管理职位页面
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == cid {
					c.HTML(http.StatusOK, "manageApply.html", "nil") //渲染管理职位页面
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v3Group.POST("/manageApply", company.Manageapply(DB))
		v3Group.GET("/viewResume", func(c *gin.Context) {
			c.HTML(http.StatusOK, "viewResume.html", "nil") //渲染企业用户登陆界面
		})
		v3Group.POST("/viewResume", company.ViewResume(DB))
	}

	// v4路由，主要用于主页面
	v4Group := r.Group("v4")
	{
		v4Group.GET("/home", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "main.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "main.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v4Group.POST("/home", MainPage.InfoTransfer(DB))

		v4Group.GET("/city", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "citymain.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "citymain.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v4Group.POST("/city", MainPage.InfoTransfer(DB))

		v4Group.GET("/scr", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "socialmain.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "socialmain.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v4Group.POST("/scr", MainPage.InfoTransfer(DB))

		v4Group.GET("/stu", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "studentmain.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "studentmain.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}

		})
		v4Group.POST("/stu", MainPage.InfoTransfer(DB))

		v4Group.GET("/search", func(c *gin.Context) {
			pid := c.Query("pid")
			if SessionSwitch == 0 {
				c.HTML(http.StatusOK, "searchdemo.html", nil)
			} else {
				authorized := Login.Getsess(c, "authorizedId")
				if authorized == pid {
					c.HTML(http.StatusOK, "searchdemo.html", nil)
				} else {
					c.String(http.StatusOK, "未登录")
				}
			}
		})
		v4Group.POST("/search", es.SearchHandler())
	}

	r.GET("/setSess", func(c *gin.Context) {
		Login.Setsess(c, "authorizedUsertel", "false")
	})
	r.GET("/getSess", func(c *gin.Context) {
		Login.Getsess(c, "authorizedUsertel")
	})
	return r
}
