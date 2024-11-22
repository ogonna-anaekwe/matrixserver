package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFlatten(t *testing.T) {
	h := Handler{}
	server := httptest.NewServer(http.HandlerFunc(h.HandlerFlatten))
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %v. Got %v ", http.StatusOK, res.StatusCode)
	}
}
