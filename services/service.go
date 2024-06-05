package services

import (
	"URLSHORTNER/utils"
	"log"
)

type ShortenURL interface {
	DomainCount() []utils.DomainCountList
	GetShorternURL(url string) (string, error)
	GetURL(shortUrl string) (string, bool)
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

// get the Domain count return 3  highest domains
func (s ShortenURLService) DomainCount() []utils.DomainCountList {
	return utils.GetHighestCountDomain(s.domainCount)
}

// GetShorternURL is used to make the short url
// that store long and short url on the map
// it will extract the domain and increment the domain map.
func (s ShortenURLService) GetShorternURL(url string) (string, error) {
	isExist := s.urlToShortUrl[url]
	if isExist != "" {
		log.Printf("the short Url already exist")
		return isExist, nil
	}
	domain, err := utils.GetDomain(url)
	if err != nil {
		log.Printf("unable to get domain from url %s got error %v", url, err)
		return "", err
	}
	s.domainCount[domain]++
	sUrl := utils.GetShortUrl(url)
	s.urlToShortUrl[url] = sUrl
	s.shortUrlToUrl[sUrl] = url
	ShortenURL := s.urlToShortUrl[url]
	return ShortenURL, nil
}

// get the url form short url map.
func (s ShortenURLService) GetURL(url string) (string, bool) {
	url, exists := s.shortUrlToUrl[url]
	return url, exists
}
