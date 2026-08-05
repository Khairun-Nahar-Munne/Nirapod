package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestStaticIndexExists(t *testing.T) {
	path := filepath.Join("public", "index.html")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s: %v", path, err)
	}
}

func TestAllImagesPresent(t *testing.T) {
	needed := []string{
		"students.jpeg",
		"teaching.jpeg",
		"workshop.jpeg",
		"students-exam.jpeg",
		"cyber-learning.jpeg",
		"security-learning.jpeg",
		"www-learning.jpeg",
	}
	for _, name := range needed {
		path := filepath.Join("public", "images", name)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("missing image %s: %v", name, err)
		}
	}
}

func TestFileServerServesRouteIndex(t *testing.T) {
	fs := http.FileServer(http.Dir("public"))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	fs.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct == "" {
		t.Fatal("expected Content-Type header")
	}
}
