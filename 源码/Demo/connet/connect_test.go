package connet

import (
	"Demo/conf"
	"Demo/setting"
	"fmt"
	"testing"
	"time"
)

func Test_connect(t *testing.T) {
	configPath := "../conf/config.ini"
	var err error
	if err = setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 需要后台开启Mysql,开启的命令，找到mysql安装路径的bin文件，./mysql -u{#usenamr} -p{#password}
	DB, err = InitMySQL(setting.Conf.MySQLConfig)
	defer DB.Close()
	if err != nil {
		t.Error("数据库连接失败")
		panic(err)
	} else {
		t.Log("第一个测试通过了")
	}
	// 需要后台开启redis,开启的命令，redis-cli -h {#host} -p {#port}
	// 关闭客户端 ./redis-cli shutdown
	PoolConn, err = InitRedis(setting.Conf.RedisConfig)
	if err != nil {
		t.Error("Redis连接失败")
		panic(err)
	} else {
		t.Log("第二个测试通过了")
	}
	defer PoolConn.Close()
}

func Test_Merge(t *testing.T) {
	configPath := "../conf/config.ini"
	var err error
	if err = setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	DB, err = InitMySQL(setting.Conf.MySQLConfig)
	defer DB.Close()
	if err != nil {
		t.Error("数据库连接失败")
		panic(err)
	} else {
		t.Log("第一个测试通过")
	}
	// 结构体结构变量赋值，直接写入mysql
	basicInfoElement := conf.BasicInfoElement{
		Name:     "12345",
		Gender:   "男",
		Identity: "学生",
		Birth:    "1999年2月5号",
		Area:     "中科创新广场",
		Phone:    "17524242424",
		Email:    "67555558",
	}
	//var BasicInfo *conf.BasicInfoElement
	//res1 := GetFieldName(BasicInfo)
	res1 := GetFieldName(conf.StructMap[8])

	fmt.Println(res1)
	if res1[0] != "Phone" {
		t.Error("获取结构体字段失败")
	} else {
		t.Log("第二个测试通过")
	}
	res2 := GetValueName(basicInfoElement)
	fmt.Println(res2)
	if res2[0] != "17524242424" {
		t.Error("获取结构体字段失败")
		t.Logf("res[0]=:%s", res2[0])
	} else {
		t.Log("第三个测试通过")
	}
	MergeString := Merge(basicInfoElement)
	fmt.Println(MergeString)
	ConstString := "17524242424#12345#男#学生#1999年2月5号#中科创新广场#67555558#"
	if ConstString != MergeString {
		t.Error("字符的拼接出错")
	} else {
		t.Log("第四个测试通过")
	}
	SelectStruct := 8
	// 返回的是一个interface{},必须类型转化
	ReceiverStruct := SetValueToStruct(res1, res2, SelectStruct)
	var newStruct *conf.BasicInfoElement
	newStruct = ReceiverStruct.(*conf.BasicInfoElement)
	fmt.Println(newStruct.Name)
	if newStruct.Gender != "男" {
		t.Error("结构体赋值失败")
	} else {
		t.Log("第五个测试通过")
	}
	// 把模型与数据库中的表对应起来
	DB.AutoMigrate(&conf.CompanyInfo{})
	companyInfo := conf.CompanyInfo{
		Cid:         "123456",
		ShortName:   "象大",
		FullName:    "象大科技有限公司",
		CompanyType: "民企",
		Address:     "南通市海安市中科创新广场",
		Industry:    "互联网",
		Scale:       "50人以上",
		LogoUrl:     "https://www.smarteleph.com",
		Brief:       "人才先育，产业升级",
		CreatTime:   time.Now(),
		ChangeTime:  time.Now(),
	}
	DB.Save(&companyInfo)
	// 读取
	var TestCompanyInfo conf.CompanyInfo
	DB.Where("Cid = ?", "123456").Find(&TestCompanyInfo)
	if TestCompanyInfo.Industry != companyInfo.Industry {
		t.Error("数据库读取和写入不一致")
		panic(err)
	} else {
		t.Log("第六个测试通过")
	}
	DB.Delete(&TestCompanyInfo)
}

func Test_PersonInfoSTr(t *testing.T) {
	configPath := "../conf/config.ini"
	var err error
	if err = setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	DB, err = InitMySQL(setting.Conf.MySQLConfig)
	defer DB.Close()
	if err != nil {
		t.Error("数据库连接失败")
		panic(err)
	} else {
		t.Log("第一个测试通过")
	}
	// 把模型与数据库中的表对应起来
	DB.AutoMigrate(&conf.PersonInfoStr{})
	var personInfoStr conf.PersonInfoStr
	personInfoStr.Pid = "anytime12345"
	BasicInfo := conf.BasicInfoElement{
		Phone:    "1311032245",
		Name:     "eospeng",
		Gender:   "男",
		Identity: "职场人",
		Birth:    "2011-09-03",
		Area:     "江苏南通",
		Email:    "763231448@qq.com",
	}
	personInfoStr.BasicInfo = Merge(BasicInfo)

	var ExpectInfo, ExpectInfo2 conf.ExpectJobElement
	ExpectInfo.ExpectedOccupation = "总监"
	ExpectInfo.ExpectedCity = "上海"
	ExpectInfo.SalaryRequirementsMin = "100000"
	ExpectInfo.SalaryRequirementsMax = "无限"
	ExpectInfo.JobType = "全职，兼职"
	ExpectInfo.ExpectTime = "一个月内"

	ExpectInfo2.ExpectedOccupation = "小冰"
	ExpectInfo2.ExpectedCity = "成都"
	ExpectInfo2.SalaryRequirementsMin = "2000"
	ExpectInfo2.SalaryRequirementsMax = "10000"
	ExpectInfo2.JobType = "兼职"
	ExpectInfo2.ExpectTime = "三个月内"

	personInfoStr.ExpectJob1 = Merge(ExpectInfo)
	personInfoStr.ExpectJob2 = Merge(ExpectInfo2)
	personInfoStr.ExpectJob3 = Merge(ExpectInfo)
	EduInfo := conf.EducationElement{
		SchoolName:                  "南通理工大学",
		HighestEducation:            "Doctor",
		MajorStudied:                "Science",
		SchoolTimeStart:             "2011年9月8日",
		SchoolTimeEnd:               "2015年7月5日",
		CertificateOfAcademicDegree: "Master",
	}
	personInfoStr.Education1 = Merge(EduInfo)
	personInfoStr.Education2 = Merge(EduInfo)
	personInfoStr.SelfEvaluation = "天生我才必有用，千金散尽还复来。哈哈，这就是我的介绍"
	personInfoStr.CreatTime = time.Now()
	personInfoStr.ChangeTime = time.Now()
	DB.Save(&personInfoStr)
	fmt.Println("personInfoStr 写入成功")
}

func TestPostStatus(t *testing.T) {
	configPath := "../conf/config.ini"
	var err error
	if err = setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	DB, err = InitMySQL(setting.Conf.MySQLConfig)
	defer DB.Close()
	if err != nil {
		t.Error("数据库连接失败")
		panic(err)
	} else {
		t.Log("第一个测试通过")
	}

	var JobList []conf.JobInfo
	var cid = "c9945fe4260aa2c3"
	var pid = "225741f979f1189e62aa1ce2662a7b1c"
	DB.Where("cid=?", cid).Find(&JobList)
	var TitleList = [4]string{"产品经理", "JAVA 开发", "GO工程师", "PHP工程师"}
	var NameList = [4]string{"张三", "李四", "王五", "赵六"}
	var personPost conf.PersonPost
	for i := 0; i < len(JobList); i++ {
		personPost.Pid = pid
		personPost.Cid = cid
		personPost.Jid = JobList[i].Jid
		personPost.Status = i % 4
		personPost.Title = TitleList[personPost.Status]
		personPost.Experience = "无经验"
		personPost.Education = "本科"
		personPost.Age = i
		personPost.Address = "北京"
		personPost.Salary = "8000-10000"
		personPost.Name = NameList[personPost.Status]
		personPost.CreatTime = time.Now()
		personPost.ChangeTime = time.Now()
		DB.Save(&personPost)
		fmt.Printf("personPost, index=[%d]\n", i+1)
	}
}
