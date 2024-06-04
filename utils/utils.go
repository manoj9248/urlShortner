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
	var domainList []DomainCountList
	for k, v := range domainCount {
		domainList = append(domainList, DomainCountList{Key: k, Value: v})
	}
	sort.Slice(domainList, func(i, j int) bool {
		return domainList[i].Value > domainList[j].Value
	})
	if len(domainList) > 3 {
		return domainList[:3]
	}
	return domainList
}
