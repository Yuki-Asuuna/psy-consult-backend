package sessions

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/boj/redistore.v1"
	"psy-consult-backend/database"
	"psy-consult-backend/utils/helper"
)

var client *redistore.RediStore

const (
	redis_address   = "127.0.0.1:6379"
	redis_password  = ""
	redis_network   = "tcp"
	redis_size      = 10
	redis_secretkey = "secret key"
	redis_maxage    = 7 * 24 * 3600
)

func SessionInit() error {
	var err error
	client, err = redistore.NewRediStore(redis_size, redis_network, redis_address, redis_password, []byte(redis_secretkey))
	if err != nil {
		return err
	}
	client.SetMaxAge(redis_maxage)
	return nil
}

func GetSessionClient() *redistore.RediStore {
	return client
}

func GetCounsellorNameBySession(c *gin.Context) string {
	session, _ := client.Get(c.Request, "dotcomUser")
	ret, ok := session.Values["username"]
	if !ok {
		return ""
	}
	return ret.(string)
}

func GetCounsellorInfoBySession(c *gin.Context) *database.CounsellorUser {
	currentCounsellorName := GetCounsellorNameBySession(c)
	currentCounsellorID := helper.S2MD5(currentCounsellorName)
	counsellor, err := database.GetCounsellorUserByCounsellorID(currentCounsellorID)
	if err != nil {
		return nil
	}
	return counsellor
}
