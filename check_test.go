package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckSite(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))
	defer ts.Close()

	res, err := checkSite(ts.URL)
	fmt.Println(res, err)

	if err != nil {
		t.Errorf("assert error %v is not nil", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("assert http code %v is not 200", res.StatusCode)
	}

	if res.Length != 2 {
		t.Errorf("assert length %v is not 2", res.Length)
	}

	if res.Hash != 656748223988434799 {
		t.Errorf("assert hash %v is not 656748223988434799", res.Hash)
	}

}
