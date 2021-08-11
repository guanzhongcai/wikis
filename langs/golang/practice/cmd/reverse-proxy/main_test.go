package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestRP(t *testing.T) {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("headers: ", r.Header)
	}))
	target, _ := url.Parse(backendServer.URL)
	fmt.Println("backendServer.URL::", backendServer.URL)

	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	proxyServer := httptest.NewServer(reverseProxy)

	fmt.Println("proxyServer.URL::", proxyServer.URL)
	req, err := http.NewRequest("POST", proxyServer.URL+"/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "abcdefg")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
}
