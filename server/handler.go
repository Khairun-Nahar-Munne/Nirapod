package server

import (
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

func SiteURLFromEnv() string {
	if u := strings.TrimSpace(os.Getenv("SITE_URL")); u != "" {
		return strings.TrimRight(u, "/")
	}
	if u := strings.TrimSpace(os.Getenv("VERCEL_URL")); u != "" {
		return "https://" + strings.TrimRight(u, "/")
	}
	return ""
}

func staticFS() fs.FS {
	for _, dir := range []string{"static", "../static"} {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return os.DirFS(dir)
		}
	}
	return os.DirFS("static")
}

func NewHandler() http.Handler {
	root := staticFS()
	mux := http.NewServeMux()

	serveHTML := func(w http.ResponseWriter, r *http.Request, filePath, seoPath string) {
		content, err := fs.ReadFile(root, filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		html := injectSEO(string(content), seoPath, SiteURLFromEnv())
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=300")
		_, _ = w.Write([]byte(html))
	}

	fileServer := http.FileServer(http.FS(root))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path

		if reqPath == "/news" || reqPath == "/news/" || reqPath == "/news.html" {
			http.Redirect(w, r, "/activities", http.StatusMovedPermanently)
			return
		}
		if reqPath == "/index/" || reqPath == "/index.html" {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		if strings.HasSuffix(reqPath, ".html") {
			http.Redirect(w, r, strings.TrimSuffix(reqPath, ".html"), http.StatusMovedPermanently)
			return
		}
		if reqPath != "/" && strings.HasSuffix(reqPath, "/") {
			http.Redirect(w, r, strings.TrimSuffix(reqPath, "/"), http.StatusMovedPermanently)
			return
		}
		if reqPath == "/robots.txt" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=86400")
			_, _ = w.Write([]byte(buildRobotsTxt(SiteURLFromEnv())))
			return
		}
		if reqPath == "/sitemap.xml" {
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			w.Header().Set("Cache-Control", "public, max-age=86400")
			_, _ = w.Write([]byte(buildSitemapXML(SiteURLFromEnv())))
			return
		}
		if reqPath == "/" {
			serveHTML(w, r, path.Join("index", "index.html"), "/")
			return
		}
		if !strings.Contains(path.Base(reqPath), ".") {
			candidate := path.Join(strings.TrimPrefix(reqPath, "/"), "index.html")
			if _, err := fs.Stat(root, candidate); err == nil {
				serveHTML(w, r, candidate, reqPath)
				return
			}
		}

		if strings.HasPrefix(reqPath, "/images/") || strings.HasPrefix(reqPath, "/css/") || strings.HasPrefix(reqPath, "/js/") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		}

		if strings.HasPrefix(reqPath, "/images/") || strings.HasPrefix(reqPath, "/css/") || strings.HasPrefix(reqPath, "/js/") || reqPath == "/favicon.svg" {
			fileServer.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		if content, err := fs.ReadFile(root, "404.html"); err == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write(content)
		}
	})

	return mux
}
