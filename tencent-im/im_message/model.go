package im_message

type TextMessageContent struct {
	Text string `json:"Text"`
}

type TextMessage struct {
	MsgType    string             `json:"MsgType"`
	MsgContent TextMessageContent `json:"MsgContent"`
}

type SendTextMessageRequest struct {
	SyncOtherMachine int64         `json:"SyncOtherMachine"`
	FromAccount      string        `json:"From_Account"`
	ToAccount        string        `json:"To_Account"`
	MsgRandom        int64         `json:"MsgRandom"`
	MsgBody          []TextMessage `json:"MsgBody"`
	// CloudCustomData  string        `json:"CloudCustomData"`
	// MsgSeq           int           `json:"MsgSeq"`
	// MsgTimeStamp     int           `json:"MsgTimeStamp"`

}

type SendTextMessageResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int64  `json:"ErrorCode"`
	MsgTime      int64  `json:"MsgTime"`
	MsgKey       string `json:"MsgKey"`
}

type SearchHistoryMessageRequest struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MinTime     int64  `json:"MinTime"`
	MaxTime     int64  `json:"MaxTime"`
	MaxCnt      int64  `json:"MaxCnt"`
	LastMsgKey  string `json:"LastMsgKey"`
	LastMsgTime int64  `json:"LastMsgTime"`
}

type MessageInfo struct {
	FromAccount  string `json:"From_Account"`
	ToAccount    string `json:"To_Account"`
	MsgSeq       int64  `json:"MsgSeq"`
	MsgRandom    int64  `json:"MsgRandom"`
	MsgTimeStamp int64  `json:"MsgTimeStamp"`
	MsgFlagBits  int64  `json:"MsgFlagBits"`
	MsgKey       string `json:"MsgKey"`
	MsgBody      []struct {
		MsgType    string `json:"MsgType"`
		MsgContent struct {
			Text string `json:"Text"`
		} `json:"MsgContent"`
	} `json:"MsgBody"`
	CloudCustomData string `json:"CloudCustomData"`
}

type MessageInfoSlice []MessageInfo

func (m MessageInfoSlice) Len() int           { return len(m) }
func (m MessageInfoSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MessageInfoSlice) Less(i, j int) bool { return m[i].MsgTimeStamp < m[j].MsgTimeStamp }

type SearchHistoryMessageResponse struct {
	ActionStatus string        `json:"ActionStatus"`
	ErrorInfo    string        `json:"ErrorInfo"`
	ErrorCode    int64         `json:"ErrorCode"`
	Complete     int64         `json:"Complete"`
	MsgCnt       int64         `json:"MsgCnt"`
	LastMsgTime  int64         `json:"LastMsgTime"`
	LastMsgKey   string        `json:"LastMsgKey"`
	MsgList      []MessageInfo `json:"MsgList"`
}
