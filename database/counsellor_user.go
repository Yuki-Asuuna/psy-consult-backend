package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
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
