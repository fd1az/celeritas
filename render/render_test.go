package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "foo", true, "no error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jey template"},
	{"jet_page_no_template", "jet", "foo", true, "no error rendering non-existent jet template, when one is expected"},
	{"invalid_renderer engine", "foo", "home", true, "no error rendering non-existent template engine"},
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)

		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)

		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s ", e.name, e.errorMessage)
			}

		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"

	err = testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error("Error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error("Error rendering page", err)
	}
}
