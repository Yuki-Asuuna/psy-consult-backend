package database

import (
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/snowflake"
	"time"
)

func AddArrangement(counsellorID string, role int, name string, arrangeTime time.Time) error {
	arrangement := Arrangement{
		ArrangeID:    snowflake.GenID(),
		CounsellorID: counsellorID,
		ArrangeTime:  arrangeTime,
		Role:         role,
	}
	if err := mysql.GetMySQLClient().Create(&arrangement).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddArrangement Failed, err= %v", err)
		return err
	}
	return nil
}

func GetArrangementsByArrangeTime(arrangeTime time.Time) ([]*Arrangement, error) {
	ret := make([]*Arrangement, 0)
	err := mysql.GetMySQLClient().Where("arrange_time = (?)", arrangeTime).Find(&ret).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetArrangementsByArrangeTime Failed, err= %v", err)
		return nil, err
	}
	return ret, nil
}

func DeleteArrangement(arrangeID int64) error {
	err := mysql.GetMySQLClient().Where("arrange_id = (?)", arrangeID).Delete(&Arrangement{}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"DeleteArrangement Failed, err= %v", err)
		return err
	}
	return nil
}
