package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/mysql"
	"time"
)

func GetCounsellorUserByCounsellorID(counsellorID string) (*CounsellorUser, error) {
	counsellor := new(CounsellorUser)
	if err := mysql.GetMySQLClient().First(&counsellor, "counsellor_id = ?", counsellorID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetCounsellorUserByCounsellorID Failed, err= %v", err)
		return nil, err
	}
	return counsellor, nil
}

func AddCounsellorUser(username string, password string, role int) error {
	counsellor := CounsellorUser{
		Username:     username,
		Password:     password,
		CounsellorID: helper.S2MD5(username),
		Role:         role,
		LastLogin:    time.Unix(0, 0),
	}
	if err := mysql.GetMySQLClient().Create(&counsellor).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddCounsellorUser Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateCounsellorUserByCounsellorID(counsellorID string, name string, password string, status int, gender int, age int, identityNumber string, phoneNumber string, avatar string, email string, title string, department string, qualification string, introduction string, maxConsults int) error {
	err := mysql.GetMySQLClient().Model(&CounsellorUser{}).Where("counsellor_id = (?)", counsellorID).Update(map[string]interface{}{
		"name":            name,
		"password":        password,
		"status":          status,
		"gender":          gender,
		"age":             age,
		"identity_number": identityNumber,
		"phone_number":    phoneNumber,
		"avatar":          avatar,
		"email":           email,
		"title":           title,
		"department":      department,
		"qualification":   qualification,
		"introduction":    introduction,
		"max_consults":    maxConsults,
	}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateCounsellorUserByCounsellorID Failed, err= %v, counsellorID= %v", err, counsellorID)
		return err
	}
	return nil
}
