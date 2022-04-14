package database

import (
	"github.com/jinzhu/gorm"
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

func CheckBindingExist(counsellorID string, supervisorID string) (bool, error) {
	binding := &Binding{}
	err := mysql.GetMySQLClient().Where("counsellor_id = (?) and supervisor_id = (?)", counsellorID, supervisorID).First(binding).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		logrus.Errorf(constant.DAO+"CheckBindingExist Failed, err= %v", err)
		return false, err
	}
	return true, nil
}

func GetBindingByBindingID(bindingID int64) (*Binding, error) {
	binding := &Binding{}
	err := mysql.GetMySQLClient().Where("binding_id = (?)", bindingID).First(binding).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf(constant.DAO+"GetBindingByBindingID Failed, err= %v", err)
		return nil, err
	}
	return binding, nil
}
