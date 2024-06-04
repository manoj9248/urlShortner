package services

import "URLSHORTNER/utils"

type ShortenURL interface {
	DomainCount() []utils.DomainCountList
	GetShorternURL(url string) (string, error)
	GetURL(shortUrl string) string
}
type ShortenURLService struct {
	urlToShortUrl map[string]string
	shortUrlToUrl map[string]string
	domainCount   map[string]int
}

func NewShortenURLServicer() ShortenURL {
	return &ShortenURLService{
		urlToShortUrl: make(map[string]string),
		shortUrlToUrl: make(map[string]string),
		domainCount:   make(map[string]int),
	}
}
func (s ShortenURLService) DomainCount() []utils.DomainCountList {
	return utils.GetHighestCountDomain(s.domainCount)
}
func (s ShortenURLService) GetShorternURL(url string) (string, error) {
	isExist := s.urlToShortUrl[url]
	if isExist != "" {
		return isExist, nil
	}
	domain, err := utils.GetDomain(url)
	if err != nil {
		return "", err
	}
	s.domainCount[domain]++
	sUrl := utils.GetShortUrl(url)
	s.urlToShortUrl[url] = sUrl
	s.shortUrlToUrl[sUrl] = url
	ShortenURL := s.urlToShortUrl[url]
	return ShortenURL, nil
}
func (s ShortenURLService) GetURL(url string) string {
	return s.shortUrlToUrl[url]
}
