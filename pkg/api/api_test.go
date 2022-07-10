package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCandidates = []string{"google.com", "eggwald.com", "172.253.122.138", "cloud.bugatti"}

func TestNetQuery(t *testing.T) {
	for i := range testCandidates {
		resp := QueryNet(testCandidates[i])
		assert.NotNil(t, resp)
	}

}

func TestVTQuery(t *testing.T) {
	for i := range testCandidates {
		resp := QueryVirusTotal(testCandidates[i])
		assert.NotNil(t, resp)
	}

}
