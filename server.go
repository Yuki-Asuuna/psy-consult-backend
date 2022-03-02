package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/sessions"
)

var r *gin.Engine

func loggerInit() error {
	logFile := "./log/sys.log"
	logrus.SetReportCaller(true)
	file_writer, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("Cannot open sys.log")
		return err
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, file_writer)
	logrus.SetOutput(io.MultiWriter(os.Stdout, file_writer))
	r = gin.Default()
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

	// init session
	if err := sessions.SessionInit(); err != nil {
		logrus.Errorf(constant.Main+"Init Session Failed, err= %v", err)
		return
	}
	logrus.Infof(constant.Main + "Init Session Success!")

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
