package api

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
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
		key := GetSecret()
		os.Setenv("VTKEY", key.Virustotal)
		response, err := Ingest(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Url)
	}
}

type SecretData struct {
	Virustotal string `json:"virustotal"`
}

var (
	secretName   string = "VT-Key"
	region       string = "us-east-1"
	versionStage string = "AWSCURRENT"
)

//Calls on AWS Go SDK to get API key from secrets manager
//
//yoinked from https://gist.github.com/xlyk/f2f2246ee259415c05f84eb21218ac73 which is much better than the auto generated
//aws boilerplate
func GetSecret() SecretData {
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err.Error())
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	var secretData SecretData
	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		panic(err.Error())
	}

	return secretData
}
