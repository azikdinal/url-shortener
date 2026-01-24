package handler

import (
	"shorten/internal/domain"
)

func toCreateResponse(link *domain.Link, host string) linkResponse {
	return linkResponse{
		FullURL:   string(link.FullURL()),
		ShortCode: host + string(link.ShortCode()),
	}
}

func toGetResponse(link *domain.Link, host string) linkResponse {
	return linkResponse{
		ShortCode: host + string(link.ShortCode()),
	}
}
