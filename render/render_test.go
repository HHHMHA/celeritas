package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, _ := http.NewRequest("GET", "/some-url", nil)
	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err := testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error("Error rendering page", err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page with jet", err)
	}
}
