package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/t3h2mas/passpi/hash"
)

func TestHandleHash(t *testing.T) {
	// disable the global logger
	log.SetOutput(ioutil.Discard)

	// create test server
	s := &server{hash: &hash.HashSha512{}}
	ts := httptest.NewServer(s.handleHash())
	defer ts.Close()

	// it should not accept GET requests
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status %v, got %v", http.StatusMethodNotAllowed, res.StatusCode)
	}

	// it should not accept bodies missing the 'password' key
	res, err = http.PostForm(ts.URL, url.Values{
		"passwor": {"bazbarfoo"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, res.StatusCode)
	}

	// it should return the correct hash otherwise
	res, err = http.PostForm(ts.URL, url.Values{
		"password": {"bazbarfoo"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	expectedHash := "4YhI0sJo3g14rZkAzwUEtOMpbdcUuTZvCuVxRaXgs0FCjyx8uMhS68aKhaxbaMPCAf7v75ODhB8zg5mzco5c6w=="
	if string(body) != expectedHash {
		t.Errorf("expected body '%s', got '%s'", expectedHash, body)
	}
}
