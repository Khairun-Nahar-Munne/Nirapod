package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerServesRobots(t *testing.T) {
	t.Setenv("SITE_URL", "https://nirapod-browse.vercel.app")

	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/api/seo?doc=robots", nil))

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/plain") {
		t.Errorf("Content-Type = %q, want text/plain", ct)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "User-agent: *") {
		t.Errorf("missing user-agent directive: %q", body)
	}
	if want := "Sitemap: https://nirapod-browse.vercel.app/sitemap.xml"; !strings.Contains(body, want) {
		t.Errorf("missing %q in %q", want, body)
	}
}

func TestHandlerServesSitemap(t *testing.T) {
	t.Setenv("SITE_URL", "https://nirapod-browse.vercel.app")

	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/api/seo?doc=sitemap", nil))

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/xml") {
		t.Errorf("Content-Type = %q, want application/xml", ct)
	}
	body := rr.Body.String()
	for _, want := range []string{
		"<urlset",
		"<loc>https://nirapod-browse.vercel.app/</loc>",
		"<loc>https://nirapod-browse.vercel.app/learning/fake-news</loc>",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("missing %q in sitemap", want)
		}
	}
}

func TestHandlerFallsBackToRequestPath(t *testing.T) {
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

func TestHandlerRejectsUnknownDocument(t *testing.T) {
	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodGet, "/api/seo", nil))

	if rr.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rr.Code)
	}
}

func TestHandlerRejectsNonGet(t *testing.T) {
	rr := httptest.NewRecorder()
	Handler(rr, httptest.NewRequest(http.MethodPost, "/api/seo?doc=robots", nil))

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", rr.Code)
	}
	if allow := rr.Header().Get("Allow"); allow == "" {
		t.Error("expected Allow header")
	}
}
