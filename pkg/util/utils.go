package util

import (
	"crypto/sha256"
	"encoding/hex"
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

func hashURL(in string) (hash string) {
	h := sha256.New()
	h.Write([]byte(in))
	hash = hex.EncodeToString(h.Sum(nil))
	return
}
