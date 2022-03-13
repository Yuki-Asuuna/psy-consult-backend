package im_message

type TextMessageContent struct {
	Text string `json:"Text"`
}

type TextMessage struct {
	MsgType    string             `json:"MsgType"`
	MsgContent TextMessageContent `json:"MsgContent"`
}

type SendTextMessageRequest struct {
	SyncOtherMachine int           `json:"SyncOtherMachine"`
	FromAccount      string        `json:"From_Account"`
	ToAccount        string        `json:"To_Account"`
	MsgRandom        int           `json:"MsgRandom"`
	MsgBody          []TextMessage `json:"MsgBody"`
	// CloudCustomData  string        `json:"CloudCustomData"`
	// MsgSeq           int           `json:"MsgSeq"`
	// MsgTimeStamp     int           `json:"MsgTimeStamp"`

}

type SendTextMessageResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
	MsgTime      int    `json:"MsgTime"`
	MsgKey       string `json:"MsgKey"`
}
