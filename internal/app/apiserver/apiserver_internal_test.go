package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_HandleGetList(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	s.getTodoList().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello world")
}
