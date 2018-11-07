# sitest

[![Build Status](https://travis-ci.org/tomplus/sitest.svg?branch=master)](https://travis-ci.org/tomplus/sitest)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomplus/airly-exporter)](https://goreportcard.com/report/github.com/tomplus/sitest)

Prometheus exporter to test HTTP(s) sites

## Overview

Sitest (web-sites + test = sitest) is a simple tool to check websites. It gets a defined url and create metrics
with respons code, response time, page size, page hash etc. The metrics are exposed in the Prometheus format
and can be used to trigger alerts if the website is down, content was changed or responds slowly.

## Configuration

A listen address and a path to configuration file can be changed via command line switches:

```
Usage of sitest:
  -config_file string
        path to config-file (default "./sitest.yaml")
  -listen_addr string
        listen address (default "0.0.0.0:8080")
```

List of URL with parameters are stored in the configuration file. It's yaml file with a simple stucture:

```
# default section for each site
default:
  interval: 1m  # how often test site

# site list with configuration
sites:
  "https://golang.org":
    interval: 15s  # own settings

  "https://api.myip.com/": {}  # default configuration

```

## Running

TODO

## List of exposed metrics

Metrics:
```
# HELP sitest_code Response code
# TYPE sitest_code gauge
sitest_code{site="http://example.com/"} 200
sitest_code{site="https://api.myip.com/"} 200
sitest_code{site="https://golang.org"} 200
# HELP sitest_count Total number of performed check
# TYPE sitest_count counter
sitest_count{site="http://example.com/"} 1
sitest_count{site="https://api.myip.com/"} 1
sitest_count{site="https://golang.org"} 6
# HELP sitest_duration_seconds Histogram of request duration
# TYPE sitest_duration_seconds histogram
sitest_duration_seconds_bucket{site="http://example.com/",le="0.005"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.01"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.025"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.05"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.1"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.25"} 0
sitest_duration_seconds_bucket{site="http://example.com/",le="0.5"} 1
sitest_duration_seconds_bucket{site="http://example.com/",le="1"} 1
sitest_duration_seconds_bucket{site="http://example.com/",le="2.5"} 1
sitest_duration_seconds_bucket{site="http://example.com/",le="5"} 1
sitest_duration_seconds_bucket{site="http://example.com/",le="10"} 1
sitest_duration_seconds_bucket{site="http://example.com/",le="+Inf"} 1
sitest_duration_seconds_sum{site="http://example.com/"} 0.294034706
sitest_duration_seconds_count{site="http://example.com/"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.005"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.01"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.025"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.05"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.1"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.25"} 0
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="0.5"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="1"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="2.5"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="5"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="10"} 1
sitest_duration_seconds_bucket{site="https://api.myip.com/",le="+Inf"} 1
sitest_duration_seconds_sum{site="https://api.myip.com/"} 0.377305916
sitest_duration_seconds_count{site="https://api.myip.com/"} 1
sitest_duration_seconds_bucket{site="https://golang.org",le="0.005"} 0
sitest_duration_seconds_bucket{site="https://golang.org",le="0.01"} 0
sitest_duration_seconds_bucket{site="https://golang.org",le="0.025"} 0
sitest_duration_seconds_bucket{site="https://golang.org",le="0.05"} 0
sitest_duration_seconds_bucket{site="https://golang.org",le="0.1"} 0
sitest_duration_seconds_bucket{site="https://golang.org",le="0.25"} 5
sitest_duration_seconds_bucket{site="https://golang.org",le="0.5"} 6
sitest_duration_seconds_bucket{site="https://golang.org",le="1"} 6
sitest_duration_seconds_bucket{site="https://golang.org",le="2.5"} 6
sitest_duration_seconds_bucket{site="https://golang.org",le="5"} 6
sitest_duration_seconds_bucket{site="https://golang.org",le="10"} 6
sitest_duration_seconds_bucket{site="https://golang.org",le="+Inf"} 6
sitest_duration_seconds_sum{site="https://golang.org"} 1.089489004
sitest_duration_seconds_count{site="https://golang.org"} 6
# HELP sitest_hash Page hash
# TYPE sitest_hash gauge
sitest_hash{site="http://example.com/"} 2.2135241933328225e+18
sitest_hash{site="https://api.myip.com/"} 1.7321780815323423e+19
sitest_hash{site="https://golang.org"} 1.3824383691894628e+19
# HELP sitest_length Page length
# TYPE sitest_length gauge
sitest_length{site="http://example.com/"} 1270
sitest_length{site="https://api.myip.com/"} 51
sitest_length{site="https://golang.org"} 8099
```
