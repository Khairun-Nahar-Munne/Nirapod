# নিরাপদ (NIRAPOD) — Capstone MVP

Bangla-first digital safety website for Class 4–10 students: learning hub, awareness, Google Form complaints, and a demo dashboard.

## Stack

- **Go** — static file server (`main.go`)
- **HTML / CSS / vanilla JS** — pages in `static/`

## Run

```bash
cd Nirapod
go run .
```

Open [http://localhost:8080](http://localhost:8080)

Optional port:

```bash
set PORT=3000
go run .
```

## Google Form (অভিযোগ করুন)

Edit `static/js/config.js`:

```js
window.NIRAPOD = {
  siteName: "নিরাপদ",
  motto: "সচেতন হই, নিরাপদ থাকি",
  googleFormUrl: "https://docs.google.com/forms/d/e/YOUR_REAL_FORM_ID/viewform"
};
```

The red **অভিযোগ করুন** button in the header (and on Report / related pages) opens this URL in a new tab. Responses are collected in Google Forms.

## Pages

| Path | Description |
|------|-------------|
| `/` | Home (hero photo) |
| `/learning/` | Learning hub + search |
| `/learning/<module>/` | 12 safety modules + quizzes |
| `/safe/` | Platform safety tips |
| `/report/` | Complaint CTA → Google Form |
| `/dashboard/` | Demo stats |
| `/team/` | Team |
| `/resources/` | Resources |
| `/news/` | Activities gallery |
| `/contact/` | Contact |

## Images

Project photos live in `static/images/` (copied from `Images/` with kebab-case names). All seven photos are used across Home, Learning, Team, News, and Resources.
