package link

type Link struct {
	FullURL   string
	ShortCode string
}

func NewLink(fullURL string) (*Link, error) {
	sc, err := newShortCode(fullURL)
	if err != nil {
		return nil, err
	}

	return &Link{
		ShortCode: sc,
		FullURL:   fullURL,
	}, nil
}
