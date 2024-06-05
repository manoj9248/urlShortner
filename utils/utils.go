package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/url"
	"sort"
	"strconv"
)

type DomainCountList struct {
	Key   string
	Value int
}

// create the short url based on md5 hash and return the first 7 hash code
func GetShortUrl(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:7])
}

// create the short url based on md5 hash with counter and return the first 7 hash code
func GetShortUrlwithCounter(url string, count int) string {
	hash := md5.Sum([]byte(url + strconv.Itoa(count)))
	return hex.EncodeToString(hash[:7])
}

// parse and return the Domain name from url.
func GetDomain(inputUrl string) (string, error) {
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		log.Printf("unable to get domain %v", err)
		return "", err
	}
	domain := parsedURL.Host
	return domain, nil
}

// loop the Domain map in descending and return top 3 Domain count
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
