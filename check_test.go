package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckSite(t *testing.T) {

	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))
	defer ts.Close()

	res, err := checkSite(ts.URL)

	assert.Nil(err, "assert error")
	assert.Equal(res.StatusCode, 200, "assert http code")
	assert.Equal(res.Length, 2, "assert length")
	assert.Equal(res.Hash, uint64(656748223988434799), "assert hash")
}
