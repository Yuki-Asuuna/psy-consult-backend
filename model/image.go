package model

type GetImageUrlResponse struct {
	StatusCode int `json:"status_code"`
	Success    struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"success"`
	Image struct {
		Url string `json:"url"`
	} `json:"image"`
	StatusTxt string `json:"status_txt"`
}
