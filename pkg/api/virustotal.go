package api

import (
	"log"

	"github.com/VirusTotal/vt-go"
	"github.com/antixcode6/cyderes/pkg/util"
)

func QueryVirusTotal(url string) {

	apikey := util.GetSecret()
	client := vt.NewClient(apikey.Virustotal)

	resp, err := client.GetObject(vt.URL("urls/%s", url))
	if err != nil {
		log.Fatal(err)
	}
	ls, err := resp.GetTime("last_submission_date")
	if err != nil {
		log.Fatal(err)
	}

	rep, _ := resp.Get("reputation")
	respurl, _ := resp.Get("url")
	log.Printf("%s (%s) was submitted for the last time on %v\n", respurl, resp.ID(), ls)
	log.Printf("Reputation of %s : %s", respurl, rep)
}
