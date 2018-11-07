package main

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMetricCounting(t *testing.T) {
	assert := assert.New(t)

	dr := prometheus.NewRegistry()
	var reg prometheus.Registerer = dr
	var gat prometheus.Gatherer = dr

	promColl := PromCollectors{}
	promColl.Register(reg)

	result := Result{StatusCode: 200, Length: 5, Duration: 6, Hash: 7}
	promColl.Update("my-site.example.com", result, nil)

	srv := httptest.NewServer(promhttp.HandlerFor(gat, promhttp.HandlerOpts{}))
	defer srv.Close()
	resp, _ := http.Get(srv.URL)

	assert.Equal(resp.StatusCode, 200, "response from prometheus handler")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	expectedBody := `# HELP sitest_code Response code
# TYPE sitest_code gauge
sitest_code{site="my-site.example.com"} 200
# HELP sitest_count Total number of performed check
# TYPE sitest_count counter
sitest_count{site="my-site.example.com"} 1
# HELP sitest_duration_seconds Histogram of request duration
# TYPE sitest_duration_seconds histogram
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.005"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.01"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.025"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.05"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.1"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.25"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="1"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="2.5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="10"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="+Inf"} 1
sitest_duration_seconds_sum{site="my-site.example.com"} 6e-09
sitest_duration_seconds_count{site="my-site.example.com"} 1
# HELP sitest_hash Page hash
# TYPE sitest_hash gauge
sitest_hash{site="my-site.example.com"} 7
# HELP sitest_length Page length
# TYPE sitest_length gauge
sitest_length{site="my-site.example.com"} 5
`
	assert.Equal(string(body[:]), expectedBody, "expected body")

}

func TestMetricError(t *testing.T) {
	assert := assert.New(t)

	dr := prometheus.NewRegistry()
	var reg prometheus.Registerer = dr
	var gat prometheus.Gatherer = dr

	promColl := PromCollectors{}
	promColl.Register(reg)

	result := Result{}
	promColl.Update("my-site.example.com", result, errors.New("test-error"))

	srv := httptest.NewServer(promhttp.HandlerFor(gat, promhttp.HandlerOpts{}))
	defer srv.Close()
	resp, _ := http.Get(srv.URL)

	assert.Equal(resp.StatusCode, 200, "response from prometheus handler")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	expectedBody := `# HELP sitest_code Response code
# TYPE sitest_code gauge
sitest_code{site="my-site.example.com"} 0
# HELP sitest_count Total number of performed check
# TYPE sitest_count counter
sitest_count{site="my-site.example.com"} 1
# HELP sitest_duration_seconds Histogram of request duration
# TYPE sitest_duration_seconds histogram
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.005"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.01"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.025"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.05"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.1"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.25"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="0.5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="1"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="2.5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="5"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="10"} 1
sitest_duration_seconds_bucket{site="my-site.example.com",le="+Inf"} 1
sitest_duration_seconds_sum{site="my-site.example.com"} 0
sitest_duration_seconds_count{site="my-site.example.com"} 1
# HELP sitest_error Total number of error
# TYPE sitest_error counter
sitest_error{site="my-site.example.com"} 1
# HELP sitest_hash Page hash
# TYPE sitest_hash gauge
sitest_hash{site="my-site.example.com"} 0
# HELP sitest_length Page length
# TYPE sitest_length gauge
sitest_length{site="my-site.example.com"} 0
`
	assert.Equal(string(body[:]), expectedBody, "expected body")

}
