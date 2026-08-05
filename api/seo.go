// Package handler exposes the generated SEO documents as a Vercel Go function.
// The rest of the site is plain static output from public/, so only the files
// that must be rendered from Go live here.
package handler

import (
	"net/http"
	"strings"

	"nirapod/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	siteURL := server.SiteURLFromEnv()

	var body, contentType string
	switch requestedDoc(r) {
	case "robots":
		body, contentType = server.RobotsTxt(siteURL), "text/plain; charset=utf-8"
	case "sitemap":
		body, contentType = server.SitemapXML(siteURL), "application/xml; charset=utf-8"
	default:
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=86400, stale-while-revalidate=604800")
	if r.Method == http.MethodHead {
		return
	}
	_, _ = w.Write([]byte(body))
}

// requestedDoc reads the document from the rewrite's query parameter, falling
// back to the request path so the function still works when invoked directly.
func requestedDoc(r *http.Request) string {
	if doc := r.URL.Query().Get("doc"); doc != "" {
		return doc
	}
	switch path := strings.TrimSuffix(r.URL.Path, "/"); {
	case strings.HasSuffix(path, "robots.txt"):
		return "robots"
	case strings.HasSuffix(path, "sitemap.xml"):
		return "sitemap"
	}
	return ""
}
