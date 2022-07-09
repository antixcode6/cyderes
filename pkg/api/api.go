package api

import (
	"log"
	"net/http"

	"github.com/antixcode6/cyderes/pkg/util"
)

func Ingest(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Query().Get("req")
	hashedRequest, err := util.ValidateQuery(request)
	if err == nil {
		QueryVirusTotal(hashedRequest)
		QueryWhoIs(request)
	}
	log.Println(err)
}
