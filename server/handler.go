package server

import (
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

// SiteURLFromEnv resolves the absolute site origin used for canonical links,
// sitemap entries and Open Graph URLs. VERCEL_PROJECT_PRODUCTION_URL is
// preferred over VERCEL_URL because the latter is unique per deployment and
// would leak build-specific hostnames into indexable metadata.
func SiteURLFromEnv() string {
	for _, key := range []string{"SITE_URL", "VERCEL_PROJECT_PRODUCTION_URL", "VERCEL_URL"} {
		u := strings.TrimSpace(os.Getenv(key))
		if u == "" {
			continue
		}
		u = strings.TrimRight(u, "/")
		if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
			return u
		}
		return "https://" + u
	}
	return ""
}

// RobotsTxt renders robots.txt for the given site origin.
func RobotsTxt(siteURL string) string { return buildRobotsTxt(siteURL) }

// SitemapXML renders sitemap.xml for the given site origin.
func SitemapXML(siteURL string) string { return buildSitemapXML(siteURL) }

func staticFS() fs.FS {
	for _, dir := range []string{"public", "../public"} {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return os.DirFS(dir)
		}
	}
	return os.DirFS("public")
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
			serveHTML(w, r, "index.html", "/")
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
