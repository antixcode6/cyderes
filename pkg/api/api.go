package api

import (
	"strings"

	"github.com/antixcode6/cyderes/pkg/util"
	"github.com/aws/aws-lambda-go/events"
)

type IQ struct {
	Url string `json:"url"`
}

//Ingest takes a query param from AWS API Gateway and passes that along to
//	ValidateQuery(url)
func Ingest(request events.APIGatewayProxyRequest) (Response, error) {
	url := request.QueryStringParameters["req"]
	if strings.Contains(url, ".") { //probably a URL or IP
		if strings.Contains(url, "https://") {
			url = util.StripURL(url)
		}
		hashedRequest, err := util.ValidateQuery(url)
		dnsQuery := QueryNet(url)
		if err == nil {
			resp := QueryVirusTotal(hashedRequest)
			return Response{
				LastSubDate: resp.LastSubDate,
				Url:         url,
				Rep:         resp.Rep,
				DNS:         dnsQuery,
			}, nil
		} else {
			return Response{}, err
		}
	} else { //probably already a hash
		resp := QueryVirusTotal(url)
		return Response{
			LastSubDate: resp.LastSubDate,
			Url:         url,
			Rep:         resp.Rep,
		}, nil
	}
}
