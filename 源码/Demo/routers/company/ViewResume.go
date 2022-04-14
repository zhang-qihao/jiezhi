package company

import (
	"Demo/conf"
	"Demo/connet"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func ViewResume(DB *gorm.DB) gin.HandlerFunc {
	var mode int
	return func(context *gin.Context) {
		//fmt.Printf("1: URL:=%s\n",context.Request.URL)
		cid := context.PostForm("cid")
		pid := context.PostForm("pid")
		jid := context.PostForm("jid")
		clickStatus := context.PostForm("clickStatus")
		var personInfoStr conf.PersonInfoStr
		DB.Where("pid=?", pid).Find(&personInfoStr)
		fmt.Printf("[ViewResume] pid=%s,cid=%s, jid=%s\n", pid, cid, jid)
		fmt.Printf("[ViewResume] status=%s\n", clickStatus)
		if personInfoStr.Pid != "" {
			if clickStatus == "" {
				// 进行字段的解析,标号为8
				mode = 8
				basicField := connet.GetFieldName(conf.StructMap[mode])
				basicInfoStr := connet.Split(personInfoStr.BasicInfo)
				BasicStruct := connet.SetValueToStruct(basicField, basicInfoStr, mode)
				BasicInfo := *BasicStruct.(*conf.BasicInfoElement)

				// 进行字段的解析,标号为9
				var ExpectJob [3]conf.ExpectJobElement
				mode = 9
				var jobCount = 0
				JobField := connet.GetFieldName(conf.StructMap[mode])
				Job1InfoStr := connet.Split(personInfoStr.ExpectJob1)
				Job1Struct := connet.SetValueToStruct(JobField, Job1InfoStr, mode)
				ExpectJob[0] = *Job1Struct.(*conf.ExpectJobElement)
				Job2InfoStr := connet.Split(personInfoStr.ExpectJob2)
				Job2Struct := connet.SetValueToStruct(JobField, Job2InfoStr, mode)
				ExpectJob[1] = *Job2Struct.(*conf.ExpectJobElement)
				Job3InfoStr := connet.Split(personInfoStr.ExpectJob3)
				Job3Struct := connet.SetValueToStruct(JobField, Job3InfoStr, mode)
				ExpectJob[2] = *Job3Struct.(*conf.ExpectJobElement)

				for i := 0; i < 3; i++ {
					if ExpectJob[i].ExpectedOccupation != "" {
						jobCount++
					}
				}

				// 进行字段的解析,标号为10
				var Education [3]conf.EducationElement
				mode = 10
				var eduCount = 0
				EduField := connet.GetFieldName(conf.StructMap[mode])
				Edu1InfoStr := connet.Split(personInfoStr.Education1)
				Edu1Struct := connet.SetValueToStruct(EduField, Edu1InfoStr, mode)
				Education[0] = *(Edu1Struct.(*conf.EducationElement))
				Edu2InfoStr := connet.Split(personInfoStr.Education2)
				Edu2Struct := connet.SetValueToStruct(EduField, Edu2InfoStr, mode)
				Education[1] = *(Edu2Struct.(*conf.EducationElement))
				Edu3InfoStr := connet.Split(personInfoStr.Education3)
				Edu3Struct := connet.SetValueToStruct(EduField, Edu3InfoStr, mode)
				Education[2] = *(Edu3Struct.(*conf.EducationElement))
				for i := 0; i < 3; i++ {
					if Education[i].SchoolName != "" {
						eduCount++
					}
				}

				// 进行字段的解析,标号为11
				var Work [3]conf.WorkExperienceElement
				mode = 11
				var workCount = 0
				WorkField := connet.GetFieldName(conf.StructMap[mode])
				Work1InfoStr := connet.Split(personInfoStr.Work1)
				Work1Struct := connet.SetValueToStruct(WorkField, Work1InfoStr, mode)
				Work2InfoStr := connet.Split(personInfoStr.Work2)
				Work2Struct := connet.SetValueToStruct(WorkField, Work2InfoStr, mode)
				Work3InfoStr := connet.Split(personInfoStr.Work3)
				Work3Struct := connet.SetValueToStruct(WorkField, Work3InfoStr, mode)
				Work[0] = *Work1Struct.(*conf.WorkExperienceElement)
				Work[1] = *Work2Struct.(*conf.WorkExperienceElement)
				Work[2] = *Work3Struct.(*conf.WorkExperienceElement)
				for i := 0; i < 3; i++ {
					if Work[i].Job != "" {
						workCount++
					}
				}
				// 进行字段的解析,标号为12
				var Item [3]conf.ItemElement
				mode = 12
				var itemCount = 0
				ItemField := connet.GetFieldName(conf.StructMap[mode])
				Item1InfoStr := connet.Split(personInfoStr.Work1)
				Item1Struct := connet.SetValueToStruct(ItemField, Item1InfoStr, mode)
				Item2InfoStr := connet.Split(personInfoStr.Work2)
				Item2Struct := connet.SetValueToStruct(ItemField, Item2InfoStr, mode)
				Item3InfoStr := connet.Split(personInfoStr.Work3)
				Item3Struct := connet.SetValueToStruct(ItemField, Item3InfoStr, mode)
				Item[0] = *Item1Struct.(*conf.ItemElement)
				Item[1] = *Item2Struct.(*conf.ItemElement)
				Item[2] = *Item3Struct.(*conf.ItemElement)
				for i := 0; i < 3; i++ {
					if Item[i].Title != "" {
						itemCount++
					}
				}
				// 进行字段的解析,标号为14
				var School [3]conf.SchoolPostElement
				mode = 14
				var schoolCount = 0
				SchoolField := connet.GetFieldName(conf.StructMap[mode])
				School1InfoStr := connet.Split(personInfoStr.SchoolPost1)
				School1Struct := connet.SetValueToStruct(SchoolField, School1InfoStr, mode)
				School2InfoStr := connet.Split(personInfoStr.SchoolPost1)
				School2Struct := connet.SetValueToStruct(SchoolField, School2InfoStr, mode)
				School3InfoStr := connet.Split(personInfoStr.SchoolPost1)
				School3Struct := connet.SetValueToStruct(SchoolField, School3InfoStr, mode)
				School[0] = *School1Struct.(*conf.SchoolPostElement)
				School[1] = *School2Struct.(*conf.SchoolPostElement)
				School[2] = *School3Struct.(*conf.SchoolPostElement)
				for i := 0; i < 3; i++ {
					if School[i].Award != "" {
						schoolCount++
					}
				}

				// 进行字段的解析,标号为16
				var Skill [3]conf.SpecialElement
				mode = 16
				var skillCount = 0
				SkillField := connet.GetFieldName(conf.StructMap[mode])
				Skill1InfoStr := connet.Split(personInfoStr.SpecialSkill1)
				Skill1Struct := connet.SetValueToStruct(SkillField, Skill1InfoStr, mode)
				Skill2InfoStr := connet.Split(personInfoStr.SpecialSkill2)
				Skill2Struct := connet.SetValueToStruct(SkillField, Skill2InfoStr, mode)
				Skill3InfoStr := connet.Split(personInfoStr.SpecialSkill3)
				Skill3Struct := connet.SetValueToStruct(SkillField, Skill3InfoStr, mode)
				Skill[0] = *Skill1Struct.(*conf.SpecialElement)
				Skill[1] = *Skill2Struct.(*conf.SpecialElement)
				Skill[2] = *Skill3Struct.(*conf.SpecialElement)
				for i := 0; i < 3; i++ {
					if Skill[i].Skill != "" {
						skillCount++
					}
				}
				// 进行字段的解析,标号为17
				var Language [3]conf.LanguageElement
				mode = 17
				var langCount = 0
				LanguageField := connet.GetFieldName(conf.StructMap[mode])
				Language1InfoStr := connet.Split(personInfoStr.LanguageSkill1)
				Language1Struct := connet.SetValueToStruct(LanguageField, Language1InfoStr, mode)
				Language2InfoStr := connet.Split(personInfoStr.SpecialSkill2)
				Language2Struct := connet.SetValueToStruct(LanguageField, Language2InfoStr, mode)
				Language3InfoStr := connet.Split(personInfoStr.SpecialSkill3)
				Language3Struct := connet.SetValueToStruct(LanguageField, Language3InfoStr, mode)
				Language[0] = *Language1Struct.(*conf.LanguageElement)
				Language[1] = *Language2Struct.(*conf.LanguageElement)
				Language[2] = *Language3Struct.(*conf.LanguageElement)
				for i := 0; i < 3; i++ {
					if Language[i].LanguageType != "" {
						langCount++
					}
				}

				// 进行字段的解析,标号为18
				var Train [3]conf.TrainingElement
				mode = 18
				var trainCount = 0
				TrainField := connet.GetFieldName(conf.StructMap[mode])
				Train1InfoStr := connet.Split(personInfoStr.TrainSkill1)
				Train1Struct := connet.SetValueToStruct(TrainField, Train1InfoStr, mode)
				Train2InfoStr := connet.Split(personInfoStr.TrainSkill2)
				Train2Struct := connet.SetValueToStruct(TrainField, Train2InfoStr, mode)
				Train3InfoStr := connet.Split(personInfoStr.TrainSkill3)
				Train3Struct := connet.SetValueToStruct(TrainField, Train3InfoStr, mode)
				Train[0] = *Train1Struct.(*conf.TrainingElement)
				Train[1] = *Train2Struct.(*conf.TrainingElement)
				Train[2] = *Train3Struct.(*conf.TrainingElement)
				for i := 0; i < 3; i++ {
					if Train[i].Course != "" {
						trainCount++
					}
				}

				// 进行字段的解析,标号为19
				var Certificate [3]conf.CertificateElement
				mode = 19
				var certiCount = 0
				CertificateField := connet.GetFieldName(conf.StructMap[mode])
				Certificate1InfoStr := connet.Split(personInfoStr.TrainSkill1)
				Certificate1Struct := connet.SetValueToStruct(CertificateField, Certificate1InfoStr, mode)
				Certificate2InfoStr := connet.Split(personInfoStr.TrainSkill2)
				Certificate2Struct := connet.SetValueToStruct(CertificateField, Certificate2InfoStr, mode)
				Certificate3InfoStr := connet.Split(personInfoStr.TrainSkill3)
				Certificate3Struct := connet.SetValueToStruct(CertificateField, Certificate3InfoStr, mode)
				Certificate[0] = *Certificate1Struct.(*conf.CertificateElement)
				Certificate[1] = *Certificate2Struct.(*conf.CertificateElement)
				Certificate[2] = *Certificate3Struct.(*conf.CertificateElement)
				for i := 0; i < 3; i++ {
					if Certificate[i].CertificateType != "" {
						certiCount++
					}
				}
				selfEvaluation := personInfoStr.SelfEvaluation

				context.JSON(http.StatusOK, gin.H{
					"name":            BasicInfo.Name,
					"age":             25,
					"gender":          BasicInfo.Gender,
					"area":            BasicInfo.Area,
					"status":          BasicInfo.Identity,
					"phone":           BasicInfo.Phone,
					"email":           BasicInfo.Email,
					"expectJob":       ExpectJob[0:jobCount],
					"educationList":   Education[0:eduCount],
					"workList":        Work[0:workCount],
					"itemList":        Item[0:itemCount],
					"schoolList":      School[0:schoolCount],
					"skillList":       Skill[0:skillCount],
					"languageList":    Language[0:langCount],
					"trainList":       Train[0:trainCount],
					"certificateList": Certificate[0:certiCount],
					"selfEvaluation":  selfEvaluation,
				})
			} else {
				// 更新PersonPost信息的状态值
				status, err := strconv.Atoi(clickStatus)
				//DB.Model(&personPost).Update("status", status)
				// 有多个选择条件之时最好按照下面方式进行更新，避免陷入坑
				DB.Model(&conf.PersonPost{}).Where("jid=? and pid=?", jid, pid).Update("status", status)
				if err != nil {
					context.JSON(http.StatusOK, gin.H{
						"msg": "更新失败",
					})
				} else {
					context.JSON(http.StatusOK, gin.H{
						"msg": "更新成功",
					})
				}
			}

		}
		//对邀请面试的操作
		viewTime := context.PostForm("viewTime")
		viewPlace := context.PostForm("viewPlace")
		viewTel := context.PostForm("viewTel")
		var Person conf.PersonPost
		Person.ViewTime = viewTime
		Person.ViewPlace = viewPlace
		Person.ViewTel = viewTel
		DB.Model(&conf.PersonPost{}).Where("cid = ? AND jid = ?", cid, jid).Update(&Person)
	}
}
