package rest

import (
	"shorten/internal/domain"
)

func toCreateResponse(sc domain.ShortCode, host string) linkResponse {
	return linkResponse{
		ShortCode: host + string(sc),
	}
}

func toGetResponse(link *domain.Link, host string) linkResponse {
	return linkResponse{
		ShortCode: host + string(link.ShortCode()),
	}
}
