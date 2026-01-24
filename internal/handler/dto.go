package handler

type linkResponse struct {
	ShortCode string `json:"short_code"`
	FullURL   string `json:"full_url"`
}

type createRequest struct {
	FullURL string `json:"full_url" binding:"required,url"`
}
