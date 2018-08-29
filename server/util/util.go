package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Brightscout/mattermost-plugin-googledrive/server/config"
	"github.com/pkg/errors"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	// ExpirySingleUse ...
	// single use URLs are deleted on first use.
	ExpirySingleUse = 0

	// ExpiryPermanent ...
	// permanent URLs can be used any number of times
	ExpiryPermanent = 1
)

// Split ...
// cant use strings.split as it includes empty string if deliminator
// is the last character in input string
func Split(data string, delim rune) []string {
	return strings.FieldsFunc(data, func(c rune) bool { return c == delim })
}

// RandString ...
// Copied from https://stackoverflow.com/a/31832326/1589165
// About 4x faster than simply choosing a random character from array
func RandString(n int) string {
	// we create a new source of randomness for each execution to
	// keep this function thread/goroutine safe
	var rndSrc = rand.NewSource(time.Now().UnixNano())

	result := make([]byte, n)
	// A rndSrc.Int63() generates 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := n-1, rndSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rndSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(result)
}

// ShortenURL ...
func ShortenURL(longURL string, expiryType ...int) (string, error) {
	var expiry int

	if len(expiryType) == 0 {
		expiry = ExpirySingleUse
	} else {
		expiry = expiryType[0]
	}

	var key = RandString(4)
	var shortURL = fmt.Sprintf("/plugins/%s/redirect?key=%s", config.PluginName, key)

	// URL saved is in format <expiry type><generated URL key>.
	var err = config.Mattermost.KeyValueStore().Set(config.UrlMappingKeyPrefix+key, []byte(strconv.Itoa(expiry)+longURL))
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// LengthenURL ...
func LengthenURL(shortURL string) (string, error) {
	var longURL, err = config.Mattermost.KeyValueStore().Get(config.UrlMappingKeyPrefix + shortURL)
	if err != nil {
		return "", err
	}
	if len(longURL) == 0 {
		return "", errors.New("No such URL could be found")
	}

	var expiry = longURL[0]
	longURL = longURL[1:]

	if expiry == ExpirySingleUse {
		defer config.Mattermost.KeyValueStore().Delete(config.UrlMappingKeyPrefix + shortURL)
	}

	return string(longURL), nil
}

// SplitArgs ...
func SplitArgs(s string) ([]string, error) {
	var indexes = regexp.MustCompile("\"").FindAllStringIndex(s, -1)
	if len(indexes)%2 != 0 {
		return []string{}, errors.New("quotes not closed")
	}

	indexes = append([][]int{{0, 0}}, indexes...)

	if indexes[len(indexes)-1][1] < len(s) {
		indexes = append(indexes, [][]int{{len(s), 0}}...)
	}

	var args []string
	for i := 0; i < len(indexes)-1; i++ {
		var start = indexes[i][1]
		var end = Min(len(s), indexes[i+1][0])

		if i%2 == 0 {
			args = append(args, strings.Split(strings.Trim(s[start:end], " "), " ")...)
		} else {
			args = append(args, s[start:end])
		}

	}

	return args, nil
}

// Min ...
// Because math.Min is for floats and
// casting to and from floats is dangerous.
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Max ...
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
