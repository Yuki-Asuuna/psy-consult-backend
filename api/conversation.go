package api

import "time"

type ConversationSearchResponse struct {
	ConversationID int64                   `json:"conversationID"`
	StartTime      time.Time               `json:"startTime"`
	EndTime        time.Time               `json:"endTime"`
	Counsellor     *CounsellorInfoResponse `json:"counsellorInfo"`
	Visitor        *VisitorInfoResponse    `json:"visitorInfo"`
	Status         int64                   `json:"status"`
	IsHelped       int64                   `json:"isHelped"`
	Supervisor     *CounsellorInfoResponse `json:"supervisorInfo"`
	Evaluation     *EvaluationInfoResponse `json:"evaluation"`
	GroupID        string                  `json:"groupID"`
}

type TodayStatResponse struct {
	TotalCount int `json:"totalCount"`
	Hour       int `json:"hour"`
	Minute     int `json:"minute"`
	Second     int `json:"second"`
}

type TodayStatAllResponse struct {
	TotalCount int `json:"totalCount"`
	Hour       int `json:"hour"`
	Minute     int `json:"minute"`
	Second     int `json:"second"`
}

type NStatResponse struct {
	DateList  []string `json:"dateList"`
	CountList []int    `json:"countList"`
}

type AddConversationResponse struct {
	ConversationID int64  `json:"conversationID"`
	GroupID        string `json:"groupID"`
}
