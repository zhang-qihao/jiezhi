package es

import (
	"Demo/conf"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

//func init() {
//	errLog := log.New(os.Stdout, "app", log.LstdFlags)
//	var err error
//	client, err = elastic.NewClient(elastic.SetErrorLog(errLog), elastic.SetURL(host))
//	if err != nil {
//		panic(err)
//	}
//	info, code, err := client.Ping(host).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Es returned with code %d and version %s\n", code, info.Version.Number)
//	versionCode, err := client.ElasticsearchVersion(host)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("es version %s\n", versionCode)
//}

//初始化
func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(host...))
	if err != nil {
		fmt.Printf("create client failed, err: %v", err)
	}
}

//PingNode 连接测试
func PingNode() {
	start := time.Now()

	info, code, err := client.Ping(host[0]).Do(context.Background())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("cost time: %v\n", duration)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func Create(index, typ, id string, data conf.JobInfo) {
	// 1使用结构体存入ES
	//var data conf.JobInfo
	//data.Jid="test12345"
	//data.Title = "测试案例"
	//data.ChangeTime=time.Now()
	//data.CreatTime=time.Now()
	//el := `{"firstname":"John","lastname":"Dan","age":17,"about":"I like rose and jack","interests":["computer","football"]}`
	put, err := client.Index().Index(index).Type(typ).Id(id).BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println("error")
		fmt.Printf("indexed %s to index %s,type %s\n", put.Id, put.Index, put.Type)
	}

}

func Get(index, typ, id string) {
	get, err := client.Get().Index(index).Type(typ).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Printf("get document %s in version %d from index %s,type %s\n", get.Id, get.Version, get.Index, get.Type)
	}
	var job conf.JobInfo
	json.Unmarshal(get.Source, &job)
	fmt.Println("get info, company:", job.Company)
}

func Update(index, typ, id string, a map[string]interface{}) {
	res, err := client.Update().Index(index).Type(typ).Id(id).Doc(a).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("update title %s\n", res.Result)
}

func Delete(index, typ, id string) {
	res, err := client.Delete().Index(index).Type(typ).Id(id).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("update age %s\n", res.Result)
}
func Query(index, typ string) {
	var res *elastic.SearchResult
	//var err error
	res, _ = client.Search(index).Type(typ).Do(context.Background())
	fmt.Println(res.Hits.TotalHits)
}

func Query1(index, typ string) {
	var res *elastic.SearchResult
	//var err error
	q := elastic.NewQueryStringQuery("nature:应届")
	res, _ = client.Search(index).Type(typ).Query(q).Do(context.Background())
	if res.Hits.TotalHits.Value > 0 {
		//for _, hit := range res.Hits.Hits {
		//	var t Employee
		//	err = json.Unmarshal(hit.Source, &t)
		//	if err != nil {
		//		fmt.Println("failed")
		//	}
		//	fmt.Printf("employee name %s:%s\n", t.FirstName, t.LastName)
		//}
		fmt.Printf("found a total of %d Employee\n", res.Hits.TotalHits.Value)
		for id, hit := range res.Hits.Hits {
			fmt.Printf(" 第 %d 个 索引的得分为：%.4f\n", id, *hit.Score)
		}
	} else {
		fmt.Printf("found no data")
	}
}

func NatureMultiSearch(index, typ, natureType, inputText string) [10]conf.JobInfo {
	var res *elastic.SearchResult
	//var err error
	boolq := elastic.NewBoolQuery()
	boolq.Must(elastic.NewMatchQuery("nature", natureType))
	boolq.Must(elastic.NewMultiMatchQuery(inputText).
		FieldWithBoost("require", 0.1).
		FieldWithBoost("describe", 0.1).
		FieldWithBoost("title", 1))
	//boolq.Must(elastic.NewMatchQuery("education", "本科及以上"))

	//boolq.Filter(elastic.NewRangeQuery("nature").Gt(30))
	res, _ = client.Search(index).Type(typ).Query(boolq).Do(context.Background())

	var result [10]conf.JobInfo
	if res.Hits.TotalHits.Value > 0 {
		var job conf.JobInfo
		fmt.Printf("found a total of %d job\n", res.Hits.TotalHits.Value)
		for id, hit := range res.Hits.Hits {
			json.Unmarshal(hit.Source, &job)
			fmt.Printf(" 第 %d 个 索引的得分为：%.4f\n", id, *hit.Score)
			result[id] = job
			if id > 10 {
				break
			}
		}
	} else {
		fmt.Println("found no data")
	}
	return result
}

// Query4 包含查询
func Query4(index, typ string) {
	var res *elastic.SearchResult
	var err error
	matchPhrase := elastic.NewMatchPhraseQuery("title", "开发")
	res, err = client.Search(index).Type(typ).Query(matchPhrase).Do(context.Background())
	if err != nil {
		if res.Hits.TotalHits.Value > 0 {
			fmt.Printf("found a total of %d data\n", res.Hits.TotalHits.Value)
			for id, hit := range res.Hits.Hits {
				fmt.Printf(" 第 %d 个 索引的得分为：%.4f\n", id, *hit.Score)
			}
		} else {
			fmt.Printf("found no data")
		}
	} else {
		panic(err)
	}
}

// Match 多个Field进行搜索
func Match(index, typ, inputText, field string) {
	var res *elastic.SearchResult
	//var err error
	q := elastic.NewMatchQuery(field, inputText)
	res, _ = client.Search().Index(index).Type(typ).Query(q).Do(context.Background())
	fmt.Println("Total number :=", res.Hits.TotalHits.Value)
	if res.Hits.TotalHits.Value > 0 {
		var job conf.JobInfo
		fmt.Printf("found a total of %d job\n", res.Hits.TotalHits.Value)
		for id, hit := range res.Hits.Hits {
			json.Unmarshal(hit.Source, &job)
			fmt.Printf(" 第 %d 个 索引, 标题为：%s, 得分为：%.4f\n", id, job.Title, *hit.Score)
		}
	} else {
		fmt.Printf("found no data")
	}

}

// MultiMatch 多个Field进行搜索
func MultiMatch(index, typ, inputText string) [10]conf.JobInfo {
	var res *elastic.SearchResult
	//var err error
	matchQuery := elastic.NewMultiMatchQuery(inputText).
		FieldWithBoost("require", 0.1).
		FieldWithBoost("describe", 0.1).
		FieldWithBoost("title", 1)
	res, _ = client.Search().Index(index).Type(typ).Query(matchQuery).Do(context.Background())
	var result [10]conf.JobInfo
	if res.Hits.TotalHits.Value > 0 {
		var job conf.JobInfo
		fmt.Printf("found a total of %d job\n", res.Hits.TotalHits.Value)
		for id, hit := range res.Hits.Hits {
			json.Unmarshal(hit.Source, &job)
			fmt.Printf(" 第 %d 个 索引的得分为：%.4f\n", id, *hit.Score)
			result[id] = job
			if id > 10 {
				break
			}
		}
	} else {
		fmt.Printf("found no data")
	}
	return result
}
