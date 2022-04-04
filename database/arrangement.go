package database

import (
	"github.com/jinzhu/gorm"
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

func CheckUnique(counsellorID string, arrangeTime time.Time) bool {
	ret := Arrangement{}
	err := mysql.GetMySQLClient().Where("arrange_time = (?) and counsellor_id = (?)", arrangeTime, counsellorID).First(&ret).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}
		logrus.Errorf(constant.DAO+"CheckUnique Failed, err= %v", err)
	}
	return false
}

func GetArrangementsByArrangeTimeList(arrangeTimeList []time.Time) (map[int64][]*Arrangement, error) {
	arr := make([]*Arrangement, 0)
	err := mysql.GetMySQLClient().Where("arrange_time in (?)", arrangeTimeList).Find(&arr).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetArrangementsByArrangeTimeList Failed, err= %v", err)
		return nil, nil
	}
	ret := make(map[int64][]*Arrangement)
	for _, a := range arr {
		unixTime := a.ArrangeTime.Unix()
		_, ok := ret[unixTime]
		if !ok {
			list := make([]*Arrangement, 0)
			list = append(list, a)
			ret[unixTime] = list
		} else {
			ret[unixTime] = append(ret[unixTime], a)
		}
	}
	return ret, nil
}
