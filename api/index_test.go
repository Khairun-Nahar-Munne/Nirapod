package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerServesHome(t *testing.T) {
	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/", nil))

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Errorf("Content-Type = %q, want text/html", ct)
	}
	if !strings.Contains(rr.Body.String(), "নিরাপদ") && !strings.Contains(rr.Body.String(), "NIRAPOD") {
		t.Error("expected home page body to mention নিরাপদ or NIRAPOD")
	}
}

func TestHandlerServesRobotsAndSitemap(t *testing.T) {
	t.Setenv("SITE_URL", "https://nirapod-browse.vercel.app")

	for path, want := range map[string]string{
		"/robots.txt":  "User-agent: *",
		"/sitemap.xml": "<urlset",
	} {
		rr := httptest.NewRecorder()
		Handler(rr, httptest.NewRequest(http.MethodGet, path, nil))

		if rr.Code != http.StatusOK {
			t.Fatalf("%s: status = %d, want 200", path, rr.Code)
		}
		if !strings.Contains(rr.Body.String(), want) {
			t.Errorf("%s: missing %q", path, want)
		}
	}
}

func TestHandlerRedirectsNews(t *testing.T) {
	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/news", nil))

	if rr.Code != http.StatusMovedPermanently {
		t.Fatalf("status = %d, want 301", rr.Code)
	}
	if loc := rr.Header().Get("Location"); loc != "/activities" {
		t.Errorf("Location = %q, want /activities", loc)
	}
}

func TestHandlerNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/nope", nil))

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rr.Code)
	}
}
