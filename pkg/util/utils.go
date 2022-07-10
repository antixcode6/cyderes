package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"net/url"
	"strings"
)

//ValidateQuery checks an incoming query value to try and determine if it's a
//url or otherwise.
//
//If input url value is missing http:// it will build the url and then hash that returned value
//
//Returns string(url -> SHA256), error
func ValidateQuery(message string) (string, error) {
	if strings.Contains(message, "http://") { //correct url fmt
		url := hashURL(message)
		return url, nil
	} else if strings.Contains(message, ".") { //url missing http://
		uri := buildURL(message)
		url := hashURL(uri)
		return url, nil
	} else if strings.Contains(message, ":") {
		return "", fmt.Errorf("ipv6 addresses are not supported")
	} else { //probably already a hash
		return message, nil
	}
}

func buildURL(in string) (url string) {
	if strings.Contains(in, ".") && !strings.Contains(in, "http://") {
		url = `http://` + in + `/`
		return
	}
	return
}

func StripURL(in string) string {
	u, err := url.Parse(in)
	if u.Host != "" {
		if host, _, err := net.SplitHostPort(u.Host); err == nil {
			return host
		} else {
			return u.Host
		}
	}
	if err != nil {
		return ""
	}
	return ""
}

func hashURL(in string) (hash string) {
	h := sha256.New()
	h.Write([]byte(in))
	hash = hex.EncodeToString(h.Sum(nil))
	return
}
