package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"time"
)

func AddConversation(counsellorID string, visitorID string, conversationID int64, groupID string) error {
	conversation := Conversation{
		ConversationID: conversationID,
		EndTime:        time.Unix(0, 0),
		StartTime:      time.Now(),
		CounsellorID:   counsellorID,
		VisitorID:      visitorID,
		Status:         0,
		IsHelped:       0,
		GroupID:        groupID,
	}
	if err := mysql.GetMySQLClient().Create(&conversation).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddConversation Failed, err= %v", err)
		return err
	}
	return nil
}

func GetConversationByConversationID(conversationID int64) (*Conversation, error) {
	conversation := &Conversation{}
	err := mysql.GetMySQLClient().Where("conversation_id = (?)", conversationID).Find(conversation).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetConversationByConversationID Failed, err= %v", err)
		return nil, err
	}
	return conversation, nil
}

func GetConversationByGroupID(groupID string) (*Conversation, error) {
	conversation := &Conversation{}
	err := mysql.GetMySQLClient().Where("group_id = (?)", groupID).Find(conversation).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetConversationByGroupID Failed, err= %v", err)
		return nil, err
	}
	return conversation, nil
}

func UpdateConversationByConversationID(conversationID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&Conversation{}).Where("conversation_id = (?)", conversationID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateConversationByConversationID Failed, err= %v, conversationID= %v", err, conversationID)
		return err
	}
	return nil
}

func GetConversationList(page int64, size int64, startTime int64, endTime int64) ([]*Conversation, error) {
	conversations := make([]*Conversation, 0)
	query := mysql.GetMySQLClient()
	// role == 0 表示选择全部
	//if role != 0 {
	//	query = query.Where("role = (?)", role)
	//}
	query = query.Where("start_time >= (?)", time.Unix(startTime, 0))
	query = query.Where("end_time <= (?)", time.Unix(endTime, 0))
	err := query.Offset(page * size).Limit(size).Find(&conversations).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetConversationList Failed, err= %v", err)
		return nil, err
	}
	return conversations, nil
}

func GetConversationListByCounsellorIDAndTimeInterval(counsellorID string, startTime int64, endTime int64) ([]*Conversation, error) {
	conversations := make([]*Conversation, 0)
	query := mysql.GetMySQLClient()
	query = query.Where("counsellor_id = (?)", counsellorID)
	query = query.Where("start_time >= (?)", time.Unix(startTime, 0))
	query = query.Where("end_time <= (?)", time.Unix(endTime, 0))
	err := query.Find(&conversations).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetConversationListByCounsellorIDAndTimeInterval Failed, err= %v", err)
		return nil, err
	}
	return conversations, nil
}

func GetConversationBySender(fromAccount string, toAccount string) (*Conversation, error) {
	conversation := &Conversation{}
	query := mysql.GetMySQLClient()
	// 正在进行
	query = query.Where("status = (0)")
	query = query.Where("counsellor_id = (?)", fromAccount)
	query = query.Where("visitor_id = (?)", toAccount)
	err := query.First(conversation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetConversationBySender Failed, err= %v", err)
		return nil, err
	}
	return conversation, nil
}
