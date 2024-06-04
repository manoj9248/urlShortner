package endpoints

type ShortenURL interface {
	DomainCount(domainName string) int
	GetShorternerURL(url string) string
	GetURL(shortUrl string) string
}
type ShortenURLService struct {
	urlToShort map[string]string
	shortToUrl map[string]string
	count      map[string]int
}

func NewShortenURLServicer() ShortenURL {
	return &ShortenURLService{
		urlToShort: make(map[string]string),
		shortToUrl: make(map[string]string),
		count:      make(map[string]int),
	}
}
func (s ShortenURLService) DomainCount(domainName string) int {
	return 0
}
func (s ShortenURLService) GetShorternerURL(url string) string {
	return url
}
func (s ShortenURLService) GetURL(url string) string {
	return s.shortToUrl[url]
}
