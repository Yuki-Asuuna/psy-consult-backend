package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
)

func AddVisitorUser(visitorID string) error {
	visitor := VisitorUser{
		VisitorID: visitorID,
	}
	if err := mysql.GetMySQLClient().Create(&visitor).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddVisitorUser Failed, err= %v", err)
		return err
	}
	return nil
}

func GetVisitorUserByVisitorID(visitorID string) (*VisitorUser, error) {
	visitor := new(VisitorUser)
	if err := mysql.GetMySQLClient().First(&visitor, "visitor_id = ?", visitorID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetVisitorUserByVisitorID Failed, err= %v", err)
		return nil, err
	}
	return visitor, nil
}

func GetVisitorUserList(page int, size int) ([]*VisitorUser, error) {
	visitors := make([]*VisitorUser, 0)
	query := mysql.GetMySQLClient()
	err := query.Offset(page * size).Limit(size).Find(&visitors).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetVisitorUserList Failed, err= %v", err)
		return nil, err
	}
	return visitors, nil
}

func UpdateVisitorUserByVisitorID(visitorID string, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&VisitorUser{}).Where("visitor_id = (?)", visitorID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateVisitorUserByVisitorID Failed, err= %v, applicationID= %v", err, visitorID)
		return err
	}
	return nil
}
