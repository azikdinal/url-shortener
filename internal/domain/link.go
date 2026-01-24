package domain

type Link struct {
	fullURL   FullURL
	shortCode ShortCode
}

func NewLink(sc ShortCode, fu FullURL) *Link {
	return &Link{
		shortCode: sc,
		fullURL:   fu,
	}
}

func (l *Link) ShortCode() ShortCode {
	return l.shortCode
}

func (l *Link) FullURL() FullURL {
	return l.fullURL
}

func NewLinkWithGeneratedCode(fu FullURL) *Link {
	sc := generateShortCode()
	return &Link{
		fullURL:   fu,
		shortCode: sc,
	}
}
