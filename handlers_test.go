package main

import (
	"testing"
	"fmt"
	//"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func TestPublicKeysHandler(t *testing.T) {
	record := httptest.NewRecorder()
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "publickeys/"},
		Form:   url.Values{
			  "addresses": {"44ljb4mt7rbo3fue@scramble.io"},
			},
		}
	publicKeysHandler(record, req)
	fmt.Println(record.Code, record.Body.String())
}
