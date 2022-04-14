package es

import (
	"Demo/conf"
	"Demo/connet"
	"Demo/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = []string{
	"http://127.0.0.1:9200/",
}
var (
	DB   *gorm.DB
	ESID = 1182
)

var Jobmapping = `{
	"settings":{
		"number_of_shards": 3,
		"number_of_replicas": 1
	},
	"mappings":{
		"doc":{
			"properties":{
				"jid":{
					"type":"keyword"
				},
				"title":{
					"type": "text",
					"index": "true"
				},
				"cid":{
					"type":"keyword"
				},
				"company":{
					"type": "text",
					"index": "true"
				},
				"nature":{
					"type":"keyword"
				},
				"salary":{
					"type":"keyword"
				},
				"education":{
					"type":"keyword"
				},
				"experience":{
					"type":"keyword"
				},
				"province":{
					"type":"keyword"
				},
				"address":{
					"type": "text",
					"index": "true"
				},
				"require":{
					"type": "text",
					"index": "true"
				},
				"describe":{
					"type": "text",
					"index": "true"
				},
				"createTime":{
					"type":"date"
				},
				"changeTime":{
					"type":"date"
				}
			}
		}
	}
}`

func GetDataFromSql(DB *gorm.DB) []conf.JobInfo {
	configPath := "../conf/config.ini"
	if err := setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库,在全局变量DB中
	DB, err := connet.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}

	var JobList []conf.JobInfo

	DB.Find(&JobList)
	return JobList
}
