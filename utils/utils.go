package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
)

type DomainCountList struct {
	Key   string
	Value int
}

func GetShortUrl(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:7])
}
func GetDomain(inputUrl string) (string, error) {
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}
	domain := parsedURL.Host
	return domain, nil
}
func GetHighestCountDomain(domainCount map[string]int) []DomainCountList {
	var pairs []DomainCountList
	for k, v := range domainCount {
		pairs = append(pairs, DomainCountList{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	if len(pairs) > 3 {
		return pairs[:3]
	}
	return pairs
}
