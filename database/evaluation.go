package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/utils/mysql"
	"psy-consult-backend/utils/snowflake"
)

func AddEvaluation(rating int64, evaluation string, conversationID int64) (*Evaluation, error) {
	eval := &Evaluation{
		EvaluationID:   snowflake.GenID(),
		Rating:         rating,
		Evaluation:     evaluation,
		ConversationID: conversationID,
	}
	if err := mysql.GetMySQLClient().Create(&eval).Error; err != nil {
		logrus.Errorf(constant.DAO+"AddEvaluation Failed, err= %v", err)
		return nil, err
	}
	return eval, nil
}

func GetEvaluationByConversationID(conversationID int64) (*Evaluation, error) {
	evaluation := new(Evaluation)
	if err := mysql.GetMySQLClient().First(&evaluation, "conversation_id = ?", conversationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logrus.Errorf(constant.DAO+"GetEvaluationByConversationID Failed, err= %v", err)
		return nil, err
	}
	return evaluation, nil
}

func GetEvaluationsByConversationIDs(conversationIDs []int64) (map[int64]*Evaluation, error) {
	ret := make(map[int64]*Evaluation)
	evaluations := make([]*Evaluation, 0)
	err := mysql.GetMySQLClient().Where("conversation_id in (?)", conversationIDs).Find(&evaluations).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetEvaluationsByConversationIDs Failed, err= %v", err)
		return nil, err
	}
	for _, e := range evaluations {
		ret[e.ConversationID] = e
	}
	return ret, nil
}
