package database

type Evaluation struct {
	EvaluationID   int64
	Rating         int64
	Evaluation     string
	ConversationID int64
}

func (Evaluation) TableName() string {
	return "evaluation"
}
