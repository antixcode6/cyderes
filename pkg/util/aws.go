package util

import (
	"os"
)

func GetSecret() (apiKey string) {
	apikey := os.Getenv("VTKEY")
	return apikey
}
