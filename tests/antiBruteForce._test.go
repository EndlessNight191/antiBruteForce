package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	app "test/internal"
	"testing"
    "github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

func getOk(body io.ReadCloser) bool {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response struct {
		IsAccess bool `json:"isAccess"`
	}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	return response.IsAccess
}	

func TestAuthService(t *testing.T) {
    app.Run()
    port := viper.GetInt("port")
    baseUrl := fmt.Sprintf("localhost:%d/api", port)
    
    t.Run("whitelist", func(t *testing.T) {
		body := `{"ip": "whiteTest", "login": "whiteTest", "password": "test"}`
        resp, err := http.Post(baseUrl, "application/json", strings.NewReader(body))
        assert.NoError(t, err)
        assert.Equal(t, true, getOk(resp.Body))
    })

    t.Run("blacklist", func(t *testing.T) {
		body := `{"ip": "blackTest", "login": "blackTest", "password": "test"}`
        resp, err := http.Post(baseUrl, "application/json", strings.NewReader(body))
        assert.NoError(t, err)
        assert.Equal(t, false, getOk(resp.Body))
    })

    t.Run("normal", func(t *testing.T) {
		body := `{"ip": "127.0.0.2", "login": "testOk", "password": "testOk"}`
        resp, err := http.Post(baseUrl, "application/json", strings.NewReader(body))
        assert.NoError(t, err)
        assert.Equal(t, true, resp.Body)
    })

    t.Run("bruteforce", func(t *testing.T) {
		body := `{"ip": "127.0.0.1", "login": "test", "password": "test"}`
        count := viper.GetInt("MAX_LIMIT_COMMON")

        for i := 0; i <= count; i++ {
            resp, err := http.Post(baseUrl, "application/json", strings.NewReader(body))
            assert.NoError(t, err)
            if i < 10 {
                assert.Equal(t, true, getOk(resp.Body))
            } else {
                assert.Equal(t, false, getOk(resp.Body))
            }
        }
    })

}
