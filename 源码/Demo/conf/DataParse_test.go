package conf

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// Test_CompanyInfo: 是一个单纯的结构体，不包含结构体嵌套，它的第一种赋值语句如下：
func Test_CompanyInfo_1(t *testing.T) {
	var companyInfo CompanyInfo
	companyInfo.Cid = "12345"
	companyInfo.ShortName = "象大"
	companyInfo.FullName = "象大科技有限公司"
	companyInfo.CompanyType = "民企"
	companyInfo.Address = "南通市海安市中科创新广场"
	companyInfo.Industry = "互联网"
	companyInfo.Scale = "50人以上"
	companyInfo.LogoUrl = "https://www.smarteleph.com"
	companyInfo.Brief = "人才先育，产业升级"
	companyInfo.CreatTime = time.Now()
	companyInfo.ChangeTime = time.Now()
	if companyInfo.Brief != "人才先育，产业升级" { //try a unit test on function
		t.Error("比对通不过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") // 记录一些你期望记录的信息
	}
	// 将结构体json marshal化，然后字符串形式输出
	keyBytes, err := json.Marshal(companyInfo)
	fmt.Printf("keyBytes:=%s\n", keyBytes)
	fmt.Println("struct is: ", companyInfo)
	if err != nil {
		fmt.Println("json unmarshal err")
	}
	//fmt.Println(string(keyBytes))

	// 将Marshal后的结构体可以以json.Unmarshal的形式绑定到另一个结构体中
	CompanyUse := CompanyInfo{}
	err = json.Unmarshal(keyBytes, &CompanyUse)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}

	// 两个不同的struct的实例能不能比较 ==  != 答案：可以能、也可以不能，自行百度
	if CompanyUse.Address != companyInfo.Address {
		t.Error("两个结构体不匹配") // 如果不是如预期的那么就报错
	} else {
		t.Log("第二个测试通过了") // 记录一些你期望记录的信息
	}
	//fmt.Printf("basic:%v\n", CompanyUse)
}

// Test_CompanyInfo: 是一个单纯的结构体，不包含结构体嵌套，它的第一种赋值语句如下：
func Test_CompanyInfo_2(t *testing.T) {
	companyInfo := CompanyInfo{
		Cid:         "12345",
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
	if companyInfo.Brief != "人才先育，产业升级" { //try a unit test on function
		t.Error("比对通不过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") // 记录一些你期望记录的信息
	}
	// 将结构体json marshal化，然后字符串形式输出
	keyBytes, err := json.Marshal(companyInfo)
	if err != nil {
		fmt.Println("json unmarshal err")
	}
	//fmt.Println(string(keyBytes))

	// 将Marshal后的结构体可以以json.Unmarshal的形式绑定到另一个结构体中
	CompanyUse := CompanyInfo{}
	err = json.Unmarshal(keyBytes, &CompanyUse)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}

	// 两个不同的struct的实例能不能比较 ==  != 答案：可以能、也可以不能，自行百度
	if CompanyUse.Address != companyInfo.Address {
		t.Error("两个结构体不匹配") // 如果不是如预期的那么就报错
	} else {
		t.Log("第二个测试通过了") // 记录一些你期望记录的信息
	}
	//fmt.Printf("basic:%v\n", CompanyUse)
}

// SchoolTimeElement: 是一个有个两层结构的结构体，其赋值语句为：
func Test_SchoolTimeElement_1(t *testing.T) {

	schoolPost1 := SchoolPostElement{
		Post:     "外联部干事",
		Describe: "2011.6-2011.5,外联部干事，负责跑腿",
	}
	schoolPost2 := SchoolPostElement{
		Post:     "外联部部长",
		Describe: "2012.6-2013.5,担任外联部部长，负责拉赞助",
	}
	schoolPost3 := SchoolPostElement{
		Post:     "学生会主席",
		Describe: "2013.5-2014.5,担任学生会主席，统筹全局",
	}
	SchoolTime := SchoolTimeElement{
		SchoolPost1: schoolPost1,
		SchoolPost2: schoolPost2,
		SchoolPost3: schoolPost3,
	}
	if SchoolTime.SchoolPost3.Post != "学生会主席" { //try a unit test on function
		t.Error("比对通不过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第四个测试通过了") // 记录一些你期望记录的信息
	}
}

// SchoolTimeElement: 是一个有个两层结构的结构体，其赋值语句为：
func Test_JobElement_1(t *testing.T) {

	schoolPost1 := SchoolPostElement{
		Post:     "外联部干事",
		Describe: "2011.6-2011.5,外联部干事，负责跑腿",
	}
	schoolPost2 := SchoolPostElement{
		Post:     "外联部部长",
		Describe: "2012.6-2013.5,担任外联部部长，负责拉赞助",
	}
	schoolPost3 := SchoolPostElement{
		Post:     "学生会主席",
		Describe: "2013.5-2014.5,担任学生会主席，统筹全局",
	}
	SchoolTime := SchoolTimeElement{
		SchoolPost1: schoolPost1,
		SchoolPost2: schoolPost2,
		SchoolPost3: schoolPost3,
	}
	if SchoolTime.SchoolPost3.Post != "学生会主席" { //try a unit test on function
		t.Error("比对通不过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第四个测试通过了") // 记录一些你期望记录的信息
	}
}
