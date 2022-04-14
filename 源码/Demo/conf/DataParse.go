package conf

import (
	"encoding/json"
	"time"
)

var StructMap = map[int]interface{}{
	1:  &PrivateAccount{},
	2:  &BusinessAccount{},
	3:  &CompanyInfo{},
	4:  &JobInfo{},
	5:  &PersonPost{},
	6:  &PersonInfo{},
	7:  &PersonInfoStr{},
	8:  &BasicInfoElement{},
	9:  &ExpectJobElement{},
	10: &EducationElement{},
	11: &WorkExperienceElement{},
	12: &ItemElement{},
	13: &SchoolTimeElement{},
	14: &SchoolPostElement{},
	15: &SkillGoodElement{},
	16: &SpecialElement{},
	17: &LanguageElement{},
	18: &TrainingElement{},
	19: &CertificateElement{},
}

// PrivateAccount 用来存放个人用户的账号(电话号码tel), 编号为1
type PrivateAccount struct {
	Pid       string `json:"pid"`
	Telephone string `json:"telephone"`
}

//BusinessAccount 用来存放企业用户的账号(电话号码tel、密码password)，编号为2
type BusinessAccount struct {
	Cid       string `json:"cid"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

// CompanyInfo 定义公司的基础信息结构体，标号为3
type CompanyInfo struct {
	Cid         string    `json:"cid"`
	ShortName   string    `json:"shortname"`
	FullName    string    `json:"fullName"`
	CompanyType string    `json:"companyType"`
	Address     string    `json:"address"`
	Industry    string    `json:"industry"`
	Scale       string    `json:"scale"`
	LogoUrl     string    `json:"logoUrl"`
	Brief       string    `json:"brief"`
	CreatTime   time.Time `json:"createTime"`
	ChangeTime  time.Time `json:"changeTime"`
}

// JobInfo 定义职位的基础信息结构体，标号为4
type JobInfo struct {
	Jid        string    `json:"jid"`
	Title      string    `json:"title"`
	Company    string    `json:"company"`
	Cid        string    `json:"cid"`
	Nature     string    `json:"nature"`
	Salary     string    `json:"salary"`
	Education  string    `json:"education"`
	Experience string    `json:"experience"`
	Province   string    `json:"province"`
	Address    string    `json:"address"`
	Require    string    `json:"require"`
	Describe   string    `json:"describe"`
	CreatTime  time.Time `json:"createTime"`
	ChangeTime time.Time `json:"changeTime"`
}

// PersonPost 定义个人投递情况数据表，标号为5
//Status, 简历的四种状态，0：未查看， 1：拒绝，2：待定，3：发起面试
type PersonPost struct {
	Pid        string    `json:"pid"`
	Jid        string    `json:"Jid"`
	Cid        string    `json:"cid"`
	Status     int       `json:"status"`
	Title      string    `json:"title"`
	Address    string    `json:"address"`
	Experience string    `json:"experience"`
	Education  string    `json:"education"`
	Salary     string    `json:"salary"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	ViewTime   string  	 `json:"view_time"`
	ViewPlace  string    `json:"view_place"`
	ViewTel    string    `json:"view_tel"`
	CreatTime  time.Time `json:"createTime"`
	ChangeTime time.Time `json:"changeTime"`
}

// PersonInfo 个人简历数据体,因为包含结构体，不能直接用来生成数据库中的表格，标号为6
type PersonInfo struct {
	BasicInfo      BasicInfoElement `json:"basicInfo"`
	ExpectJob1     ExpectJobElement
	ExpectJob2     ExpectJobElement
	ExpectJob3     ExpectJobElement
	Education1     EducationElement
	Education2     EducationElement
	Education3     EducationElement
	Work1          WorkExperienceElement
	Work2          WorkExperienceElement
	Work3          WorkExperienceElement
	Item1          ItemElement
	Item2          ItemElement
	Item3          ItemElement
	SchoolTime     SchoolTimeElement
	SkillGood      SkillGoodElement
	SelfEvaluation string
	CreatTime      time.Time `json:"createTime"`
	ChangeTime     time.Time `json:"changeTime"`
}

// PersonInfoStr 收录拼和之后的个人简历数据，标号为7
type PersonInfoStr struct {
	Pid            string    `json:"pid"`
	BasicInfo      string    `json:"basicInfo"`
	ExpectJob1     string    `json:"expectJob1"`
	ExpectJob2     string    `json:"expectJob2"`
	ExpectJob3     string    `json:"expectJob3"`
	Education1     string    `json:"education1"`
	Education2     string    `json:"education2"`
	Education3     string    `json:"education3"`
	Work1          string    `json:"work1"`
	Work2          string    `json:"work2"`
	Work3          string    `json:"work3"`
	Item1          string    `json:"item1"`
	Item2          string    `json:"item2"`
	Item3          string    `json:"item3"`
	SchoolPost1    string    `json:"schoolPost1"`
	SchoolPost2    string    `json:"schoolPost2"`
	SchoolPost3    string    `json:"schoolPost3"`
	SpecialSkill1  string    `json:"specialSkill1"`
	SpecialSkill2  string    `json:"specialSkill2"`
	SpecialSkill3  string    `json:"specialSkill3"`
	LanguageSkill1 string    `json:"languageSkill1"`
	LanguageSkill2 string    `json:"languageSkill2"`
	LanguageSkill3 string    `json:"languageSkill3"`
	TrainSkill1    string    `json:"trainSkill1"`
	TrainSkill2    string    `json:"trainSkill2"`
	TrainSkill3    string    `json:"trainSkill3"`
	Certificate1   string    `json:"certificate1"`
	Certificate2   string    `json:"certificate2"`
	Certificate3   string    `json:"certificate3"`
	SelfEvaluation string    `json:"selfEvaluation"`
	CreatTime      time.Time `json:"createTime"`
	ChangeTime     time.Time `json:"changeTime"`
}

// BasicInfoElement 个人基础信息，标号为8
type BasicInfoElement struct {
	Phone    string `form:"phone"`
	Name     string `form:"name"`
	Gender   string `form:"gender"`
	Identity string `form:"identity"`
	Birth    string `form:"birth"`
	Area     string `form:"area"`
	Email    string `form:"email"`
}

// ExpectJobElement 求职意向表，标号为9
type ExpectJobElement struct {
	ExpectedOccupation    string `form:"expected_occupation"`
	ExpectedCity          string `form:"expected_city"`
	SalaryRequirementsMin string `form:"salary_requirements_min"`
	SalaryRequirementsMax string `form:"salary_requirements_max"`
	JobType               string `form:"nature_of_works[]"`
	ExpectTime            string `form:"arrival_time"`
}

// EducationElement 教育经历，标号为10
type EducationElement struct {
	SchoolName                  string `form:"school_name"`
	HighestEducation            string `form:"highest_education"`
	MajorStudied                string `form:"major_studied"`
	SchoolTimeStart             string `form:"school_time_start"`
	SchoolTimeEnd               string `form:"school_time_start_end"`
	CertificateOfAcademicDegree string `form:"certificate_of_academic_degree"`
}

// WorkExperienceElement 工作经历，标号为11
type WorkExperienceElement struct {
	Company       string `form:"company"`
	Job           string `form:"job"`
	InTime        string `form:"inTime"`
	EndTime       string `form:"end_time"`
	CurrentSalary string `form:"current_monthly_salary"`
	WorkType      string `form:"workType[]"`
	Describe      string `form:"describe"`
}

// ItemElement 项目经验，标号为12
type ItemElement struct {
	Title    string `form:"title"`
	InTime   string `form:"inTime"`
	EndTime  string `form:"end_time"`
	Company  string `form:"company"`
	Describe string `form:"describe"`
}

// SchoolTimeElement 校内表现，标号13
type SchoolTimeElement struct {
	SchoolPost1 SchoolPostElement
	SchoolPost2 SchoolPostElement
	SchoolPost3 SchoolPostElement
}

// SchoolPostElement 校内职务，标号14
type SchoolPostElement struct {
	Award    string `form:"award"`
	Post     string `form:"post"`
	Describe string `form:"describe"`
}

// SkillGoodElement 技能特长，标号15
type SkillGoodElement struct {
	SpecialSkill1     SpecialElement
	SpecialSkill2     SpecialElement
	SpecialSkill3     SpecialElement
	LanguageSkill1    LanguageElement
	LanguageSkill2    LanguageElement
	LanguageSkill3    LanguageElement
	TrainSkill1       TrainingElement
	TrainSkill2       TrainingElement
	TrainSkill3       TrainingElement
	CertificateSkill1 CertificateElement
	CertificateSkill2 CertificateElement
	CertificateSkill3 CertificateElement
}

// SpecialElement 标号为16
type SpecialElement struct {
	Skill  string `form:"skill"`
	InTime string `form:"inTime"`
}

// LanguageElement 标号为17
type LanguageElement struct {
	LanguageType string `form:"languageType"`
	InTime       string `form:"inTime"`
}

// TrainingElement 标号为18
type TrainingElement struct {
	Organization string `form:"organization"`
	Course       string `form:"course"`
	Period       string `form:"period"`
	Describe     string `form:"describe"`
}

// CertificateElement 标号为19
type CertificateElement struct {
	CertificateType string `form:"certificate"`
	ObtainTime      string `form:"obtainTime"`
}

type SelfEvaluation struct {
	SelfEvaluation string `form:"selfEvaluation"`
}

type Modify struct {
	Modify string `form:"modify"`
}

func (res BasicInfoElement) JSONBlindStruct(a string) {
	json.Unmarshal([]byte(a), &res)
}
