package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRoutingGet(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/tweet", srv.URL))

	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status NotAllowed; got %v", res.Status)
	}
}

func TestRoutingPOST(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	var jsonStr = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque interdum rutrum sodales. Nullam mattis fermentum libero, non volutpat."

	url := fmt.Sprintf("%s/tweet?tweet_data=%s", srv.URL, url.QueryEscape(jsonStr))
	res, err := http.Post(url, "application/json", nil)

	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	var tweet Tweet
	if err := json.Unmarshal(b, &tweet); err != nil {
		t.Fatalf("could not read response: %v", err)
	}

}
