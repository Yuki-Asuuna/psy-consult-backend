package database

import (
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/snowflake"
	"time"
)

func AddConversation(counsellorID string, visitorID string) (int64, error) {
	conversation := Conversation{
		ConversationID: snowflake.GenID(),
		StartTime:      time.Now(),
		CounsellorID:   counsellorID,
		VisitorID:      visitorID,
		Status:         0,
		IsHelped:       0,
	}
	if err := mysql.GetMySQLClient().Create(&conversation).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddConversation Failed, err= %v", err)
		return 0, err
	}
	return conversation.ConversationID, nil
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

func UpdateConversationByConversationID(conversationID int64, updateMap map[string]interface{}) error {
	err := mysql.GetMySQLClient().Model(&Conversation{}).Where("conversation_id = (?)", conversationID).Update(updateMap).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateConversationByConversationID Failed, err= %v, conversationID= %v", err, conversationID)
		return err
	}
	return nil
}
