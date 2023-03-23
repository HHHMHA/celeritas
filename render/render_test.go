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

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page with jet", err)
	}

	testRenderer.Renderer = ""

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("Error rendering with no engine", err)
	}
}

func TestRender_GoPage(t *testing.T) {
	r, _ := http.NewRequest("GET", "/some-url", nil)
	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err := testRenderer.GoPage(w, r, "home", nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.GoPage(w, r, "no-file", nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	r, _ := http.NewRequest("GET", "/some-url", nil)
	w := httptest.NewRecorder()

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	err := testRenderer.JetPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.JetPage(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}
}
