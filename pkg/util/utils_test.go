package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCandidates = []string{"google.com", "https://cloud.bugatti", "172.253.122.138", "2607:f8b0:4004:837::200e"}

func TestValidateQuery(t *testing.T) {
	urls := make([]string, 0)
	urls = append(urls, testCandidates...)

	for i := range urls {
		testQuery, err := ValidateQuery(urls[i])
		if err != nil {
			assert.NotNil(t, err)
			break
		}
		assert.Nil(t, err)
		assert.NotEmpty(t, testQuery)
		fmtUrl := buildURL(urls[i])
		hashed := hashURL(fmtUrl)
		assert.Equal(t, hashed, testQuery)
	}

}

func TestStripUrl(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{
			[]string{"https://cloud.bugatti", "https://cloud.bugatti:900"},
			"cloud.bugatti",
		},
	}
	for i, test := range tests {
		strippedUrl := StripURL(tests[i].in[i])
		assert.Equal(t, test.want, strippedUrl)
	}

}

func TestAws(t *testing.T) {
	os.Setenv("VTKEY", "testKey")
	expected := os.Getenv("VTKEY")
	actual := GetSecret()
	assert.Equal(t, expected, actual)
}
