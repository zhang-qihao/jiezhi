package company

import (
	"time"
)

type JobInfo struct {
	Jid         string    `json:"jid"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Cid         string    `json:"cid"`
	Salary      string    `json:"salary"`
	Nature      string    `json:"nature"`
	Education   string    `json:"education"`
	Experience  string    `json:"experience"`
	Address     string    `json:"address"`
	FullAddress string    `json:"fullAddress"`
	Require     string    `json:"require"`
	Describe    string    `json:"describe"`
	CreatTime   time.Time `json:"createTime"`
	ChangeTime  time.Time `json:"changeTime"`
	pid         string    `json:"pid"`
}

type Rank struct {
	//用到json记得一定大写首字母
	ShortName   string `json:"shortname"`
	Title       string `json:"title"`
	Salary      string `json:"salary"`
	Scale       string `json:"scale"`
	Address     string `json:"address"`
	CompanyType string `json:"companytype"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
}

//func Laceinvite(DB *gorm.DB) gin.HandlerFunc {
//	return func(context *gin.Context) {
//		cid := context.PostForm("cid")
//		fmt.Println("cid = ", cid)
//		var title string
//		rank0 := make([]Rank, 0)
//		rank1 := make([]Rank, 0)
//		rank2 := make([]Rank, 0)
//		rank3 := make([]Rank, 0)
//		rank4 := make([]Rank, 0)
//		jopinfo := make([]JobInfo, 0)
//		var jobinfo []JobInfo
//		var personpost []conf.PersonPost
//		var ShortName string
//		var CompanyType string
//		var Address string
//		var Scale string
//		//找出这个公司的全部岗位以及相关信息
//		DB.Where("cid = ?", cid).Find(&jobinfo)
//		DB.Where("cid = ?", cid).Find(&personpost)
//		//找出这个公司的全部岗位的相关信息
//		for _, value := range jobinfo {
//			var jop = JobInfo{
//				Jid:        value.Jid,
//				Title:      value.Title,
//				Company:    value.Company,
//				Cid:        value.Cid,
//				Salary:     value.Salary,
//				Education:  value.Education,
//				Experience: value.Experience,
//				Address:    value.Address,
//				Require:    value.Require,
//				Describe:   value.Describe,
//				CreatTime:  value.CreatTime,
//				ChangeTime: value.ChangeTime,
//				Nature:     value.Nature,
//			}
//			jopinfo = append(jopinfo, jop)
//
//		}
//		var company conf.CompanyInfo
//		DB.Where("cid = ?", cid).Find(&company)
//		ShortName = company.ShortName
//		CompanyType = company.CompanyType
//		Address = company.Address
//		Scale = company.Scale
//		///找出这个公司全部的岗位以及在前端需要显示的内容
//		for _, value := range jopinfo {
//			var ran1 = Rank{
//				Title:       value.Title,
//				Salary:      value.Salary,
//				Education:   value.Education,
//				Experience:  value.Experience,
//				ShortName:   ShortName,
//				Address:     Address,
//				CompanyType: CompanyType,
//				Scale:       Scale,
//			}
//			rank0 = append(rank0, ran1)
//		}
//		fmt.Println("企业id为：" + cid + "的全部岗位相关信息如下")
//		fmt.Println(rank0)
//		//laceinvite.html中 全部 按钮的功能实现
//		for _, value := range personpost {
//			for _, value1 := range jopinfo {
//				if value.Jid == value1.Jid {
//					title = value1.Title
//				}
//			}
//			for _, value2 := range rank0 {
//				if value2.Title == title {
//					rank1 = append(rank1, value2)
//				}
//			}
//		}
//		fmt.Println("企业id为：" + cid + "的全部被投递岗位相关信息如下")
//		fmt.Println(rank1)
//		//laceinvite.html中 待处理 按钮的功能实现
//		for _, value := range personpost {
//			if value.Status == "0" {
//				for _, value1 := range jopinfo {
//					if value.Jid == value1.Jid {
//						title = value1.Title
//					}
//				}
//				for _, value2 := range rank0 {
//					if value2.Title == title {
//						rank2 = append(rank2, value2)
//					}
//				}
//			}
//		}
//		fmt.Println("企业id为：" + cid + "的待处理岗位相关信息如下")
//		fmt.Println(rank2)
//		//laceinvite.html中 已接受 按钮的功能实现
//		for _, value := range personpost {
//			if value.Status == "2" {
//				for _, value1 := range jopinfo {
//					if value.Jid == value1.Jid {
//						title = value1.Title
//					}
//				}
//				for _, value2 := range rank0 {
//					if value2.Title == title {
//						rank3 = append(rank3, value2)
//					}
//				}
//			}
//		}
//		fmt.Println("企业id为：" + cid + "的已接收岗位相关信息如下")
//		fmt.Println(rank3)
//		//laceinvite.html中 已拒绝 按钮的功能实现
//		for _, value := range personpost {
//			if value.Status == "3" {
//				for _, value1 := range jopinfo {
//					if value.Jid == value1.Jid {
//						title = value1.Title
//					}
//				}
//				for _, value2 := range rank0 {
//					if value2.Title == title {
//						rank4 = append(rank4, value2)
//					}
//				}
//			}
//		}
//		fmt.Println("企业id为：" + cid + "的已拒绝岗位相关信息如下")
//		fmt.Println(rank4)
//
//		context.JSON(http.StatusOK, gin.H{
//			"rank1": rank1,
//			"rank2": rank2,
//			"rank3": rank3,
//			"rank4": rank4,
//		})
//	}
//}
