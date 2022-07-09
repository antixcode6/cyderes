package api

import (
	"log"

	"github.com/VirusTotal/vt-go"
)

func QueryVirusTotal(url string) {
	apikey := ""
	client := vt.NewClient(apikey)

	resp, err := client.GetObject(vt.URL("urls/%s", url))
	if err != nil {
		log.Fatal(err)
	}
	ls, err := resp.GetTime("last_submission_date")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File %s was submitted for the last time on %v\n", resp.ID(), ls)

}
