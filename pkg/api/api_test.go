package api

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var urls = []string{"google.com", "eggwald.com", "172.253.122.138", "cloud.bugatti"}

func TestNetQuery(t *testing.T) {
	for i := range urls {
		resp := QueryNet(urls[i])
		assert.NotNil(t, resp)
	}

}

func TestHandler(t *testing.T) {
	m := make(map[string]string)
	m["req"] = "google.com"
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{QueryStringParameters: m},
			expect:  "google.com",
			err:     nil,
		},
	}
	for _, test := range tests {
		response, err := Ingest(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Url)
	}
}
