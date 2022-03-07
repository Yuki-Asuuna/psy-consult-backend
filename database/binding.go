package database

import (
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/snowflake"
)

func AddBinding(counsellorID string, supervisorID string) error {
	binding := Binding{
		BindingID:    snowflake.GenID(),
		SupervisorID: supervisorID,
		CounsellorID: counsellorID,
	}
	if err := mysql.GetMySQLClient().Create(&binding).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddBinding Failed, err= %v", err)
		return err
	}
	return nil
}

func DeleteBinding(arrangeID int64) error {
	err := mysql.GetMySQLClient().Where("binding_id = (?)", arrangeID).Delete(&Binding{}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"DeleteBinding Failed, err= %v", err)
		return err
	}
	return nil
}

func GetBindingByCounsellorID(counsellorID string) ([]*Binding, error) {
	binding := make([]*Binding, 0)
	err := mysql.GetMySQLClient().Where("counsellor_id = (?)", counsellorID).Find(&binding).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetBindingByCounsellorID Failed, err= %v", err)
		return nil, err
	}
	return binding, nil
}
