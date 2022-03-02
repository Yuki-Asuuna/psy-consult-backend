package model

type GetImageUrlResponse struct {
	StatusCode int `json:"status_code"`
	Success    struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"success"`
	Image struct {
		Name             string      `json:"name"`
		Extension        string      `json:"extension"`
		Size             int         `json:"size"`
		Width            int         `json:"width"`
		Height           int         `json:"height"`
		Date             string      `json:"date"`
		DateGmt          string      `json:"date_gmt"`
		StorageId        interface{} `json:"storage_id"`
		Description      interface{} `json:"description"`
		Nsfw             string      `json:"nsfw"`
		Md5              string      `json:"md5"`
		Storage          string      `json:"storage"`
		OriginalFilename string      `json:"original_filename"`
		OriginalExifdata interface{} `json:"original_exifdata"`
		Views            string      `json:"views"`
		IdEncoded        string      `json:"id_encoded"`
		Filename         string      `json:"filename"`
		Ratio            float64     `json:"ratio"`
		SizeFormatted    string      `json:"size_formatted"`
		Mime             string      `json:"mime"`
		Bits             int         `json:"bits"`
		Channels         interface{} `json:"channels"`
		Url              string      `json:"url"`
		UrlViewer        string      `json:"url_viewer"`
		Thumb            struct {
			Filename      string      `json:"filename"`
			Name          string      `json:"name"`
			Width         int         `json:"width"`
			Height        int         `json:"height"`
			Ratio         int         `json:"ratio"`
			Size          int         `json:"size"`
			SizeFormatted string      `json:"size_formatted"`
			Mime          string      `json:"mime"`
			Extension     string      `json:"extension"`
			Bits          int         `json:"bits"`
			Channels      interface{} `json:"channels"`
			Url           string      `json:"url"`
		} `json:"thumb"`
		Medium struct {
			Filename      string      `json:"filename"`
			Name          string      `json:"name"`
			Width         int         `json:"width"`
			Height        int         `json:"height"`
			Ratio         float64     `json:"ratio"`
			Size          int         `json:"size"`
			SizeFormatted string      `json:"size_formatted"`
			Mime          string      `json:"mime"`
			Extension     string      `json:"extension"`
			Bits          int         `json:"bits"`
			Channels      interface{} `json:"channels"`
			Url           string      `json:"url"`
		} `json:"medium"`
		ViewsLabel string `json:"views_label"`
		DisplayUrl string `json:"display_url"`
		HowLongAgo string `json:"how_long_ago"`
	} `json:"image"`
	StatusTxt string `json:"status_txt"`
}
