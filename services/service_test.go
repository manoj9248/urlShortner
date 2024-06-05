package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDomainCount(t *testing.T) {
	svcObj := helper()
	resp := svcObj.DomainCount()
	require.Equal(t, len(resp), 1)
}
func TestGetShorternURL(t *testing.T) {
	svcObj := helper()
	resp, err := svcObj.GetShorternURL("https://github.com/manoj")
	require.NoError(t, err)
	require.Equal(t, resp, "452785")
	resp, err = svcObj.GetShorternURL("https://github.com/manoj1")
	require.NoError(t, err)
	require.Equal(t, resp, "90746e7ac11204")
}
func TestGetUrl(t *testing.T) {
	svcObj := helper()
	resp, _ := svcObj.GetURL("452785")
	require.Equal(t, resp, "https://github.com/manoj")
	resp, exist := svcObj.GetURL("https://github.com/manoj1")
	require.Equal(t, resp, "")
	require.Equal(t, exist, false)
}
func helper() ShortenURLService {
	urlToShortUrl := make(map[string]string)
	shortUrlToUrl := make(map[string]string)
	domainCount := make(map[string]int)
	urlToShortUrl["https://github.com/manoj"] = "452785"
	shortUrlToUrl["452785"] = "https://github.com/manoj"
	domainCount["github.com"] = 1
	urlobj := ShortenURLService{
		urlToShortUrl: urlToShortUrl,
		shortUrlToUrl: shortUrlToUrl,
		domainCount:   domainCount,
	}
	return urlobj
}
