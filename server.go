package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"psy-consult-backend/constant"
	"psy-consult-backend/middleware"
	"psy-consult-backend/utils/ip"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/redis"
	"psy-consult-backend/utils/sessions"
	"psy-consult-backend/utils/snowflake"
)

var r *gin.Engine

const (
	es_server = "http://124.221.197.218:9200"
)

func loggerInit() error {
	ip, err := ip.GetOutBoundIP()
	if err != nil {
		logrus.Error("Get LocalIP failed, err= %v", err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(es_server))
	if err != nil {
		panic(err)
		return err
	}
	logrus.Info("Init elastic client sucess")

	hook, err := elogrus.NewElasticHook(client, ip, logrus.TraceLevel, "psy-backend-log")
	if err != nil {
		panic(err)
		return err
	}
	logrus.AddHook(hook)
	logrus.Info("Init elastic hook sucess")

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 本地日志废除
	//logFile := "./log/sys.log"
	//file_writer, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	//if err != nil {
	//	panic("Cannot open sys.log")
	//	return err
	//}
	//gin.DefaultWriter = io.MultiWriter(os.Stdout, file_writer)
	//logrus.SetOutput(io.MultiWriter(os.Stdout, file_writer))
	r = gin.Default()
	r.Use(middleware.Log4Gin())
	return nil
}

func main() {
	// logger init
	err := loggerInit()
	if err != nil {
		return
	}

	// handler init
	httpHandlerInit()

	// init redis
	redis.RedisInit()

	// init session
	if err := sessions.SessionInit(); err != nil {
		logrus.Errorf(constant.Main+"Init Session Failed, err= %v", err)
		return
	}
	logrus.Infof(constant.Main + "Init Session Success!")

	// init unique id generator (twitter/snowflake)
	if err := snowflake.SnowflakeInit(); err != nil {
		logrus.Errorf(constant.Main+"Init Snowflake Failed, err= %v", err)
		return
	}

	// init mysql database
	if err := mysql.MysqlInit(); err != nil {
		logrus.Error(constant.Main+"Init Mysql Failed, err= %v", err)
		return
	}
	logrus.Infof(constant.Main + "Init Mysql Success!")

	// start gin
	if err := r.Run(":8000"); err != nil {
		logrus.Error(constant.Main+"Run Gin Server Failed, err= %v", err)
	}
}
