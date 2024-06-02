package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pharsha1995/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	resRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ping(resRecorder, req)
	res := resRecorder.Result()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}