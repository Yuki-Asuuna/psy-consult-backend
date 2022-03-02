package sessions

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/boj/redistore.v1"
)

var client *redistore.RediStore

const (
	redis_address  = "127.0.0.1:6379"
	redis_password = ""
	redis_network  = "tcp"
	redis_size     = 10
)

func SessionInit() error {
	var err error
	client, err = redistore.NewRediStore(redis_size, redis_network, redis_address, redis_password, []byte("secret key"))
	if err != nil {
		return err
	}
	return nil
}

func GetSessionClient() *redistore.RediStore {
	return client
}

func GetUserNameBySession(c *gin.Context) string {
	session, _ := client.Get(c.Request, "dotcomUser")
	ret, ok := session.Values["username"]
	if !ok {
		return ""
	}
	return ret.(string)
}
