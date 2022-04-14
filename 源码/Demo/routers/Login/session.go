package Login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

//初始化sesson
func InitSession(engine *gin.Engine)  {
	store, err := redis.NewStore(10,"tcp","127.0.0.1:6379", "", []byte("secret"))
	if err != nil {
		log.Println(err.Error())
	}
	//store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession",store))
}

//设置session
func Setsess(context *gin.Context ,key interface{}, value interface{}) error {
	log.SetPrefix("<<<SET SESSION>>>")
	log.Println("保存session:",key,",",value)
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

//获取session
func Getsess(context *gin.Context ,key interface{}) interface{} {
	log.SetPrefix("<<<GET SESSION>>>")
	log.Println("获取session:",key)
	session := sessions.Default(context)
	getSess := session.Get(key)
	log.Println("sess:", getSess)
	return getSess
}