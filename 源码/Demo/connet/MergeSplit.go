package connet

import (
	"Demo/conf"
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func generateJSONStr(a, b []string) string {
	if len(a) != len(b) {
		fmt.Println(" 请确保两个数组的长度一致")
	}
	fixLen := len(a)
	var buffer bytes.Buffer
	buffer.WriteString("`{")
	for i := 0; i < fixLen-1; i++ {
		buffer.WriteString("\"")
		buffer.WriteString(a[i])
		buffer.WriteString("\"")
		buffer.WriteString(":")
		buffer.WriteString("\"")
		buffer.WriteString(b[i])
		buffer.WriteString("\"")
		buffer.WriteString(",")
	}
	buffer.WriteString("\"")
	buffer.WriteString(a[fixLen-1])
	buffer.WriteString("\"")
	buffer.WriteString(":")
	buffer.WriteString("\"")
	buffer.WriteString(b[fixLen-1])
	buffer.WriteString("\"")
	buffer.WriteString("}`")
	return buffer.String()
}

func Merge(structName interface{}) string {
	t := reflect.ValueOf(structName)
	fieldNum := t.NumField()
	//FieldList := GetFieldName(structName)
	//fieldNum := len(FieldList)
	var buffer bytes.Buffer
	for i := 0; i < fieldNum; i++ {
		buffer.WriteString(t.Field(i).String())
		buffer.WriteString("#")
	}
	return buffer.String()
}

func Split(a string) []string {
	receiver := strings.Split(a, "#")
	return receiver
}

//GetFieldName 获取结构体中字段的名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

func GetValueName(structName interface{}) []string {
	t := reflect.ValueOf(structName)
	//FieldList := GetFieldName(structName)
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).String())
		//result = append(result, t.FieldByName(FieldList[i]).String())
	}
	return result
}
func SetValue(FieldList, ValueList []string, p interface{}) {
	lenValue := len(ValueList)
	defaultValue := ""
	v := reflect.ValueOf(p).Elem()
	fieldNum := v.NumField()
	for i := 0; i < fieldNum; i++ {
		if lenValue <= 1 {
			v.FieldByName(FieldList[i]).Set(reflect.ValueOf(defaultValue))
		} else {
			v.FieldByName(FieldList[i]).Set(reflect.ValueOf(ValueList[i]))
		}
	}
}

// SetValueToStruct 选择不同的结构体，传入结构体的Field和Value, 自动给结构体赋值
func SetValueToStruct(FieldList, ValueList []string, mode int) interface{} {
	flag := 1
	switch mode {
	case 1:
		p := &conf.PrivateAccount{}
		SetValue(FieldList, ValueList, p)
		return p
	case 2:
		p := &conf.BusinessAccount{}
		SetValue(FieldList, ValueList, p)
		return p
	case 3:
		p := &conf.CompanyInfo{}
		SetValue(FieldList, ValueList, p)
		return p
	case 4:
		p := &conf.JobInfo{}
		SetValue(FieldList, ValueList, p)
		return p
	case 5:
		p := &conf.PersonPost{}
		SetValue(FieldList, ValueList, p)
		return p
	case 6:
		p := &conf.PersonInfo{}
		SetValue(FieldList, ValueList, p)
		return p
	case 7:
		p := &conf.PersonInfoStr{}
		SetValue(FieldList, ValueList, p)
		return p
	case 8:
		p := &conf.BasicInfoElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 9:
		p := &conf.ExpectJobElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 10:
		p := &conf.EducationElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 11:
		p := &conf.WorkExperienceElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 12:
		p := &conf.ItemElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 13:
		p := &conf.SchoolTimeElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 14:
		p := &conf.SchoolPostElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 15:
		p := &conf.SkillGoodElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 16:
		p := &conf.SpecialElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 17:
		p := &conf.LanguageElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 18:
		p := &conf.TrainingElement{}
		SetValue(FieldList, ValueList, p)
		return p
	case 19:
		p := &conf.CertificateElement{}
		SetValue(FieldList, ValueList, p)
		return p
	}
	return flag
}
