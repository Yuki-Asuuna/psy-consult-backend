package snowflake

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"time"
)

var client *snowflake.Snowflake

const (
	dataCenter_id = 0
	workder_id    = 0
)

func SnowflakeInit() error {
	var err error
	client, err = snowflake.NewSnowflake(dataCenter_id, workder_id)
	if err != nil {
		return err
	}
	return nil
}

func GetSnowflakeClient() *snowflake.Snowflake {
	return client
}

func GenID() int64 {
	return client.NextVal()
}

func GetGenTime(ID int64) time.Time {
	timeTemplate := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate, snowflake.GetGenTime(ID), time.Local)
	return stamp
}
