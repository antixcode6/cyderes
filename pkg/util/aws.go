package util

import (
	"os"
)

//Loads the VirusTotal API key from an ENV var
func GetSecret() (apiKey string) {
	apikey := os.Getenv("VTKEY")
	return apikey
}
