package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStatusPage(t *testing.T) {
	assert := assert.New(t)

	sitest := Sitest{ConfigFile: "/tmp/mock", Sites: map[string]Config{"http://example.com": Config{Interval: time.Duration(9)}}}
	ts := httptest.NewServer(sitest)
	defer ts.Close()

	resp, _ := http.Get(ts.URL)
	defer resp.Body.Close()

	assert.Equal(resp.StatusCode, 200, "response from prometheus handler")

	expectedBody := `<html>
<head>
<title>Sitest</title>
</head>
<body>
<p>
<h1>Defined sites:</h1>
<ul>

   <li><strong>http://example.com</strong>: {9ns}</li>

</ul>
configuration from file /tmp/mock</pre>
</p>
<p>
Metrics <a href="/metrics">/metrics</a>
</p>
</body>
</html>`

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(string(body[:]), expectedBody, "expected body")
}
