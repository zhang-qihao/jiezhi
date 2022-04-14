package person

import (
	"Demo/conf"
	"Demo/connet"
	"Demo/routers/Login"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GenerateResume，根据pid 去数据库中查询用户的信息，然后调用resumePreview.html模版进行渲染结果
func GetPerson(DB *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

		fmt.Println("GetPerson 函数内部")
		pid := context.Query("pid")

		var personInfoStr conf.PersonInfoStr
		DB.Where("Pid = ?", pid).Find(&personInfoStr) //数据库查找
		// 进行字段的解析,标号为8
		var BasicInfo *conf.BasicInfoElement
		basicField := connet.GetFieldName(BasicInfo)
		basicInfoStr := connet.Split(personInfoStr.BasicInfo)
		BasicStruct := connet.SetValueToStruct(basicField, basicInfoStr, 8)
		BasicInfo = BasicStruct.(*conf.BasicInfoElement)

		// 进行字段的解析,标号为9
		var ExpectJob [3]conf.ExpectJobElement
		JobField := connet.GetFieldName(ExpectJob[0])
		Job1InfoStr := connet.Split(personInfoStr.ExpectJob1)
		Job1Struct := connet.SetValueToStruct(JobField, Job1InfoStr, 9)
		Job2InfoStr := connet.Split(personInfoStr.ExpectJob2)
		Job2Struct := connet.SetValueToStruct(JobField, Job2InfoStr, 9)
		Job3InfoStr := connet.Split(personInfoStr.ExpectJob3)
		Job3Struct := connet.SetValueToStruct(JobField, Job3InfoStr, 9)
		ExpectJob[0] = *Job1Struct.(*conf.ExpectJobElement)
		ExpectJob[1] = *Job2Struct.(*conf.ExpectJobElement)
		ExpectJob[2] = *Job3Struct.(*conf.ExpectJobElement)

		// 进行字段的解析,标号为10
		var Education [3]conf.EducationElement
		EduField := connet.GetFieldName(&conf.EducationElement{})
		Edu1InfoStr := connet.Split(personInfoStr.Education1)
		Education[0] = *connet.SetValueToStruct(EduField, Edu1InfoStr, 10).(*conf.EducationElement)
		Edu2InfoStr := connet.Split(personInfoStr.Education2)
		Education[1] = *connet.SetValueToStruct(EduField, Edu2InfoStr, 10).(*conf.EducationElement)
		Edu3InfoStr := connet.Split(personInfoStr.Education3)
		Edu3Struct := connet.SetValueToStruct(EduField, Edu3InfoStr, 10)
		Education[2] = *Edu3Struct.(*conf.EducationElement)

		// 进行字段的解析,标号为11
		var Work [3]conf.WorkExperienceElement
		WorkField := connet.GetFieldName(Work[0])
		Work1InfoStr := connet.Split(personInfoStr.Work1)
		Work1Struct := connet.SetValueToStruct(WorkField, Work1InfoStr, 11)
		Work2InfoStr := connet.Split(personInfoStr.Work2)
		Work2Struct := connet.SetValueToStruct(WorkField, Work2InfoStr, 11)
		Work3InfoStr := connet.Split(personInfoStr.Work3)
		Work3Struct := connet.SetValueToStruct(WorkField, Work3InfoStr, 11)
		Work[0] = *Work1Struct.(*conf.WorkExperienceElement)
		Work[1] = *Work2Struct.(*conf.WorkExperienceElement)
		Work[2] = *Work3Struct.(*conf.WorkExperienceElement)

		// 进行字段的解析,标号为12
		var Item [3]conf.ItemElement
		ItemField := connet.GetFieldName(Item[0])
		Item1InfoStr := connet.Split(personInfoStr.Item1)
		Item1Struct := connet.SetValueToStruct(ItemField, Item1InfoStr, 12)
		Item2InfoStr := connet.Split(personInfoStr.Item2)
		Item2Struct := connet.SetValueToStruct(ItemField, Item2InfoStr, 12)
		Item3InfoStr := connet.Split(personInfoStr.Item3)
		Item3Struct := connet.SetValueToStruct(ItemField, Item3InfoStr, 12)
		Item[0] = *Item1Struct.(*conf.ItemElement)
		Item[1] = *Item2Struct.(*conf.ItemElement)
		Item[2] = *Item3Struct.(*conf.ItemElement)

		// 进行字段的解析,标号为14
		var School [3]conf.SchoolPostElement
		SchoolField := connet.GetFieldName(School[0])
		School1InfoStr := connet.Split(personInfoStr.SchoolPost1)
		School1Struct := connet.SetValueToStruct(SchoolField, School1InfoStr, 14)
		School2InfoStr := connet.Split(personInfoStr.SchoolPost2)
		School2Struct := connet.SetValueToStruct(SchoolField, School2InfoStr, 14)
		School3InfoStr := connet.Split(personInfoStr.SchoolPost3)
		School3Struct := connet.SetValueToStruct(SchoolField, School3InfoStr, 14)
		School[0] = *School1Struct.(*conf.SchoolPostElement)
		School[1] = *School2Struct.(*conf.SchoolPostElement)
		School[2] = *School3Struct.(*conf.SchoolPostElement)

		// 进行字段的解析,标号为16
		var Skill [3]conf.SpecialElement
		SkillField := connet.GetFieldName(Skill[0])
		Skill1InfoStr := connet.Split(personInfoStr.SpecialSkill1)
		Skill1Struct := connet.SetValueToStruct(SkillField, Skill1InfoStr, 16)
		Skill2InfoStr := connet.Split(personInfoStr.SpecialSkill2)
		Skill2Struct := connet.SetValueToStruct(SkillField, Skill2InfoStr, 16)
		Skill3InfoStr := connet.Split(personInfoStr.SpecialSkill3)
		Skill3Struct := connet.SetValueToStruct(SkillField, Skill3InfoStr, 16)
		Skill[0] = *Skill1Struct.(*conf.SpecialElement)
		Skill[1] = *Skill2Struct.(*conf.SpecialElement)
		Skill[2] = *Skill3Struct.(*conf.SpecialElement)

		// 进行字段的解析,标号为17
		var Language [3]conf.LanguageElement
		LanguageField := connet.GetFieldName(Language[0])
		Language1InfoStr := connet.Split(personInfoStr.LanguageSkill1)
		Language1Struct := connet.SetValueToStruct(LanguageField, Language1InfoStr, 17)
		Language2InfoStr := connet.Split(personInfoStr.LanguageSkill2)
		Language2Struct := connet.SetValueToStruct(LanguageField, Language2InfoStr, 17)
		Language3InfoStr := connet.Split(personInfoStr.LanguageSkill3)
		Language3Struct := connet.SetValueToStruct(LanguageField, Language3InfoStr, 17)
		Language[0] = *Language1Struct.(*conf.LanguageElement)
		Language[1] = *Language2Struct.(*conf.LanguageElement)
		Language[2] = *Language3Struct.(*conf.LanguageElement)

		// 进行字段的解析,标号为18
		var Train [3]conf.TrainingElement
		TrainField := connet.GetFieldName(Train[0])
		Train1InfoStr := connet.Split(personInfoStr.TrainSkill1)
		Train1Struct := connet.SetValueToStruct(TrainField, Train1InfoStr, 18)
		Train2InfoStr := connet.Split(personInfoStr.TrainSkill2)
		Train2Struct := connet.SetValueToStruct(TrainField, Train2InfoStr, 18)
		Train3InfoStr := connet.Split(personInfoStr.TrainSkill3)
		Train3Struct := connet.SetValueToStruct(TrainField, Train3InfoStr, 18)
		Train[0] = *Train1Struct.(*conf.TrainingElement)
		Train[1] = *Train2Struct.(*conf.TrainingElement)
		Train[2] = *Train3Struct.(*conf.TrainingElement)

		// 进行字段的解析,标号为18
		var Certificate [3]conf.CertificateElement
		CertificateField := connet.GetFieldName(Certificate[0])
		Certificate1InfoStr := connet.Split(personInfoStr.Certificate1)
		Certificate1Struct := connet.SetValueToStruct(CertificateField, Certificate1InfoStr, 19)
		Certificate2InfoStr := connet.Split(personInfoStr.Certificate2)
		Certificate2Struct := connet.SetValueToStruct(CertificateField, Certificate2InfoStr, 19)
		Certificate3InfoStr := connet.Split(personInfoStr.Certificate3)
		Certificate3Struct := connet.SetValueToStruct(CertificateField, Certificate3InfoStr, 19)
		Certificate[0] = *Certificate1Struct.(*conf.CertificateElement)
		Certificate[1] = *Certificate2Struct.(*conf.CertificateElement)
		Certificate[2] = *Certificate3Struct.(*conf.CertificateElement)

		SelfEvaluation := personInfoStr.SelfEvaluation
		ChangeTime := personInfoStr.ChangeTime

		authorized := Login.Getsess(context, "authorizedId")
		if authorized == pid {
			context.HTML(http.StatusOK, "resumeModify.html", gin.H{
				//基础信息
				"phone":    BasicInfo.Phone,
				"name":     BasicInfo.Name,
				"gender":   BasicInfo.Gender,
				"identity": BasicInfo.Identity,
				"birth":    BasicInfo.Birth,
				"area":     BasicInfo.Area,
				"email":    BasicInfo.Email,
				//期望职业
				"expected_occupation1":     ExpectJob[0].ExpectedOccupation,
				"expected_city1":           ExpectJob[0].ExpectedCity,
				"salary_requirements_min1": ExpectJob[0].SalaryRequirementsMin,
				"salary_requirements_max1": ExpectJob[0].SalaryRequirementsMax,
				"nature_of_work1":          ExpectJob[0].JobType,
				"arrival_time1":            ExpectJob[0].ExpectTime,
				//教育经历
				"school_name1":       Education[0].SchoolName,
				"major_studied1":     Education[0].MajorStudied,
				"highest_education1": Education[0].HighestEducation,
				"school_time1":       Education[0].SchoolTimeStart + "至" + Education[0].SchoolTimeEnd,
				"school_name2":       Education[1].SchoolName,
				"major_studied2":     Education[1].MajorStudied,
				"highest_education2": Education[1].HighestEducation,
				"school_time2":       Education[1].SchoolTimeStart + "至" + Education[1].SchoolTimeEnd,
				"school_name3":       Education[2].SchoolName,
				"major_studied3":     Education[2].MajorStudied,
				"highest_education3": Education[2].HighestEducation,
				"school_time3":       Education[2].SchoolTimeStart + "至" + Education[2].SchoolTimeEnd,
				//工作经历
				"corporate_name1":         Work[0].Company,
				"job_title1":              Work[0].Job,
				"time_of_employment1":     Work[0].InTime,
				"current_monthly_salary1": Work[0].CurrentSalary,
				"type_of_work1":           Work[0].WorkType,
				"job_description1":        Work[0].Describe,
				"corporate_name2":         Work[1].Company,
				"job_title2":              Work[1].Job,
				"time_of_employment2":     Work[1].InTime,
				"current_monthly_salary2": Work[1].CurrentSalary,
				"type_of_work2":           Work[1].WorkType,
				"job_description2":        Work[1].Describe,
				"corporate_name3":         Work[2].Company,
				"job_title3":              Work[2].Job,
				"time_of_employment3":     Work[2].InTime,
				"current_monthly_salary3": Work[2].CurrentSalary,
				"type_of_work3":           Work[2].WorkType,
				"job_description3":        Work[2].Describe,
				//项目经历
				"project_name1":        Item[0].Title,
				"project_time1":        Item[0].InTime + "-" + Item[0].EndTime,
				"project_description1": Item[0].Describe,
				"project_name2":        Item[1].Title,
				"project_time2":        Item[1].InTime + "-" + Item[1].EndTime,
				"project_description2": Item[1].Describe,
				"project_name3":        Item[2].Title,
				"project_time3":        Item[2].InTime + "-" + Item[2].EndTime,
				"project_description3": Item[2].Describe,
				//校内表现
				"award1":            School[0].Award,
				"post1":             School[0].Post,
				"post_description1": School[0].Describe,
				"award2":            School[1].Award,
				"post2":             School[1].Post,
				"post_description2": School[1].Describe,
				"award3":            School[2].Award,
				"post3":             School[2].Post,
				"post_description3": School[2].Describe,
				//技能特长
				"skill_name1": Skill[0].Skill,
				"skill_time1": Skill[0].InTime,
				"skill_name2": Skill[1].Skill,
				"skill_time2": Skill[1].InTime,
				"skill_name3": Skill[2].Skill,
				"skill_time3": Skill[2].InTime,
				//语言特长
				"language_name1": Language[0].LanguageType,
				"language_time1": Language[0].InTime,
				"language_name2": Language[1].LanguageType,
				"language_time2": Language[1].InTime,
				"language_name3": Language[2].LanguageType,
				"language_time3": Language[2].InTime,
				//培训经历
				"training_institution1": Train[0].Organization,
				"training_courses1":     Train[0].Course,
				"training_description1": Train[0].Describe,
				"training_institution2": Train[1].Organization,
				"training_courses2":     Train[1].Course,
				"training_description2": Train[1].Describe,
				"training_institution3": Train[2].Organization,
				"training_courses3":     Train[2].Course,
				"training_description3": Train[2].Describe,
				//获得证书
				"certificate_name1": Certificate[0].CertificateType,
				"acquired_date1":    Certificate[0].ObtainTime,
				"certificate_name2": Certificate[1].CertificateType,
				"acquired_date2":    Certificate[1].ObtainTime,
				"certificate_name3": Certificate[2].CertificateType,
				"acquired_date3":    Certificate[2].ObtainTime,
				//自我评价
				"self_evaluation": SelfEvaluation,
				//更新时间
				"changeTime": ChangeTime,
			})
		} else {
			context.String(http.StatusOK, "未登录")
		}
	}
}
