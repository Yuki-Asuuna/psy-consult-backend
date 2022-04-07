package api

type EvaluationInfoResponse struct {
	EvaluationID   string `json:"evaluationID"`
	ConversationID string `json:"conversationID"`
	Rating         int64  `json:"rating"`
	Evaluation     string `json:"evaluation"`
}
