package account_manage

type AccountImportRequest struct {
	UserID  string `json:"UserID"`
	Nick    string `json:"Nick"`
	FaceUrl string `json:"FaceUrl"`
}

type AccountImportResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

type DeleteItem struct {
	UserID string `json:"UserID"`
}

type AccountDeleteRequest struct {
	DeleteItem []DeleteItem `json:"DeleteItem"`
}

type DeleteResultItem struct {
	ResultCode int    `json:"ResultCode"`
	ResultInfo string `json:"ResultInfo"`
	UserID     string `json:"UserID"`
}

type AccountDeleteResponse struct {
	ActionStatus string             `json:"ActionStatus"`
	ErrorCode    int                `json:"ErrorCode"`
	ErrorInfo    string             `json:"ErrorInfo"`
	ResultItem   []DeleteResultItem `json:"ResultItem"`
}

type ProfileItem struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type PortraitSetRequest struct {
	FromAccount string        `json:"From_Account"`
	ProfileItem []ProfileItem `json:"ProfileItem"`
}

type PortraitSetResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorDisplay string `json:"ErrorDisplay"`
}
