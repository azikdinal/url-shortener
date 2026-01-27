package rest

type linkResponse struct {
	ShortCode string `json:"short_code"`
}

type createRequest struct {
	FullURL string `json:"full_url" binding:"required,url"`
}
