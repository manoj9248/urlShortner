package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetShortUrl(t *testing.T) {
	resp := GetShortUrl("https://dev.to/devniklesh/crud-api-with-go-gin-framework-production-ready-52jd")
	require.Equal(t, resp, "9cfb054a4b9e7a")
}
func TestGetDomain(t *testing.T) {
	resp, err := GetDomain("https://github.com/manoj9248/urlShortner/commits/master/")
	require.NoError(t, err)
	require.Equal(t, resp, "github.com")
}
func TestGetHighestCountDomain(t *testing.T) {
	reqMap := map[string]int{}
	reqMap["test1"] = 3
	reqMap["test2"] = 1
	reqMap["test3"] = 19
	reqMap["test4"] = 9
	resp := GetHighestCountDomain(reqMap)
	require.Equal(t, len(resp), 3)
	require.Equal(t, resp[0].Key, "test3")
	require.Equal(t, resp[1].Key, "test4")
}
