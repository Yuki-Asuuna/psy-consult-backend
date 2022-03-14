package api

type EvaluationInfoResponse struct {
	EvaluationID   int64  `json:"evaluationID"`
	ConversationID int64  `json:"conversationID"`
	Rating         int64  `json:"rating"`
	Evaluation     string `json:"evaluation"`
}
