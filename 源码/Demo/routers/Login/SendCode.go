package Login

import (
	"Demo/setting"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-gonic/gin.v1/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// SendCode 生成6位随机数字验证码
// mode 等于1，则开启发送服务，
// mode 等于0，则关闭发送服务
func SendCode(c *gin.Context, R redis.Conn, mode int) string {
	phone, _, _ := GetDataFromHTML(c) //获取前端输入的手机号
	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	//2.调用阿里云sdk，完成发送
	if mode == 1 {
		config := setting.Conf //对应到app.json里面的相关配置
		client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
		if err != nil {
			log.Fatal(err.Error())
		}
		//构造一个请求的结构体对象对象
		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https" //请求类型为https，更加安全
		request.SignName = config.SignName
		request.TemplateCode = config.TemplateCode
		request.PhoneNumbers = phone
		par, err := json.Marshal(map[string]interface{}{
			"code": code,
		})
		request.TemplateParam = string(par)
		response, err := client.SendSms(request) //向阿里云传入数据
		fmt.Println(response)                    //打印传送的结果
		if err != nil {
			log.Fatal(err.Error())
		}
		//验证是否发送成功
		if response.Code == "OK" {
			fmt.Println("验证码发送成功为： ", code)
			//手机号验证码存入redis并设置过期时间
			WriteRedis(R, phone, code, 3, 1)   //手机号、code写入db：3
			_, err = R.Do("expire", phone, 60) //设置过期时间
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{"msg": "生成验证码"}) //给前端返回JSON格式数据
		}
		return code
	} else {
		var err error
		WriteRedis(R, phone, code, 3, 1)   //手机号、code写入db：3
		_, err = R.Do("expire", phone, 60) //设置过期时间
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"msg": "生成验证码"}) //给前端返回JSON格式数据
		return code
	}

}

func CreateID(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	ID := hex.EncodeToString(h.Sum(nil))
	ID = ID[:16]
	return ID
}
