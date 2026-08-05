package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func siteURLFromEnv() string {
	if u := strings.TrimSpace(os.Getenv("SITE_URL")); u != "" {
		return strings.TrimRight(u, "/")
	}
	return ""
}

func serveHTML(w http.ResponseWriter, r *http.Request, filePath, seoPath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	html := injectSEO(string(content), seoPath, siteURLFromEnv())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=300")
	_, _ = w.Write([]byte(html))
}

func main() {
	staticDir := "static"
	if _, err := os.Stat(staticDir); err != nil {
		log.Fatalf("static directory not found: %v", err)
	}

	fs := http.FileServer(http.Dir(staticDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Keep clean, extensionless URLs canonical and redirect legacy paths.
		if path == "/news" || path == "/news/" || path == "/news.html" {
			http.Redirect(w, r, "/activities", http.StatusMovedPermanently)
			return
		}
		if path == "/index/" || path == "/index.html" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		if strings.HasSuffix(path, ".html") {
			http.Redirect(w, r, strings.TrimSuffix(path, ".html"), http.StatusMovedPermanently)
			return
		}
		if path != "/" && strings.HasSuffix(path, "/") {
			http.Redirect(w, r, strings.TrimSuffix(path, "/"), http.StatusMovedPermanently)
			return
		}
		if path == "/robots.txt" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=86400")
			_, _ = w.Write([]byte(buildRobotsTxt(siteURLFromEnv())))
			return
		}
		if path == "/sitemap.xml" {
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=86400")
			_, _ = w.Write([]byte(buildSitemapXML(siteURLFromEnv())))
			return
		}
		if path == "/" {
			serveHTML(w, r, filepath.Join(staticDir, "index", "index.html"), "/")
			return
		}
		if !strings.Contains(filepath.Base(path), ".") {
			candidate := filepath.Join(staticDir, strings.TrimPrefix(path, "/"), "index.html")
			if _, err := os.Stat(candidate); err == nil {
				serveHTML(w, r, candidate, path)
				return
			}
		}

		// Set caching for static assets
		if strings.HasPrefix(path, "/images/") || strings.HasPrefix(path, "/css/") || strings.HasPrefix(path, "/js/") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		}

		if strings.HasPrefix(path, "/images/") || strings.HasPrefix(path, "/css/") || strings.HasPrefix(path, "/js/") || path == "/favicon.svg" {
			fs.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, filepath.Join(staticDir, "404.html"))
	})

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("নিরাপদ (NIRAPOD) running at http://localhost%s", addr)
	log.Fatal(server.ListenAndServe())
}
