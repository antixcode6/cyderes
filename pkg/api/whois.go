package api

import (
	"fmt"

	"github.com/likexian/whois"
)

func QueryWhoIs(url string) {
	result, err := whois.Whois(url)
	if err == nil {
		fmt.Println(result)
	}
}
