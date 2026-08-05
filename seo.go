package main

import (
	"fmt"
	"html"
	"strings"
)

type pageSEO struct {
	Title       string
	Description string
	Keywords    string
	OGImage     string
	NoIndex     bool
}

var pageSEOMap = map[string]pageSEO{
	"/": {
		Title:       "নিরাপদ (NIRAPOD) — শিক্ষার্থীদের ডিজিটাল নিরাপত্তা প্ল্যাটফর্ম",
		Description: "বাংলাদেশের Class 4–10 শিক্ষার্থী, শিক্ষক ও অভিভাবকদের জন্য বাংলা ভাষায় ডিজিটাল নিরাপত্তা, সাইবার বুলিং, ফেইক নিউজ ও অনলাইন স্ক্যাম সম্পর্কে শেখার ও অভিযোগ করার প্ল্যাটফর্ম।",
		Keywords:    "নিরাপদ, NIRAPOD, ডিজিটাল নিরাপত্তা, cyber safety Bangladesh, শিক্ষার্থী, বাংলা, online safety",
		OGImage:     "/images/students.jpeg",
	},
	"/learning": {
		Title:       "শেখার কোণ — ডিজিটাল নিরাপত্তা মডিউল | নিরাপদ",
		Description: "১৫–২০ মিনিটে শিখুন: ফেইক নিউজ, সাইবার বুলিং, পাসওয়ার্ড, AI ও Deepfake, অনলাইন স্ক্যাম, গেমিং ও সোশ্যাল মিডিয়া নিরাপত্তা — কুইজসহ।",
		Keywords:    "ডিজিটাল নিরাপত্তা শেখা, learning modules, cyber safety quiz, বাংলা, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/ai-deepfake": {
		Title:       "AI ও Deepfake শনাক্তকরণ — শেখার কোণ | নিরাপদ",
		Description: "AI-এর ভালো ও খারাপ ব্যবহার, Deepfake ভিডিও চিনুন, ভুয়া ছবি যাচাই করার চেকলিস্ট ও মিনি কুইজ — বাংলায়।",
		Keywords:    "AI, Deepfake, ভুয়া ছবি, fake video, digital literacy, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/cyberbullying": {
		Title:       "Cyberbullying থেকে নিজেকে রক্ষা — শেখার কোণ | নিরাপদ",
		Description: "অনলাইন অত্যাচার কী, কীভাবে নিজেকে রক্ষা করবেন, কাকে জানাবেন ও কুইজ — শিক্ষার্থীদের জন্য বাংলা গাইড।",
		Keywords:    "cyberbullying, সাইবার বুলিং, online harassment, school safety, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/fake-news": {
		Title:       "Fake News শনাক্তকরণ — শেখার কোণ | নিরাপদ",
		Description: "ভুয়া খবর চিনুন, যাচাই করার সহজ বাংলা পদ্ধতি, THINK নিয়ম ও মিনি কুইজ — সোশ্যাল মিডিয়ায় শেয়ারের আগে।",
		Keywords:    "fake news, ভুয়া খবর, misinformation, fact check, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/password": {
		Title:       "Password ও Account Security — শেখার কোণ | নিরাপদ",
		Description: "শক্তিশালী পাসওয়ার্ড তৈরি, 2FA চালু করা ও অ্যাকাউন্ট নিরাপদ রাখার বাংলা গাইড ও কুইজ।",
		Keywords:    "password security, পাসওয়ার্ড, 2FA, account safety, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/online-scam": {
		Title:       "Online Scam ও Phishing — শেখার কোণ | নিরাপদ",
		Description: "OTP প্রতারণা, ফিশিং লিংক, ভুয়া অফার চিনুন ও অনলাইন স্ক্যাম থেকে বাঁচার বাংলা টিপস ও কুইজ।",
		Keywords:    "online scam, phishing, OTP fraud, প্রতারণা, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/social-media": {
		Title:       "Social Media Safety — শেখার কোণ | নিরাপদ",
		Description: "Facebook, Instagram, TikTok-এ প্রাইভেসি সেটিং, ব্লক ও রিপোর্ট করার বাংলা গাইড ও কুইজ।",
		Keywords:    "social media safety, প্রাইভেসি, Facebook, Instagram, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/gaming": {
		Title:       "Online Gaming Safety — শেখার কোণ | নিরাপদ",
		Description: "গেম চ্যাট, অজানা লিংক ও অ্যাকাউন্ট নিরাপদ রাখার উপায় — শিক্ষার্থীদের জন্য বাংলা গাইড ও কুইজ।",
		Keywords:    "gaming safety, online games, game chat, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/digital-citizenship": {
		Title:       "Digital Citizenship — শেখার কোণ | নিরাপদ",
		Description: "অনলাইনে ভদ্র, দায়িত্বশীল ও নিরাপদ ডিজিটাল নাগরিক হওয়ার বাংলা গাইড ও কুইজ।",
		Keywords:    "digital citizenship, ডিজিটাল নাগরিক, online etiquette, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/parent": {
		Title:       "Parent Corner — অভিভাবকদের ডিজিটাল নিরাপত্তা | নিরাপদ",
		Description: "অভিভাবকদের জন্য সন্তানের স্ক্রিনটাইম, সোশ্যাল মিডিয়া ও অনলাইন ঝুঁকি নিয়ন্ত্রণের বাংলা গাইড।",
		Keywords:    "parent guide, অভিভাবক, child online safety, নিরাপদ",
		OGImage:     "/images/cyber-learning.jpeg",
	},
	"/learning/teacher": {
		Title:       "Teacher Corner — শ্রেণিকক্ষে ডিজিটাল নিরাপত্তা | নিরাপদ",
		Description: "শিক্ষকদের জন্য ক্লাসরুম সচেতনতা কার্যক্রম, পোস্টার ও ডিজিটাল নিরাপত্তা শেখানোর বাংলা রিসোর্স।",
		Keywords:    "teacher resources, শিক্ষক, classroom activity, digital safety, নিরাপদ",
		OGImage:     "/images/teaching.jpeg",
	},
	"/learning/quiz": {
		Title:       "Quiz & Certificate — Digital Safety Champion | নিরাপদ",
		Description: "ডিজিটাল নিরাপত্তা কুইজ দিন, ৮০% স্কোরে Digital Safety Champion সার্টিফিকেট পান — বাংলায়।",
		Keywords:    "digital safety quiz, certificate, কুইজ, নিরাপদ",
		OGImage:     "/images/students-exam.jpeg",
	},
	"/safe": {
		Title:       "নিরাপদ থাকুন — প্ল্যাটফর্ম সেফটি গাইড | নিরাপদ",
		Description: "Facebook, WhatsApp, YouTube, TikTok ও Instagram-এ প্রাইভেসি ও রিপোর্ট সেটিংস — বাংলায় ধাপে ধাপে।",
		Keywords:    "platform safety, privacy settings, Facebook, WhatsApp, নিরাপদ",
		OGImage:     "/images/students.jpeg",
	},
	"/report": {
		Title:       "অভিযোগ করুন — গোপনীয় অনলাইন অভিযোগ | নিরাপদ",
		Description: "সাইবার বুলিং, অনলাইন হয়রানি বা ডিজিটাল নিরাপত্তা সমস্যা নিরাপদে ও গোপনীয়ভাবে Google Form-এ জানান।",
		Keywords:    "অভিযোগ, report cyberbullying, online complaint, school safety, নিরাপদ",
		OGImage:     "/images/students.jpeg",
	},
	"/dashboard": {
		Title:       "ড্যাশবোর্ড — অভিযোগ পরিসংখ্যান | নিরাপদ",
		Description: "নিরাপদ প্রকল্পের ডেমো ড্যাশবোর্ড — অভিযোগের পরিসংখ্যান ও সচেতনতা কার্যক্রমের ওভারভিউ।",
		Keywords:    "dashboard, statistics, demo, নিরাপদ",
		OGImage:     "/images/students.jpeg",
	},
	"/team": {
		Title:       "নিরাপদ দল — শিক্ষক ও শিক্ষার্থী নেতা | নিরাপদ",
		Description: "নিরাপদ প্রকল্পের শিক্ষক উপদেষ্টা, শিক্ষার্থী নেতা ও ডিজিটাল সিকিউরিটি টিম পরিচিতি।",
		Keywords:    "team, নিরাপদ দল, student leaders, digital safety project",
		OGImage:     "/images/students.jpeg",
	},
	"/resources": {
		Title:       "রিসোর্স — পোস্টার, PDF ও ইনফোগ্রাফিক | নিরাপদ",
		Description: "ডিজিটাল নিরাপত্তা পোস্টার, লিফলেট, PDF ও ইনফোগ্রাফিক ডাউনলোড — স্কুল ও ক্লাসরুমে ব্যবহারের জন্য।",
		Keywords:    "resources, poster, PDF, infographic, digital safety, নিরাপদ",
		OGImage:     "/images/session.jpeg",
	},
	"/activities": {
		Title:       "কার্যক্রম ও গ্যালারি — সচেতনতা ক্যাম্পেইন | নিরাপদ",
		Description: "স্কুলভিত্তিক ডিজিটাল সেফটি ওয়ার্কশপ, সচেতনতা সেশন, কুইজ ও শিক্ষার্থী নেতৃত্বের কার্যক্রমের গ্যালারি।",
		Keywords:    "activities, workshop, awareness campaign, gallery, নিরাপদ",
		OGImage:     "/images/session.jpeg",
	},
	"/contact": {
		Title:       "যোগাযোগ — স্কুল ও ডিজিটাল সিকিউরিটি টিম | নিরাপদ",
		Description: "নিরাপদ প্রকল্পে যোগাযোগ করুন — স্কুল, শিক্ষক ও ডিজিটাল সিকিউরিটি টিমের তথ্য।",
		Keywords:    "contact, যোগাযোগ, school, digital safety team, নিরাপদ",
		OGImage:     "/images/students.jpeg",
	},
}

var sitemapPaths = []string{
	"/",
	"/learning",
	"/learning/ai-deepfake",
	"/learning/cyberbullying",
	"/learning/fake-news",
	"/learning/password",
	"/learning/online-scam",
	"/learning/social-media",
	"/learning/gaming",
	"/learning/digital-citizenship",
	"/learning/parent",
	"/learning/teacher",
	"/learning/quiz",
	"/safe",
	"/report",
	"/dashboard",
	"/team",
	"/resources",
	"/activities",
	"/contact",
}

func seoForPath(path string) (pageSEO, bool) {
	path = strings.TrimSuffix(path, "/")
	if path == "" {
		path = "/"
	}
	seo, ok := pageSEOMap[path]
	return seo, ok
}

func injectSEO(htmlContent, path, siteURL string) string {
	if strings.Contains(htmlContent, `name="nirapod-seo"`) {
		return htmlContent
	}

	seo, ok := seoForPath(path)
	if !ok {
		return htmlContent
	}

	htmlContent = replaceTitle(htmlContent, seo.Title)
	block := buildSEOBlock(seo, path, siteURL)
	marker := `<meta name="viewport" content="width=device-width, initial-scale=1" />`
	if idx := strings.Index(htmlContent, marker); idx >= 0 {
		insertPos := idx + len(marker)
		return htmlContent[:insertPos] + "\n" + block + htmlContent[insertPos:]
	}

	headMarker := "<head>"
	if idx := strings.Index(htmlContent, headMarker); idx >= 0 {
		insertPos := idx + len(headMarker)
		return htmlContent[:insertPos] + "\n" + block + htmlContent[insertPos:]
	}

	return htmlContent
}

func replaceTitle(htmlContent, title string) string {
	start := strings.Index(htmlContent, "<title>")
	end := strings.Index(htmlContent, "</title>")
	if start < 0 || end < 0 || end <= start {
		return htmlContent
	}
	escaped := html.EscapeString(title)
	return htmlContent[:start] + "<title>" + escaped + "</title>" + htmlContent[end+len("</title>"):]
}

func buildSEOBlock(seo pageSEO, path, siteURL string) string {
	canonical := canonicalURL(siteURL, path)
	ogImage := absoluteURL(siteURL, seo.OGImage)
	siteName := "নিরাপদ (NIRAPOD)"
	title := html.EscapeString(seo.Title)
	desc := html.EscapeString(seo.Description)
	keywords := html.EscapeString(seo.Keywords)
	canonicalEsc := html.EscapeString(canonical)
	ogImageEsc := html.EscapeString(ogImage)

	robots := "index, follow, max-image-preview:large, max-snippet:-1, max-video-preview:-1"
	if seo.NoIndex {
		robots = "noindex, nofollow"
	}

	var b strings.Builder
	b.WriteString(`  <meta name="nirapod-seo" content="1" />` + "\n")
	b.WriteString(`  <meta name="description" content="` + desc + `" />` + "\n")
	b.WriteString(`  <meta name="keywords" content="` + keywords + `" />` + "\n")
	b.WriteString(`  <meta name="author" content="` + html.EscapeString(siteName) + `" />` + "\n")
	b.WriteString(`  <meta name="robots" content="` + robots + `" />` + "\n")
	b.WriteString(`  <link rel="canonical" href="` + canonicalEsc + `" />` + "\n")
	b.WriteString(`  <link rel="alternate" hreflang="bn" href="` + canonicalEsc + `" />` + "\n")
	b.WriteString(`  <meta name="theme-color" content="#0d4f8b" />` + "\n")
	b.WriteString(`  <meta property="og:type" content="website" />` + "\n")
	b.WriteString(`  <meta property="og:site_name" content="` + html.EscapeString(siteName) + `" />` + "\n")
	b.WriteString(`  <meta property="og:locale" content="bn_BD" />` + "\n")
	b.WriteString(`  <meta property="og:title" content="` + title + `" />` + "\n")
	b.WriteString(`  <meta property="og:description" content="` + desc + `" />` + "\n")
	b.WriteString(`  <meta property="og:url" content="` + canonicalEsc + `" />` + "\n")
	b.WriteString(`  <meta property="og:image" content="` + ogImageEsc + `" />` + "\n")
	b.WriteString(`  <meta name="twitter:card" content="summary_large_image" />` + "\n")
	b.WriteString(`  <meta name="twitter:title" content="` + title + `" />` + "\n")
	b.WriteString(`  <meta name="twitter:description" content="` + desc + `" />` + "\n")
	b.WriteString(`  <meta name="twitter:image" content="` + ogImageEsc + `" />` + "\n")

	if path == "/" {
		b.WriteString(buildHomeJSONLD(siteURL, seo))
	} else if strings.HasPrefix(path, "/learning/") && path != "/learning/quiz" {
		b.WriteString(buildLearningModuleJSONLD(siteURL, path, seo))
	}

	return b.String()
}

func buildHomeJSONLD(siteURL string, seo pageSEO) string {
	url := absoluteURL(siteURL, "/")
	json := fmt.Sprintf(`  <script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "WebSite",
  "name": "নিরাপদ (NIRAPOD)",
  "alternateName": "NIRAPOD",
  "url": %q,
  "description": %q,
  "inLanguage": "bn-BD",
  "audience": {
    "@type": "EducationalAudience",
    "educationalRole": "student"
  }
}
</script>
`, url, seo.Description)
	return json
}

func buildLearningModuleJSONLD(siteURL, path string, seo pageSEO) string {
	url := absoluteURL(siteURL, path)
	json := fmt.Sprintf(`  <script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "LearningResource",
  "name": %q,
  "description": %q,
  "url": %q,
  "inLanguage": "bn-BD",
  "learningResourceType": "lesson",
  "isPartOf": {
    "@type": "WebSite",
    "name": "নিরাপদ (NIRAPOD)",
    "url": %q
  }
}
</script>
`, seo.Title, seo.Description, url, absoluteURL(siteURL, "/"))
	return json
}

func canonicalURL(siteURL, path string) string {
	if siteURL == "" {
		return path
	}
	return strings.TrimRight(siteURL, "/") + path
}

func absoluteURL(siteURL, path string) string {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	if siteURL == "" {
		return path
	}
	return strings.TrimRight(siteURL, "/") + path
}

func buildRobotsTxt(siteURL string) string {
	var b strings.Builder
	b.WriteString("User-agent: *\n")
	b.WriteString("Allow: /\n")
	b.WriteString("Disallow: /404\n")
	if siteURL != "" {
		b.WriteString("\nSitemap: ")
		b.WriteString(strings.TrimRight(siteURL, "/"))
		b.WriteString("/sitemap.xml\n")
	}
	return b.String()
}

func buildSitemapXML(siteURL string) string {
	if siteURL == "" {
		siteURL = "http://localhost:8080"
	}
	base := strings.TrimRight(siteURL, "/")
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	for _, path := range sitemapPaths {
		priority := "0.7"
		changefreq := "monthly"
		if path == "/" {
			priority = "1.0"
			changefreq = "weekly"
		} else if path == "/learning" {
			priority = "0.9"
			changefreq = "weekly"
		} else if strings.HasPrefix(path, "/learning/") {
			priority = "0.8"
		}
		b.WriteString("  <url>\n")
		b.WriteString("    <loc>")
		b.WriteString(html.EscapeString(base + path))
		b.WriteString("</loc>\n")
		b.WriteString("    <changefreq>")
		b.WriteString(changefreq)
		b.WriteString("</changefreq>\n")
		b.WriteString("    <priority>")
		b.WriteString(priority)
		b.WriteString("</priority>\n")
		b.WriteString("  </url>\n")
	}
	b.WriteString("</urlset>\n")
	return b.String()
}
