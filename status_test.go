package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStatusPage(t *testing.T) {
	assert := assert.New(t)

	site := Site{Config: Config{Interval: time.Duration(9)}}
	sitest := Sitest{ConfigFile: "/tmp/mock", Sites: map[string]*Site{"http://example.com": &site}}
	ts := httptest.NewServer(sitest)
	defer ts.Close()

	resp, _ := http.Get(ts.URL)
	assert.Equal(resp.StatusCode, 200, "response from status handler")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(strings.Contains(string(body[:]), "http://example.com"), true, "expected body part")
}
