package main

import (
	"testing"
	"net/http"
)

func TestHTTPHeaderFilters(t *testing.T) {
	filters := HTTPHeaderFilters{}

	err := filters.Set("Header1:^$")
	if err != nil {
		t.Error("Should not error on Header1:^$")
	}

	err = filters.Set("Header2:^:$")
	if err != nil {
		t.Error("Should not error on Header2:^:$")
	}

	err = filters.Set("Header3-^$")
	if err == nil {
		t.Error("Should error on Header2:^:$")
	}

	req := http.Request{}
	req.Header = make(map[string][]string)
	req.Header.Add("Header1", "")
	req.Header.Add("Header2", ":")
	req.Header.Add("Header3", "Irrelevant")

	if(!filters.Good(&req)) {
		t.Error("Request should pass filters")
	}
}
