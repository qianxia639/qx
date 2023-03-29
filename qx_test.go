package qx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQx(t *testing.T) {
	engine := New()

	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/hello", nil)

	engine.ServeHTTP(recorder, req)
}
