package person

import (
	"Demo/conf"
	"Demo/connet"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func EditPerson(DB *gorm.DB) gin.HandlerFunc {
	var pid string
	return func(context *gin.Context) {
		fmt.Println("已经进入EditPerson 函数内部")
		pid = context.PostForm("pid") //获取pid
		var personInfoStr conf.PersonInfoStr
		modifyType := context.PostForm("modifyType")
		DB.Where("Pid = ?", pid).Find(&personInfoStr) //数据库中查找数据
		//修改个人基础信息
		if modifyType == "basic" {
			//basicForm
			userName := context.PostForm("userName")
			gender := context.PostForm("gender")
			identity := context.PostForm("identity")
			userBirth := context.PostForm("userBirth")
			userPlace := context.PostForm("userPlace")
			userTel := context.PostForm("userTel")
			userEmail := context.PostForm("userEmail")
			pid = context.PostForm("pid")
			personInfoStr.ChangeTime = time.Now()
			var basicForm = conf.BasicInfoElement{
				userTel,
				userName,
				gender,
				identity,
				userBirth,
				userPlace,
				userEmail,
			}
			basicFormStr := connet.Merge(basicForm)
			if basicForm.Name != "" {
				personInfoStr.BasicInfo = basicFormStr
				DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
				fmt.Println("修改基础信息成功")
			}
		}

		//修改求职意向表
		if modifyType == "expectJob" {
			//expectForm
			jobName := context.PostForm("jobName")
			expectCity := context.PostForm("expectCity")
			jobMoneyMin := context.PostForm("jobMoneyMin")
			jobMoneyMax := context.PostForm("jobMoneyMax")
			jobType := context.PostForm("jobType")
			jobTime := context.PostForm("jobTime")
			personInfoStr.ChangeTime = time.Now()
			var expectForm = conf.ExpectJobElement{
				jobName,
				expectCity,
				jobMoneyMin,
				jobMoneyMax,
				jobType,
				jobTime,
			}

			if expectForm.ExpectedOccupation != "" {
				expectFormStr := connet.Merge(expectForm)
				personInfoStr.ExpectJob1 = expectFormStr
				DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
				fmt.Println("修改期望职业成功")
			}
		}

		//修改教育经历
		if modifyType == "education1" || modifyType == "education2" || modifyType == "education3" {
			//educateForm
			schoolName := context.PostForm("school_name")
			highestEducation := context.PostForm("highest_education")
			majorStudied := context.PostForm("major_studied")
			schoolTimeStart := context.PostForm("school_time_start")
			schoolTimeEnd := context.PostForm("school_time_end")
			personInfoStr.ChangeTime = time.Now()
			var educateForm = conf.EducationElement{
				schoolName,
				highestEducation,
				majorStudied,
				schoolTimeStart,
				schoolTimeEnd,
				"",
			}
			educateFormStr := connet.Merge(educateForm)
			switch {
			case modifyType == "education1":
				{
					if educateForm.SchoolName != "" {
						personInfoStr.Education1 = educateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改教育经历1成功")
					}
					modifyType = ""
					break
				}
			case modifyType == "education2":
				{
					if educateForm.SchoolName != "" {
						personInfoStr.Education2 = educateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改教育经历2成功")
					}
					modifyType = ""
					break
				}
			case modifyType == "education3":
				{
					if educateForm.SchoolName != "" {
						personInfoStr.Education3 = educateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改教育经历3成功")
					}

				}
			}
		}
		//修改工作经历
		if modifyType == "workExperience1" || modifyType == "workExperience2" || modifyType == "workExperience3" {
			//workForm
			workCompanyName := context.PostForm("workCompanyName")
			workName := context.PostForm("workName")
			workTime1 := context.PostForm("workTime1")
			workTime2 := context.PostForm("workTime2")
			workMoney := context.PostForm("workMoney")
			workType := context.PostForm("workType")
			workLike := context.PostForm("workLike")
			personInfoStr.ChangeTime = time.Now()
			var workForm = conf.WorkExperienceElement{
				workCompanyName,
				workName,
				workTime1,
				workTime2,
				workMoney,
				workType,
				workLike,
			}
			workFormStr := connet.Merge(workForm)
			switch {
			case modifyType == "workExperience1":
				{
					if workForm.Company != "" {
						personInfoStr.Work1 = workFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改工作经历1成功")
					}
				}
			case modifyType == "workExperience2":
				{
					if workForm.Company != "" {
						personInfoStr.Work2 = workFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改工作经历2成功")
					}
				}
			case modifyType == "workExperience3":
				{
					if workForm.Company != "" {
						personInfoStr.Work3 = workFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改工作经历3成功")
					}
				}
			}
		}
		//修改项目经历
		if modifyType == "item1" || modifyType == "item2" || modifyType == "item3" {
			//projectForm
			projectName := context.PostForm("projectName")
			projectTime1 := context.PostForm("projectTime1")
			projectTime2 := context.PostForm("projectTime2")
			projectCompany := context.PostForm("projectCompany")
			projectLike := context.PostForm("projectLike")
			personInfoStr.ChangeTime = time.Now()
			var projectForm = conf.ItemElement{
				projectName,
				projectTime1,
				projectTime2,
				projectCompany,
				projectLike,
			}
			projectFormStr := connet.Merge(projectForm)
			switch {
			case modifyType == "item1":
				{
					if projectForm.Title != "" {
						personInfoStr.Item1 = projectFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改项目经历1成功")
					}
				}
			case modifyType == "item2":
				{
					if projectForm.Title != "" {
						personInfoStr.Item2 = projectFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改项目经历2成功")
					}
				}
			case modifyType == "item3":
				if projectForm.Title != "" {
					personInfoStr.Item3 = projectFormStr
					DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
					fmt.Println("修改项目经历3成功")
				}
			}
		}
		//修改校内表现
		if modifyType == "school1" || modifyType == "school2" || modifyType == "school3" {
			//schoolForm
			award := context.PostForm("award")
			post := context.PostForm("post")
			describe := context.PostForm("describe")
			personInfoStr.ChangeTime = time.Now()
			var schoolForm = conf.SchoolPostElement{
				award,
				post,
				describe,
			}
			schoolFormStr := connet.Merge(schoolForm)
			switch {
			case modifyType == "school1":
				{
					if schoolForm.Award != "" {
						personInfoStr.SchoolPost1 = schoolFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改校内表现1成功")
					}
				}
			case modifyType == "school2":
				{
					if schoolForm.Award != "" {
						personInfoStr.SchoolPost2 = schoolFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改校内表现2成功")
					}
				}
			case modifyType == "school3":
				if schoolForm.Award != "" {
					personInfoStr.SchoolPost3 = schoolFormStr
					DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
					fmt.Println("修改校内表现3成功")
				}
			}
		}
		//修改专业技能
		if modifyType == "special1" || modifyType == "special2" || modifyType == "special3" {
			//specialForm
			skill := context.PostForm("skill")
			skillTime := context.PostForm("skillTime")
			personInfoStr.ChangeTime = time.Now()
			var specialForm = conf.SpecialElement{
				skill,
				skillTime,
			}
			specialFormStr := connet.Merge(specialForm)
			switch {
			case modifyType == "special1":
				{
					if specialForm.Skill != "" {
						personInfoStr.SpecialSkill1 = specialFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改专业技能1成功")
					}
				}
			case modifyType == "special2":
				{
					if specialForm.Skill != "" {
						personInfoStr.SpecialSkill2 = specialFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改专业技能2成功")
					}
				}
			case modifyType == "special3":
				if specialForm.Skill != "" {
					personInfoStr.SpecialSkill3 = specialFormStr
					DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
					fmt.Println("修改专业技能3成功")
				}
			}

		}
		//修改语言能力
		if modifyType == "language1" || modifyType == "language2" || modifyType == "language3" {
			//languageForm
			languageType := context.PostForm("languageType")
			languageTime := context.PostForm("languageTime")
			personInfoStr.ChangeTime = time.Now()
			var languageForm = conf.LanguageElement{
				languageType,
				languageTime,
			}
			switch {
			case modifyType == "language1":
				{
					if languageForm.LanguageType != "" {
						languageFormStr := connet.Merge(languageForm)
						personInfoStr.LanguageSkill1 = languageFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改语言能力1成功")
					}
				}
			case modifyType == "language2":
				{
					if languageForm.LanguageType != "" {
						languageFormStr := connet.Merge(languageForm)
						personInfoStr.LanguageSkill2 = languageFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改语言能力2成功")
					}
				}
			case modifyType == "language3":
				{
					if languageForm.LanguageType != "" {
						languageFormStr := connet.Merge(languageForm)
						personInfoStr.LanguageSkill3 = languageFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改语言能力3成功")
					}
				}
			}
		}
		//修改培训经历
		if modifyType == "train1" || modifyType == "train2" || modifyType == "train3" {
			//trainForm
			trainCompany := context.PostForm("trainCompany")
			trainClass := context.PostForm("trainClass")
			trainTime := context.PostForm("trainTime")
			trainDescribe := context.PostForm("trainDescribe")
			personInfoStr.ChangeTime = time.Now()
			var trainForm = conf.TrainingElement{
				trainCompany,
				trainClass,
				trainTime,
				trainDescribe,
			}
			switch {
			case modifyType == "train1":
				{
					if trainForm.Organization != "" {
						trainFormStr := connet.Merge(trainForm)
						personInfoStr.TrainSkill1 = trainFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改培训经历1成功")
					}
				}
			case modifyType == "train2":
				{
					if trainForm.Organization != "" {
						trainFormStr := connet.Merge(trainForm)
						personInfoStr.TrainSkill2 = trainFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改培训经历2成功")
					}
				}
			case modifyType == "train3":
				{
					if trainForm.Organization != "" {
						trainFormStr := connet.Merge(trainForm)
						personInfoStr.TrainSkill3 = trainFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改培训经历3成功")
					}
				}
			}
		}
		//修改获得证书
		if modifyType == "certificate1" || modifyType == "certificate2" || modifyType == "certificate3" {
			//certificateForm
			certificateName := context.PostForm("certificateName")
			certificateTime := context.PostForm("certificateTime")
			personInfoStr.ChangeTime = time.Now()
			var certificateForm = conf.CertificateElement{
				certificateName,
				certificateTime,
			}
			certificateFormStr := connet.Merge(certificateForm)
			switch {
			case modifyType == "certificate1":
				{
					if certificateForm.CertificateType != "" {
						personInfoStr.Certificate1 = certificateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改获得证书1成功")
					}
				}
			case modifyType == "certificate2":
				{
					if certificateForm.CertificateType != "" {
						personInfoStr.Certificate2 = certificateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改获得证书2成功")
					}
				}
			case modifyType == "certificate3":
				{
					if certificateForm.CertificateType != "" {
						personInfoStr.Certificate3 = certificateFormStr
						DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
						fmt.Println("修改获得证书3成功")
					}
				}
			}
		}
		//修改自我评价
		if modifyType == "selfEvaluation" {
			selfEvaluation := context.PostForm("TxtSelfEvaluation")
			personInfoStr.ChangeTime = time.Now()
			var selfForm = conf.SelfEvaluation{
				selfEvaluation,
			}
			selfFormStr := connet.Merge(selfForm)
			if selfEvaluation != "" {
				personInfoStr.SelfEvaluation = selfFormStr
				DB.Model(&personInfoStr).Where("Pid = ?", pid).Update(&personInfoStr)
				fmt.Println("修改自我评价成功")
			}
		}

	}
}
