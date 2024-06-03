package main

import (
	"net/http"
	"testing"

	"github.com/pharsha1995/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app  := newTestApplication()
	server := newTestServer(t, app.routes())
	defer server.Close()

	code, _, body := server.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}