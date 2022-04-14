package MainPage

import (
	"Demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

type CompanyInfo struct {
	Cid         string    `json:"cid"`
	ShortName   string    `json:"shortname"`
	Fullname    string    `json:"fullName"`
	CompanyType string    `json:"companyType"`
	Address     string    `json:"address"`
	Industry    string    `json:"industry"`
	Scale       string    `json:"scale"`
	Logourl     string    `json:"logoUrl"`
	Brief       string    `json:"brief"`
	Creattime   time.Time `json:"createTime"`
	Changetime  time.Time `json:"changeTime"`
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

//定义结构体用于存放职位信息
type Rank struct {
	Jid         string `json:"jid"`
	Title       string `json:"title"`
	Shortname   string `json:"shortname"`
	Salary      string `json:"salary"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
	Address     string `json:"address"`
	Companytype string `json:"companytype"`
	Scale       string `json:"scale"`
}

// InfoTransfer 向前端传值
func InfoTransfer(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取成功登陆的用户Pid
		ID := c.PostForm("ID")
		//根据Pid获取用户basicInfo
		arr := getUserBasicInfo(ID, DB)
		if arr != nil {
			//获取要放在主页的职位信息
			rank1, rank2, rank3 := GetJobInfoFromDatabase(DB)
			//向前端传json数据：用户姓名、12+9+3=24组职位信息，用于主页面显示
			c.JSON(http.StatusOK, gin.H{"name": arr[1], "rank1": rank1, "rank2": rank2, "rank3": rank3})
			log.SetPrefix("<<<Login>>>")
			log.Printf("%s sign in, pid = %s\n", arr[1], ID)

		}
	}
}

// getUserBasicInfo 查找用户信息
func getUserBasicInfo(ID string, DB *gorm.DB) []string {
	//new结构体
	var info = conf.PersonInfoStr{}
	//根据Pid查找用户
	DB.Where("Pid = ?", ID).Find(&info)
	//按"#"分割数据
	sep := "#"
	//basicInfo存入arr
	arr := strings.Split(info.BasicInfo, sep)
	return arr
}

// GetJobInfoFromDatabase 从数据库中获取职位信息
func GetJobInfoFromDatabase(DB *gorm.DB) ([]Rank, []Rank, []Rank) {
	//new结构体
	rank1 := make([]Rank, 0)
	rank2 := make([]Rank, 0)
	rank3 := make([]Rank, 0)
	var job []JobInfo
	DB.Find(&job)
	var company []CompanyInfo
	DB.Find(&company)

	var Jid1 string
	var Cid string
	var Title1 string
	var Salary1 string
	var Education1 string
	var Experience1 string
	var Address1 string
	var Shortname1 string
	var CompanyType1 string
	var Scale1 string
	//定义count用来计数
	var count = 0

	//遍历数据库
	for _, value2 := range job {
		Jid1 = value2.Jid
		Title1 = value2.Title
		Salary1 = value2.Salary
		Education1 = value2.Education
		Experience1 = value2.Experience
		Address1 = value2.Address
		Cid = value2.Cid
		for _, value3 := range company {
			if value3.Cid == Cid {
				Shortname1 = value3.ShortName
				CompanyType1 = value3.CompanyType
				Scale1 = value3.Scale
			}
		}
		//new结构体
		var rank = Rank{
			Jid:         Jid1,
			Title:       Title1,
			Salary:      Salary1,
			Education:   Education1,
			Experience:  Experience1,
			Shortname:   Shortname1,
			Address:     Address1,
			Companytype: CompanyType1,
			Scale:       Scale1,
		}
		count++
		if count <= 12 {
			rank1 = append(rank1, rank)
		} else if count <= 21 {
			rank2 = append(rank2, rank)
		} else if count <= 24 {
			rank3 = append(rank3, rank)
		} else {
			//超过25就跳出循环
			break
		}
	}
	return rank1, rank2, rank3

}
