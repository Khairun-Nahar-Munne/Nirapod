(function () {
  const cfg = window.NIRAPOD || {};
  const formUrl = cfg.googleFormUrl || "#";

  const NAV = [
    { href: "/", label: "হোম" },
    { href: "/learning", label: "শেখার কোণ" },
    { href: "/safe", label: "নিরাপদ থাকুন" },
    { href: "/report", label: "অভিযোগ" },
    { href: "/dashboard", label: "ড্যাশবোর্ড" },
    { href: "/team", label: "নিরাপদ দল" },
    { href: "/resources", label: "রিসোর্স" },
    { href: "/activities", label: "কার্যক্রম" },
    { href: "/contact", label: "যোগাযোগ" }
  ];

  function currentPath() {
    return window.location.pathname.replace(/\\/g, "/");
  }

  function isActive(href) {
    const path = currentPath();
    if (href === "/") {
      return path === "/";
    }
    if (href === "/learning") {
      return path.startsWith("/learning");
    }
    return path === href;
  }

  function openComplaint(e) {
    if (e) e.preventDefault();
    if (!formUrl || formUrl.includes("YOUR_FORM_ID")) {
      alert(
        "Google Form লিংক এখনো সেট করা হয়নি। static/js/config.js ফাইলে googleFormUrl আপডেট করুন।"
      );
      return;
    }
    window.open(formUrl, "_blank", "noopener,noreferrer");
  }

  function buildHeader() {
    const navLinks = NAV.map(
      (item) =>
        `<a href="${item.href}" class="${isActive(item.href) ? "active" : ""}">${item.label}</a>`
    ).join("");

    return `
<header class="site-header">
  <div class="container header-inner">
    <a class="brand" href="/">
      <span class="brand-mark" aria-hidden="true">🛡️</span>
      <span>${cfg.siteName || "নিরাপদ"}</span>
    </a>
    <button class="nav-toggle" type="button" aria-label="মেনু" aria-expanded="false">☰ মেনু</button>
    <div class="nav-wrap" id="navWrap">
      <nav class="site-nav" aria-label="প্রধান মেনু">${navLinks}</nav>
      <button type="button" class="btn btn-complain" data-complain>🚨 অভিযোগ করুন</button>
    </div>
  </div>
</header>`;
  }

  function buildFooter() {
    return `
<footer class="site-footer">
  <div class="container footer-grid">
    <div>
      <h3>🛡️ ${cfg.siteName || "নিরাপদ"}</h3>
      <p>বাংলাদেশের শিক্ষার্থীদের জন্য ডিজিটাল নিরাপত্তা, সচেতনতা ও অভিযোগ ব্যবস্থাপনা প্ল্যাটফর্ম।</p>
      <p class="footer-motto">${cfg.motto || "সচেতন হই, নিরাপদ থাকি"}</p>
    </div>
    <div>
      <h3>দ্রুত লিংক</h3>
      <ul>
        <li><a href="/learning">শেখার কোণ</a></li>
        <li><a href="/report">অভিযোগ করুন</a></li>
        <li><a href="/dashboard">ড্যাশবোর্ড</a></li>
        <li><a href="/resources">রিসোর্স</a></li>
      </ul>
    </div>
    <div>
      <h3>যোগাযোগ</h3>
      <ul>
        <li><a href="/contact">স্কুল ও শিক্ষক</a></li>
      </ul>
    </div>
  </div>
  <div class="container footer-bottom">
    © ${new Date().getFullYear()} নিরাপদ (NIRAPOD) · Capstone Project
  </div>
</footer>`;
  }

  function injectChrome() {
    const headerMount = document.getElementById("site-header");
    const footerMount = document.getElementById("site-footer");
    if (headerMount) headerMount.outerHTML = buildHeader();
    if (footerMount) footerMount.outerHTML = buildFooter();

    document.querySelectorAll("[data-complain]").forEach((btn) => {
      btn.addEventListener("click", openComplaint);
    });

    const toggle = document.querySelector(".nav-toggle");
    const wrap = document.getElementById("navWrap");
    if (toggle && wrap) {
      toggle.addEventListener("click", () => {
        const open = wrap.classList.toggle("open");
        toggle.setAttribute("aria-expanded", open ? "true" : "false");
      });
    }
  }

  function initLearningSearch() {
    const input = document.getElementById("learningSearch");
    if (!input) return;
    const cards = document.querySelectorAll("[data-category]");
    input.addEventListener("input", () => {
      const q = input.value.trim().toLowerCase();
      cards.forEach((card) => {
        const hay = (card.getAttribute("data-category") || "").toLowerCase();
        card.classList.toggle("hidden", q && !hay.includes(q));
      });
    });
  }

  function initQuiz(root) {
    const form = root.querySelector("[data-quiz]");
    if (!form) return;
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      const questions = form.querySelectorAll("[data-answer]");
      let score = 0;
      questions.forEach((q) => {
        const correct = q.getAttribute("data-answer");
        const selected = q.querySelector("input:checked");
        if (selected && selected.value === correct) score += 1;
      });
      const total = questions.length;
      const out = form.querySelector(".quiz-result");
      if (out) {
        const pct = Math.round((score / total) * 100);
        out.textContent =
          pct >= 80
            ? `🎉 দুর্দান্ত! ${score}/${total} (${pct}%) — Digital Safety Champion!`
            : `আপনার স্কোর: ${score}/${total} (${pct}%)। আবার চেষ্টা করুন!`;
      }
    });
  }

  function initPasswordChecker() {
    const input = document.getElementById("pwInput");
    const bar = document.getElementById("pwBar");
    const label = document.getElementById("pwLabel");
    if (!input || !bar || !label) return;

    input.addEventListener("input", () => {
      const v = input.value;
      let score = 0;
      if (v.length >= 8) score += 1;
      if (v.length >= 12) score += 1;
      if (/[A-Z]/.test(v) && /[a-z]/.test(v)) score += 1;
      if (/\d/.test(v)) score += 1;
      if (/[^A-Za-z0-9]/.test(v)) score += 1;

      const widths = ["0%", "20%", "40%", "60%", "80%", "100%"];
      const colors = ["#c62828", "#ef6c00", "#f9a825", "#7cb342", "#2e7d32", "#1b5e20"];
      const texts = ["খুব দুর্বল", "দুর্বল", "মাঝারি", "ভালো", "শক্তিশালী", "অত্যন্ত শক্তিশালী"];
      bar.style.width = widths[score];
      bar.style.background = colors[score];
      label.textContent = v ? texts[score] : "পাসওয়ার্ড লিখুন (সংরক্ষণ করা হয় না)";
    });
  }

  function initDashboard() {
    const chart = document.getElementById("monthlyBars");
    if (!chart) return;
    const data = [
      { label: "জুলাই", value: 12 },
      { label: "আগস্ট", value: 18 }
    ];
    const max = Math.max(...data.map((d) => d.value));
    chart.innerHTML = data
      .map(
        (d) =>
          `<div class="bar" style="height:${Math.round((d.value / max) * 100)}%"><span>${d.label}</span></div>`
      )
      .join("");
  }

  document.addEventListener("DOMContentLoaded", () => {
    injectChrome();
    initLearningSearch();
    document.querySelectorAll(".quiz-block").forEach(initQuiz);
    initPasswordChecker();
    initDashboard();
  });

  window.NIRAPOD_OPEN_COMPLAINT = openComplaint;
})();
