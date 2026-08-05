package server

import (
	"strings"
	"testing"
)

func TestSiteURLFromEnvPrecedence(t *testing.T) {
	tests := []struct {
		name          string
		siteURL       string
		productionURL string
		deploymentURL string
		want          string
	}{
		{
			name:          "site url wins",
			siteURL:       "https://nirapod.example.org",
			productionURL: "nirapod-browse.vercel.app",
			deploymentURL: "nirapod-browse-abc123.vercel.app",
			want:          "https://nirapod.example.org",
		},
		{
			name:          "production url preferred over deployment url",
			productionURL: "nirapod-browse.vercel.app",
			deploymentURL: "nirapod-browse-abc123.vercel.app",
			want:          "https://nirapod-browse.vercel.app",
		},
		{
			name:          "deployment url is the last resort",
			deploymentURL: "nirapod-browse-abc123.vercel.app",
			want:          "https://nirapod-browse-abc123.vercel.app",
		},
		{
			name:    "trailing slash trimmed",
			siteURL: "https://nirapod.example.org/",
			want:    "https://nirapod.example.org",
		},
		{
			name: "empty when unset",
			want: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("SITE_URL", tc.siteURL)
			t.Setenv("VERCEL_PROJECT_PRODUCTION_URL", tc.productionURL)
			t.Setenv("VERCEL_URL", tc.deploymentURL)

			if got := SiteURLFromEnv(); got != tc.want {
				t.Errorf("SiteURLFromEnv() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestRobotsTxtIncludesSitemap(t *testing.T) {
	got := RobotsTxt("https://nirapod-browse.vercel.app")
	if want := "Sitemap: https://nirapod-browse.vercel.app/sitemap.xml"; !strings.Contains(got, want) {
		t.Errorf("RobotsTxt() = %q, want it to contain %q", got, want)
	}
}

func TestSitemapXMLCoversEverySEOPage(t *testing.T) {
	got := SitemapXML("https://nirapod-browse.vercel.app")
	for path := range pageSEOMap {
		loc := "<loc>https://nirapod-browse.vercel.app" + path + "</loc>"
		if !strings.Contains(got, loc) {
			t.Errorf("sitemap missing entry for %s", path)
		}
	}
}
