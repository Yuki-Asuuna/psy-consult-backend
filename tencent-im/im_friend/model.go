package im_friend

type AddFriendItem struct {
	ToAccount string `json:"To_Account"`
	AddSource string `json:"AddSource"`
	GroupName string `json:"GroupName"`
}

type AddFriendRequest struct {
	FromAccount   string          `json:"From_Account"`
	AddFriendItem []AddFriendItem `json:"AddFriendItem"`
}

type AddFriendResponse struct {
	ResultItem []struct {
		ToAccount  string `json:"To_Account"`
		ResultCode int    `json:"ResultCode"`
		ResultInfo string `json:"ResultInfo"`
	} `json:"ResultItem"`
	FailAccount  []string `json:"Fail_Account"`
	ActionStatus string   `json:"ActionStatus"`
	ErrorCode    int      `json:"ErrorCode"`
	ErrorInfo    string   `json:"ErrorInfo"`
	ErrorDisplay string   `json:"ErrorDisplay"`
}

type DeleteFriendRequest struct {
	FromAccount string   `json:"From_Account"`
	ToAccount   []string `json:"To_Account"`
	DeleteType  string   `json:"DeleteType"`
}

type DeleteFriendResponse struct {
	ResultItem []struct {
		ToAccount  string `json:"To_Account"`
		ResultCode int    `json:"ResultCode"`
		ResultInfo string `json:"ResultInfo"`
	} `json:"ResultItem"`
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorDisplay string `json:"ErrorDisplay"`
}
