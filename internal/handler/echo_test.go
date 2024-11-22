package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerEcho(t *testing.T) {
	h := Handler{}
	server := httptest.NewServer(http.HandlerFunc(h.HandlerEcho))
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %v. Got %v ", http.StatusOK, res.StatusCode)
	}
}
