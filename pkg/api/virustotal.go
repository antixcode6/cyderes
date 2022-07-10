package api

import (
	"log"
	"time"

	"github.com/VirusTotal/vt-go"
	"github.com/antixcode6/cyderes/pkg/util"
)

type Response struct {
	Rep         interface{} `json:"Reputation"`
	Url         interface{} `json:"Url"`
	LastSubDate time.Time   `json:"Last Submission Date"`
	DNS         DnsRequest  `json:"DNS Query"`
}

//Uses the Virus total sdk/client provided by VT themselves and returns some information
func QueryVirusTotal(url string) Response {
	var r Response
	apikey := util.GetSecret()
	client := vt.NewClient(apikey)

	resp, err := client.GetObject(vt.URL("urls/%s", url))
	if err != nil {
		log.Fatal(err)
	}
	ls, err := resp.GetTime("last_submission_date")
	if err != nil {
		log.Fatal(err)
	}

	respRep, _ := resp.Get("reputation")
	respUrl, _ := resp.Get("url")
	r = Response{LastSubDate: ls, Url: respUrl, Rep: respRep}
	return r
}
