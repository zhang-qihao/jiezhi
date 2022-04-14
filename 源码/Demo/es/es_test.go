package es

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPingNode(t *testing.T) {
	PingNode()
}

func TestGetDataFromSql(t *testing.T) {
	jobList := GetDataFromSql(DB)
	var id string
	index := "job"
	typ := "job_info"
	for i := 0; i < len(jobList); i++ {
		id = strconv.Itoa(i + 1)
		fmt.Printf("第%s个例子的标题为:%s\n", id, jobList[i].Title)
		Create(index, typ, id, jobList[i])
	}
}

//func TestGet(t *testing.T) {
//	index := "job"
//	typ := "job_info"
//	id := "100"
//	Get(index, typ, id)
//}
//func TestUpdate(t *testing.T) {
//	index := "job"
//	typ := "job_info"
//	id := "1"
//	var a = map[string]interface{}{"title": "这是个测试案例"}
//	Update(index, typ, id, a)
//}
//func TestQuery(t *testing.T) {
//	index := "job"
//	typ := "job_info"
//	Query(index, typ)
//}
//
//
//

//func TestMatch(t *testing.T) {
//	index := "job"
//	typ := "job_info"
//	inputText := "软件开发"
//	field := "title"
//	Match(index, typ, inputText, field)
//}

//func TestMultiMatch(t *testing.T) {
//	index := "job"
//	typ := "job_info"
//	inputText := "python软件开发京东优选"
//	MultiMatch(index, typ, inputText)
//}

func TestNatureMultiSearch(t *testing.T) {
	index := "job"
	typ := "job_info"
	inputText := "python软件开发京东优选"
	natureType := "实习"
	NatureMultiSearch(index, typ, natureType, inputText)
}

//func TestMod(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		fmt.Printf("    %d,mod= %d\n", i, i%4)
//	}
//}
