package util

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

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
