<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>VulnScan — Real-Time Security Audit Pipeline</title>
<link href="https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&family=Plus+Jakarta+Sans:wght@300;400;500;700&display=swap" rel="stylesheet">
<style>
:root {
  --bg:       #f7faff;
  --bg2:      #eef4ff;
  --bg3:      #ffffff;
  --surface:  #ffffff;
  --lblue:    #b8d9ff;
  --lblue2:   #d6eaff;
  --lblue3:   #87bfff;
  --lime:     #d4f542;
  --lime2:    #bfe320;
  --lime3:    #e8fb7a;
  --navy:     #0d2240;
  --navy2:    #1a3a5c;
  --text:     #0d2240;
  --muted:    #4a6a8a;
  --muted2:   #7a9ab8;
  --border:   rgba(13,34,64,0.1);
  --border2:  rgba(13,34,64,0.2);
  --mono:     'Space Mono', monospace;
  --sans:     'Plus Jakarta Sans', sans-serif;
}

*{margin:0;padding:0;box-sizing:border-box;}

body {
  background: var(--bg);
  color: var(--text);
  font-family: var(--sans);
  font-size: 15px;
  line-height: 1.7;
  overflow-x: hidden;
}

/* ── DOT GRID BACKGROUND ── */
body::before {
  content:'';
  position:fixed;
  inset:0;
  background-image: radial-gradient(circle, rgba(13,34,64,0.07) 1px, transparent 1px);
  background-size: 28px 28px;
  pointer-events:none;
  z-index:0;
}

.container { max-width:860px; margin:0 auto; padding:0 2rem; position:relative; z-index:1; }

/* ═══════════════ HERO ═══════════════ */
.hero {
  min-height: 88vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 5rem 2rem 4rem;
  position: relative;
  overflow: hidden;
}

/* animated blobs */
.blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(70px);
  opacity: 0.45;
  animation: drift 10s ease-in-out infinite alternate;
  pointer-events: none;
}
.blob1 { width:420px;height:340px; background:var(--lblue); top:-80px; left:-100px; animation-duration:11s; }
.blob2 { width:360px;height:300px; background:var(--lime3); top:-40px; right:-80px; opacity:0.35; animation-duration:14s; animation-delay:2s; }
.blob3 { width:280px;height:240px; background:var(--lblue2); bottom:0; left:50%; transform:translateX(-50%); opacity:0.4; animation-duration:9s; animation-delay:1s; }

@keyframes drift {
  from { transform: translate(0,0) scale(1); }
  to   { transform: translate(20px, 15px) scale(1.06); }
}
.blob3 { animation-name: drift3; }
@keyframes drift3 {
  from { transform: translateX(-50%) scale(1); }
  to   { transform: translateX(calc(-50% + 15px)) scale(1.05); }
}

/* floating particles */
.particles { position:absolute; inset:0; pointer-events:none; overflow:hidden; }
.p {
  position:absolute;
  width:4px;height:4px;
  border-radius:50%;
  background:var(--navy);
  opacity:0.12;
  animation: floatUp linear infinite;
}
.p:nth-child(1){left:10%;animation-duration:8s;animation-delay:0s;width:3px;height:3px;}
.p:nth-child(2){left:25%;animation-duration:11s;animation-delay:2s;background:var(--lime2);opacity:0.3;}
.p:nth-child(3){left:45%;animation-duration:9s;animation-delay:1s;}
.p:nth-child(4){left:65%;animation-duration:12s;animation-delay:3s;background:var(--lblue3);opacity:0.4;}
.p:nth-child(5){left:80%;animation-duration:7s;animation-delay:0.5s;width:5px;height:5px;}
.p:nth-child(6){left:90%;animation-duration:10s;animation-delay:4s;background:var(--lime2);opacity:0.25;}

@keyframes floatUp {
  0%   { bottom:-20px; opacity:0; }
  10%  { opacity:0.2; }
  90%  { opacity:0.15; }
  100% { bottom:110%; opacity:0; }
}

/* hero content */
.hero-inner { position:relative; z-index:2; max-width:700px; }

.status-badge {
  display:inline-flex; align-items:center; gap:7px;
  background:#ffffff;
  border:1.5px solid var(--lblue3);
  border-radius:100px;
  padding:5px 16px;
  font-family:var(--mono);
  font-size:10px;
  letter-spacing:0.12em;
  text-transform:uppercase;
  color:var(--navy2);
  margin-bottom:2rem;
  animation: popIn 0.6s cubic-bezier(0.34,1.56,0.64,1) both;
}
.status-dot {
  width:7px;height:7px;border-radius:50%;
  background:var(--lime2);
  animation: pulse 2s infinite;
}
@keyframes pulse{0%,100%{box-shadow:0 0 0 0 rgba(180,210,30,0.5);}50%{box-shadow:0 0 0 5px rgba(180,210,30,0);}}

.hero h1 {
  font-family:var(--mono);
  font-size:clamp(2rem,6vw,3.6rem);
  font-weight:700;
  line-height:1.1;
  letter-spacing:-0.03em;
  margin-bottom:1.2rem;
  animation: slideUp 0.7s 0.15s ease both;
}

.accent-lime { 
  color:var(--lime2);
  -webkit-text-stroke:1px var(--lime2);
  position:relative;
}
.accent-blue { color:var(--lblue3); }

.hero-desc {
  font-size:1.05rem;
  color:var(--muted);
  font-weight:300;
  max-width:520px;
  margin:0 auto 2.4rem;
  animation: slideUp 0.7s 0.25s ease both;
}

.hero-pills {
  display:flex; flex-wrap:wrap; gap:8px; justify-content:center;
  animation: slideUp 0.7s 0.35s ease both;
  margin-bottom:2.5rem;
}

.pill {
  display:inline-flex; align-items:center; gap:5px;
  padding:5px 13px;
  border-radius:7px;
  font-family:var(--mono);
  font-size:11px;
  font-weight:700;
  letter-spacing:0.04em;
  border:1.5px solid;
  transition:transform 0.15s, box-shadow 0.15s;
  background:#fff;
}
.pill:hover{transform:translateY(-2px);box-shadow:0 4px 12px rgba(13,34,64,0.1);}

.pill-go   {color:#00add8;border-color:#b8e8f5;}
.pill-rust {color:#c0622a;border-color:#f5d0b8;}
.pill-py   {color:#2b6cb0;border-color:#b8d4f5;}
.pill-grpc {color:#6b4fa0;border-color:#d8c8f5;}
.pill-q    {color:#1a7a4a;border-color:#b8f0d4;}
.pill-sb   {color:#1a8a5a;border-color:#b8f0d4;}
.pill-ai   {color:#7a5a00;border-color:#f5e8b0;}

/* scroll indicator */
.scroll-hint {
  animation: slideUp 0.7s 0.5s ease both;
  display:flex; flex-direction:column; align-items:center; gap:6px;
  font-family:var(--mono); font-size:10px; letter-spacing:0.1em;
  text-transform:uppercase; color:var(--muted2);
}
.scroll-arrow {
  width:20px;height:20px;border-right:2px solid var(--muted2);border-bottom:2px solid var(--muted2);
  transform:rotate(45deg);
  animation:bounce 1.5s infinite;
}
@keyframes bounce{0%,100%{transform:rotate(45deg) translateY(0);}50%{transform:rotate(45deg) translateY(4px);}}

/* ── ANIMATIONS ── */
@keyframes slideUp  {from{opacity:0;transform:translateY(22px);}to{opacity:1;transform:none;}}
@keyframes popIn    {from{opacity:0;transform:scale(0.85);}to{opacity:1;transform:scale(1);}}

.fade-in{opacity:0;transform:translateY(14px);transition:opacity 0.55s ease,transform 0.55s ease;}
.fade-in.vis{opacity:1;transform:none;}

/* ═══════════════ LAYOUT ═══════════════ */
hr.div {border:none;border-top:1.5px solid var(--border);margin:3.5rem 0;}
.section{margin:4rem 0;}
.section-tag{font-family:var(--mono);font-size:10px;letter-spacing:0.15em;text-transform:uppercase;color:var(--muted2);margin-bottom:0.4rem;}
h2{font-family:var(--mono);font-size:1.45rem;font-weight:700;margin-bottom:1.1rem;color:var(--navy);}
h3{font-family:var(--mono);font-size:0.95rem;font-weight:700;color:var(--navy2);margin-bottom:0.5rem;}
p{color:var(--muted);margin-bottom:0.8rem;}

/* ═══════════════ STATS ═══════════════ */
.stats{display:grid;grid-template-columns:repeat(auto-fit,minmax(140px,1fr));gap:1px;background:var(--border);border:1.5px solid var(--border);border-radius:14px;overflow:hidden;margin:2rem 0;}
.stat{background:#fff;padding:1.6rem 1rem;text-align:center;}
.stat-num{font-family:var(--mono);font-size:2rem;font-weight:700;color:var(--navy);display:block;line-height:1;margin-bottom:0.3rem;}
.stat-num.lime{color:var(--lime2);}
.stat-num.blue{color:var(--lblue3);}
.stat-label{font-size:11px;color:var(--muted2);line-height:1.4;}

/* ═══════════════ PIPELINE ═══════════════ */
.pipeline {
  background:#fff;
  border:1.5px solid var(--border);
  border-radius:14px;
  padding:2rem;
  font-family:var(--mono);
  font-size:12px;
  color:var(--navy);
  line-height:2;
  overflow-x:auto;
}
.pipeline::before{content:'// SYSTEM ARCHITECTURE';display:block;color:var(--muted2);font-size:10px;letter-spacing:0.1em;margin-bottom:1rem;border-bottom:1.5px solid var(--border);padding-bottom:0.6rem;}
.pl-dim{color:var(--muted2);}
.pl-blue{color:#2b6cb0;}
.pl-lime{color:#6a8c00;}
.pl-rust{color:#c0622a;}
.pl-purple{color:#6b4fa0;}

/* ═══════════════ SERVICE CARDS ═══════════════ */
.services-grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(260px,1fr));gap:1rem;margin-top:1.5rem;}
.scard{background:#fff;border:1.5px solid var(--border);border-radius:12px;padding:1.4rem;position:relative;overflow:hidden;transition:border-color 0.2s,transform 0.2s,box-shadow 0.2s;}
.scard:hover{border-color:var(--lblue3);transform:translateY(-3px);box-shadow:0 8px 24px rgba(13,34,64,0.08);}
.scard::before{content:'';position:absolute;top:0;left:0;right:0;height:3px;}
.scard.go::before    {background:linear-gradient(90deg,#00add8,transparent);}
.scard.py::before    {background:linear-gradient(90deg,#2b6cb0,transparent);}
.scard.rust::before  {background:linear-gradient(90deg,#c0622a,transparent);}
.scard.ui::before    {background:linear-gradient(90deg,var(--lime2),transparent);}
.scard-num{font-family:var(--mono);font-size:10px;letter-spacing:0.1em;color:var(--muted2);text-transform:uppercase;margin-bottom:0.2rem;}
.scard-title{font-family:var(--mono);font-size:0.9rem;font-weight:700;margin-bottom:0.6rem;}
.scard.go .scard-title   {color:#00add8;}
.scard.py .scard-title   {color:#2b6cb0;}
.scard.rust .scard-title {color:#c0622a;}
.scard.ui .scard-title   {color:#6a8c00;}
.scard-desc{font-size:13px;color:var(--muted);line-height:1.65;}
.scard-why{margin-top:0.9rem;padding-top:0.7rem;border-top:1.5px solid var(--border);font-family:var(--mono);font-size:11px;color:var(--muted2);}
.scard-why strong{color:var(--navy2);}

/* ═══════════════ HIGHLIGHT BOX ═══════════════ */
.hbox{background:var(--lblue2);border:1.5px solid var(--lblue3);border-left:4px solid var(--lblue3);border-radius:0 10px 10px 0;padding:1.2rem 1.5rem;margin:1.5rem 0;}
.hbox h4{font-family:var(--mono);font-size:11px;letter-spacing:0.08em;text-transform:uppercase;color:var(--navy2);margin-bottom:0.4rem;}
.hbox p{color:var(--navy2);margin:0;font-size:14px;}

.limebox{background:#f9ffd6;border:1.5px solid #c8e030;border-left:4px solid var(--lime2);border-radius:0 10px 10px 0;padding:1.2rem 1.5rem;margin:1.5rem 0;}
.limebox h4{font-family:var(--mono);font-size:11px;letter-spacing:0.08em;text-transform:uppercase;color:#4a5a00;margin-bottom:0.4rem;}
.limebox p{color:#3a4800;margin:0;font-size:14px;}

/* ═══════════════ TECH TABLE ═══════════════ */
.tech-table{width:100%;border-collapse:collapse;font-size:13px;margin-top:1rem;background:#fff;border-radius:12px;overflow:hidden;border:1.5px solid var(--border);}
.tech-table th{font-family:var(--mono);font-size:10px;letter-spacing:0.12em;text-transform:uppercase;color:var(--muted2);padding:10px 14px;text-align:left;background:var(--bg2);border-bottom:1.5px solid var(--border);}
.tech-table td{padding:11px 14px;border-bottom:1px solid var(--border);color:var(--muted);}
.tech-table tr:last-child td{border-bottom:none;}
.tech-table tr:hover td{background:var(--bg2);}
.tn{font-family:var(--mono);font-weight:700;}

/* ═══════════════ VULN DEMO ═══════════════ */
.vuln-demo{background:#fff;border:1.5px solid var(--border);border-radius:12px;overflow:hidden;margin:1.5rem 0;position:relative;}
.vuln-demo::after{content:'';position:absolute;left:0;right:0;top:0;height:2px;background:linear-gradient(90deg,transparent,#e85555,transparent);animation:scan 2.5s linear infinite;}
@keyframes scan{0%{top:0;}100%{top:100%;}}
.vuln-hdr{background:#fff5f5;border-bottom:1.5px solid #fdd;padding:10px 16px;display:flex;align-items:center;gap:8px;font-family:var(--mono);font-size:11px;}
.vtag{background:#ffe5e5;color:#c0352a;border:1.5px solid #fcc;border-radius:5px;padding:2px 8px;font-size:10px;letter-spacing:0.06em;}
.vtag-warn{background:#fff8e0;color:#a06000;border:1.5px solid #f5dea0;border-radius:5px;padding:2px 8px;font-size:10px;}
.vfile{color:var(--muted2);}
.vline{color:#c08000;margin-left:auto;}
.code-block{padding:1rem 1.2rem;font-family:var(--mono);font-size:12px;line-height:2;overflow-x:auto;}
.cl{display:flex;gap:12px;}
.cln{color:rgba(74,106,138,0.35);min-width:20px;user-select:none;}
.c-vuln{background:#fff0f0;border-left:2.5px solid #e85555;padding-left:6px;}
.c-normal{color:#4a6a8a;}
.c-kw{color:#6b4fa0;}
.c-str{color:#a06000;}
.c-fn{color:#2b6cb0;}
.c-cmt{color:#a8c0d4;}
.vuln-ftr{background:#fff5f5;border-top:1.5px solid #fdd;padding:9px 16px;font-family:var(--mono);font-size:11px;color:#c0352a;display:flex;align-items:center;gap:6px;}
.vscore{margin-left:auto;color:#a06000;}

/* ═══════════════ VS GRID ═══════════════ */
.vs-grid{display:grid;grid-template-columns:1fr 1fr;gap:1rem;margin-top:1rem;}
.vs-box{background:#fff;border:1.5px solid var(--border);border-radius:12px;padding:1.2rem;}
.vs-box h4{font-family:var(--mono);font-size:12px;font-weight:700;margin-bottom:0.9rem;padding-bottom:0.6rem;border-bottom:1.5px solid var(--border);}
.vs-box.bad  h4{color:#c0352a;}
.vs-box.good h4{color:#1a7a4a;}
.vi{display:flex;gap:8px;font-size:12px;color:var(--muted);margin-bottom:6px;align-items:flex-start;}
.vi::before{content:'';flex-shrink:0;margin-top:7px;width:5px;height:5px;border-radius:50%;}
.vs-box.bad  .vi::before{background:#e85555;}
.vs-box.good .vi::before{background:#1a9a5a;}

/* ═══════════════ TERMINAL ═══════════════ */
.terminal{background:#f4f8ff;border:1.5px solid var(--border);border-radius:12px;overflow:hidden;margin:1.5rem 0;font-family:var(--mono);font-size:12px;}
.tbar{background:var(--bg2);border-bottom:1.5px solid var(--border);padding:8px 12px;display:flex;align-items:center;gap:6px;}
.dot{width:10px;height:10px;border-radius:50%;}
.dot-r{background:#ff5f57;}.dot-y{background:#febc2e;}.dot-g{background:#28c840;}
.ttitle{margin-left:auto;font-size:10px;color:var(--muted2);letter-spacing:0.08em;}
.tbody{padding:1rem 1.2rem;line-height:2;}
.t-prompt{color:var(--lblue3);}
.t-cmd{color:var(--navy);}
.t-out{color:var(--muted2);}
.t-flag{color:#a06000;}
.t-ok{color:#1a7a4a;}
.cursor{display:inline-block;width:8px;height:14px;background:var(--lblue3);vertical-align:middle;animation:blink 1s step-end infinite;margin-left:2px;}
@keyframes blink{0%,100%{opacity:1}50%{opacity:0}}

/* ═══════════════ FEATURES ═══════════════ */
.features{display:grid;grid-template-columns:repeat(auto-fit,minmax(220px,1fr));gap:1rem;margin-top:1.5rem;}
.feat{background:#fff;border:1.5px solid var(--border);border-radius:12px;padding:1.2rem;transition:border-color 0.2s,transform 0.2s;}
.feat:hover{border-color:var(--lblue3);transform:translateY(-2px);}
.feat-icon{font-size:1.3rem;margin-bottom:0.6rem;}
.feat h4{font-family:var(--mono);font-size:13px;font-weight:700;color:var(--navy);margin-bottom:0.3rem;}
.feat p{font-size:12px;color:var(--muted2);margin:0;line-height:1.6;}

/* ═══════════════ LIME BADGE ═══════════════ */
.lime-badge{
  display:inline-block;
  background:var(--lime);
  color:#2a3800;
  font-family:var(--mono);
  font-size:10px;
  font-weight:700;
  letter-spacing:0.08em;
  padding:4px 12px;
  border-radius:6px;
  text-transform:uppercase;
  margin-bottom:1rem;
}

/* ═══════════════ FOOTER ═══════════════ */
.footer{text-align:center;padding:4rem 0 3rem;border-top:1.5px solid var(--border);margin-top:5rem;}
.footer-logo{font-family:var(--mono);font-size:1.1rem;font-weight:700;color:var(--navy);margin-bottom:0.4rem;}
.footer-sub{font-size:13px;color:var(--muted2);}
.footer-stack{margin-top:0.5rem;font-family:var(--mono);font-size:11px;color:var(--muted2);}
.footer-links{display:flex;gap:1.5rem;justify-content:center;margin-top:1.5rem;flex-wrap:wrap;}
.footer-links a{font-family:var(--mono);font-size:11px;letter-spacing:0.06em;text-transform:uppercase;color:var(--muted2);text-decoration:none;transition:color 0.15s;}
.footer-links a:hover{color:var(--navy);}

@media(max-width:600px){
  .vs-grid{grid-template-columns:1fr;}
  .hero h1{font-size:2rem;}
}
</style>
</head>
<body>

<!-- ═══════════════════ HERO ═══════════════════ -->
<div class="hero">
  <div class="blob blob1"></div>
  <div class="blob blob2"></div>
  <div class="blob blob3"></div>
  <div class="particles">
    <div class="p"></div><div class="p"></div><div class="p"></div>
    <div class="p"></div><div class="p"></div><div class="p"></div>
  </div>

  <div class="hero-inner">
    <div class="status-badge">
      <span class="status-dot"></span>
      Star Project &nbsp;·&nbsp; Startup Demo Day 2025
    </div>

    <h1>
      Real-Time<br>
      <span class="accent-lime">Security Audit</span><br>
      <span class="accent-blue">Pipeline</span>
    </h1>

    <p class="hero-desc">
      An automated security engineer that detects vulnerabilities in your code
      the moment you push — powered by vector semantics over AST, not regex.
    </p>

    <div class="hero-pills">
      <span class="pill pill-go">⬡ Go 1.22</span>
      <span class="pill pill-rust">⬢ Rust 1.78</span>
      <span class="pill pill-py">🐍 Python 3.12</span>
      <span class="pill pill-grpc">⬡ gRPC</span>
      <span class="pill pill-q">◈ Qdrant</span>
      <span class="pill pill-sb">⬡ Supabase</span>
      <span class="pill pill-ai">✦ AI Embeddings</span>
    </div>

    <div class="scroll-hint">
      <span>Scroll to explore</span>
      <div class="scroll-arrow"></div>
    </div>
  </div>
</div>

<div class="container">

<!-- ═══════════════════ STATS ═══════════════════ -->
<div class="stats fade-in">
  <div class="stat">
    <span class="stat-num lime">90%</span>
    <span class="stat-label">Cosine similarity<br>threshold for critical flags</span>
  </div>
  <div class="stat">
    <span class="stat-num">4</span>
    <span class="stat-label">Isolated microservices<br>in the pipeline</span>
  </div>
  <div class="stat">
    <span class="stat-num blue">3</span>
    <span class="stat-label">Languages in one<br>polyglot gRPC stream</span>
  </div>
  <div class="stat">
    <span class="stat-num lime">∞</span>
    <span class="stat-label">Real-time vulnerability<br>events via WebSocket</span>
  </div>
</div>

<hr class="div">

<!-- ═══════════════════ OVERVIEW ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Overview</div>
  <h2>What is this?</h2>
  <p>
    Every time a developer pushes code, this pipeline wakes up. It clones the diff,
    breaks the source files into abstract syntax trees, converts those AST tokens into
    vector embeddings via a free AI API, then runs cosine similarity against a curated
    database of known-vulnerable code patterns stored in Qdrant.
  </p>
  <p>
    Vulnerable functions flash red on a live SolidJS dashboard within seconds of the commit —
    with exact file name and line number.
  </p>

  <div class="hbox">
    <h4>Core Innovation</h4>
    <p>
      Instead of fragile regex rules, the engine uses <strong>AI vector embeddings over AST tokens</strong>.
      A function that is structurally equivalent to a SQL-injectable query gets caught — even if
      variable names, formatting, and syntax look completely different on the surface.
    </p>
  </div>
</div>

<!-- ═══════════════════ LIVE VULN DEMO ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Live Output</div>
  <h2>What the scanner catches</h2>
  <p>This is what gets flagged the moment a developer pushes the following pattern:</p>

  <div class="vuln-demo">
    <div class="vuln-hdr">
      <span class="vtag">CRITICAL</span>
      <span class="vtag-warn">SQL INJECTION</span>
      <span class="vfile">src/api/users.py</span>
      <span class="vline">line 47 ↓</span>
    </div>
    <div class="code-block">
      <div class="cl c-normal"><span class="cln">44</span><span><span class="c-kw">def</span> <span class="c-fn">get_user_by_name</span>(name):</span></div>
      <div class="cl c-normal"><span class="cln">45</span><span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="c-kw">conn</span> = db.connect()</span></div>
      <div class="cl c-normal"><span class="cln">46</span><span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="c-kw">cursor</span> = conn.cursor()</span></div>
      <div class="cl c-vuln"><span class="cln">47</span><span>&nbsp;&nbsp;&nbsp;&nbsp;query = <span class="c-str">f"SELECT * FROM users WHERE name = '{name}'"</span> <span class="c-cmt"># ⚠ UNSAFE</span></span></div>
      <div class="cl c-vuln"><span class="cln">48</span><span>&nbsp;&nbsp;&nbsp;&nbsp;cursor.execute(query)</span></div>
      <div class="cl c-normal"><span class="cln">49</span><span>&nbsp;&nbsp;&nbsp;&nbsp;<span class="c-kw">return</span> cursor.fetchall()</span></div>
    </div>
    <div class="vuln-ftr">
      ⚠ Unsanitized f-string interpolation in SQL — vector similarity 96.2%
      <span class="vscore">CVSS 9.1 / CRITICAL</span>
    </div>
  </div>
</div>

<hr class="div">

<!-- ═══════════════════ ARCHITECTURE ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Architecture</div>
  <h2>System overview</h2>
  <p>Four microservices, zero shared memory. All inter-service communication runs over gRPC with Protobuf-typed contracts.</p>

  <div class="pipeline">
<span class="pl-dim">┌──────────────────────────────────────────────────────┐</span>
<span class="pl-dim">│</span>  <span class="pl-blue">[ GitHub / GitLab Webhook ]</span>  <span class="pl-dim">→  push event fires on commit  │</span>
<span class="pl-dim">└───────────────────┬──────────────────────────────────┘</span>
                    <span class="pl-dim">│  HTTP POST</span>
                    <span class="pl-dim">▼</span>
<span class="pl-dim">╔═══════════════════════════════════════════════════════╗</span>
<span class="pl-dim">║</span>  <span class="pl-blue">SERVICE 1 — Go Webhook Receiver</span>                      <span class="pl-dim">║</span>
<span class="pl-dim">║</span>  <span class="pl-dim">Validates payload · goroutines per webhook · low idle</span>  <span class="pl-dim">║</span>
<span class="pl-dim">╚═══════════════════════════════════════════════════════╝</span>
                    <span class="pl-dim">│  gRPC Unary RPC</span>
                    <span class="pl-dim">▼</span>
<span class="pl-dim">╔═══════════════════════════════════════════════════════╗</span>
<span class="pl-dim">║</span>  <span class="pl-lime">SERVICE 2 — Python Static Analyzer</span>                   <span class="pl-dim">║</span>
<span class="pl-dim">║</span>  <span class="pl-dim">Git clone · AST parse (tree-sitter) · tokenize blocks</span> <span class="pl-dim">║</span>
<span class="pl-dim">╚═══════════════════════════════════════════════════════╝</span>
                    <span class="pl-dim">│  gRPC Bi-directional Stream</span>
                    <span class="pl-dim">▼</span>
<span class="pl-dim">╔═══════════════════════════════════════════════════════╗</span>
<span class="pl-dim">║</span>  <span class="pl-rust">SERVICE 3 — Rust Semantic Engine</span>                      <span class="pl-dim">║</span>
<span class="pl-dim">║</span>  <span class="pl-dim">AI embeddings · Qdrant query · cosine similarity ≥90%</span> <span class="pl-dim">║</span>
<span class="pl-dim">╚═══════════════════════════════════════════════════════╝</span>
          <span class="pl-dim">│  WebSocket stream         │  Async DB write</span>
          <span class="pl-dim">▼                          ▼</span>
<span class="pl-dim">  ┌─────────────────────┐  ┌─────────────────────┐</span>
  <span class="pl-dim">│</span> <span class="pl-lime">SolidJS Dashboard</span>    <span class="pl-dim">│  │</span> <span class="pl-purple">Supabase PostgreSQL</span>  <span class="pl-dim">│</span>
  <span class="pl-dim">│ live red flags + line│  │ vuln history · repos │</span>
  <span class="pl-dim">└─────────────────────┘  └─────────────────────┘</span>
  </div>
</div>

<!-- ═══════════════════ SERVICES ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Microservices</div>
  <h2>Service breakdown</h2>
  <p>Each service owns exactly one responsibility — isolated memory, isolated failure surface.</p>

  <div class="services-grid">
    <div class="scard go">
      <div class="scard-num">Service 01</div>
      <div class="scard-title">Go Webhook Receiver</div>
      <div class="scard-desc">
        Exposes a public HTTP endpoint for GitHub / GitLab push events.
        Validates HMAC signatures, extracts repository URLs and commit hashes,
        and fans out to goroutine workers for parallel processing.
      </div>
      <div class="scard-why">
        <strong>Why Go?</strong> Goroutines handle dozens of simultaneous webhooks
        at near-zero idle memory — fits comfortably inside free-tier Render / Fly.io caps.
      </div>
    </div>

    <div class="scard py">
      <div class="scard-num">Service 02</div>
      <div class="scard-title">Python Static Analyzer</div>
      <div class="scard-desc">
        Receives repo metadata via gRPC, clones the target commit into <code>/tmp</code>,
        and uses Python's built-in <code>ast</code> module and <code>tree-sitter</code>
        to decompose source into function and class blocks for downstream analysis.
      </div>
      <div class="scard-why">
        <strong>Why Python?</strong> tree-sitter supports 40+ languages — the engine
        works on Python, JavaScript, Go, Rust, and more without additional adapters.
      </div>
    </div>

    <div class="scard rust">
      <div class="scard-num">Service 03</div>
      <div class="scard-title">Rust Semantic Engine</div>
      <div class="scard-desc">
        Receives code blocks over a persistent bi-directional gRPC stream, calls a
        free AI API (Gemini / Groq) to generate embeddings, then queries Qdrant with
        cosine similarity. Any match above 90% is flagged as a critical risk immediately.
      </div>
      <div class="scard-why">
        <strong>Why Rust?</strong> Sub-millisecond vector score sorting, zero-cost
        network abstractions, and memory safety without garbage collection pauses.
      </div>
    </div>

    <div class="scard ui">
      <div class="scard-num">Layer 04</div>
      <div class="scard-title">Storage &amp; Live Dashboard</div>
      <div class="scard-desc">
        Supabase stores vulnerability histories, severity metrics, and the list of linked
        repositories. A SolidJS frontend connects directly to the Rust engine via WebSocket —
        flagged lines flash red in real-time showing exact file name and line number.
      </div>
      <div class="scard-why">
        <strong>Why SolidJS?</strong> Fine-grained reactivity means only the newly flagged
        line re-renders — no full virtual DOM diffing on every incoming event.
      </div>
    </div>
  </div>
</div>

<hr class="div">

<!-- ═══════════════════ TECH STACK ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Stack</div>
  <h2>Technology choices</h2>

  <table class="tech-table">
    <thead>
      <tr><th>Layer</th><th>Technology</th><th>Role</th></tr>
    </thead>
    <tbody>
      <tr><td>Gateway</td><td><span class="tn" style="color:#00add8">Go 1.22</span></td><td>High-throughput HTTP ingestion, goroutine concurrency</td></tr>
      <tr><td>Analysis</td><td><span class="tn" style="color:#2b6cb0">Python 3.12</span></td><td>AST parsing, tree-sitter, code block tokenization</td></tr>
      <tr><td>Inference</td><td><span class="tn" style="color:#c0622a">Rust 1.78</span></td><td>Vector math, gRPC server, Qdrant client, WebSocket emitter</td></tr>
      <tr><td>RPC</td><td><span class="tn" style="color:#6b4fa0">gRPC / Protobuf</span></td><td>Strongly-typed, bi-directional streaming across services</td></tr>
      <tr><td>Vector DB</td><td><span class="tn" style="color:#1a7a4a">Qdrant</span></td><td>Cosine similarity search over vulnerability embedding vectors</td></tr>
      <tr><td>Database</td><td><span class="tn" style="color:#1a7a4a">Supabase</span></td><td>PostgreSQL for persistent audit logs, repo registry, severity scores</td></tr>
      <tr><td>AI APIs</td><td><span class="tn" style="color:#a06000">Gemini / Groq</span></td><td>Free-tier embedding generation for code snippets</td></tr>
      <tr><td>Frontend</td><td><span class="tn" style="color:#1a6a40">SolidJS</span></td><td>Fine-grained reactive dashboard with live WebSocket binding</td></tr>
    </tbody>
  </table>
</div>

<!-- ═══════════════════ VS COMPARISON ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Differentiation</div>
  <h2>Why this beats regex scanners</h2>

  <div class="vs-grid">
    <div class="vs-box bad">
      <h4>✕ Traditional regex scanners</h4>
      <div class="vi">Miss semantically equivalent vulnerabilities</div>
      <div class="vi">Require hand-written rules per pattern</div>
      <div class="vi">Break with renamed variables or reformatting</div>
      <div class="vi">No understanding of code structure or flow</div>
      <div class="vi">High false-negative rate on obfuscated injection</div>
    </div>
    <div class="vs-box good">
      <h4>✓ Vector-semantic approach</h4>
      <div class="vi">Detects structural similarity regardless of naming</div>
      <div class="vi">Works across 40+ languages via tree-sitter</div>
      <div class="vi">Catches logical patterns, not just string patterns</div>
      <div class="vi">Trained on curated vulnerable code corpus</div>
      <div class="vi">90% threshold keeps false-positive rate low</div>
    </div>
  </div>
</div>

<!-- ═══════════════════ TERMINAL ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Getting Started</div>
  <h2>Quickstart</h2>

  <div class="terminal">
    <div class="tbar">
      <div class="dot dot-r"></div>
      <div class="dot dot-y"></div>
      <div class="dot dot-g"></div>
      <span class="ttitle">bash — vulnscan setup</span>
    </div>
    <div class="tbody">
      <div><span class="t-prompt">$ </span><span class="t-cmd">git clone https://github.com/your-org/vulnscan-pipeline</span></div>
      <div><span class="t-prompt">$ </span><span class="t-cmd">cd vulnscan-pipeline</span></div>
      <div>&nbsp;</div>
      <div><span class="t-prompt"># </span><span class="t-out">Start all four services with Docker Compose</span></div>
      <div><span class="t-prompt">$ </span><span class="t-cmd">docker compose up <span class="t-flag">--build</span></span></div>
      <div><span class="t-ok">✓ go-gateway      started on :8080</span></div>
      <div><span class="t-ok">✓ py-analyzer     started on :50051</span></div>
      <div><span class="t-ok">✓ rust-engine     started on :50052 + ws :9000</span></div>
      <div><span class="t-ok">✓ qdrant          started on :6333</span></div>
      <div><span class="t-ok">✓ solidjs-dash    started on :3000</span></div>
      <div>&nbsp;</div>
      <div><span class="t-prompt"># </span><span class="t-out">Register a repository to monitor</span></div>
      <div><span class="t-prompt">$ </span><span class="t-cmd">curl <span class="t-flag">-X POST</span> localhost:8080/repos \</span></div>
      <div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="t-cmd"><span class="t-flag">-d</span> '{"url": "https://github.com/you/myapp"}'</span></div>
      <div><span class="t-ok">{"status": "registered", "webhook_url": "https://your-host/webhook/abc123"}</span></div>
      <div>&nbsp;</div>
      <div><span class="t-prompt"># </span><span class="t-out">Open the live dashboard, then push a commit</span></div>
      <div><span class="t-prompt">$ </span><span class="t-cmd">open http://localhost:3000</span><span class="cursor"></span></div>
    </div>
  </div>
</div>

<!-- ═══════════════════ FEATURES ═══════════════════ -->
<div class="section fade-in">
  <div class="section-tag">Capabilities</div>
  <h2>What this project demonstrates</h2>

  <div class="features">
    <div class="feat">
      <div class="feat-icon">⬡</div>
      <h4>Polyglot gRPC Pipelines</h4>
      <p>Complex payloads — AST tokens, source code, vector arrays — flow across Go, Python, and Rust over Protobuf contracts.</p>
    </div>
    <div class="feat">
      <div class="feat-icon">◈</div>
      <h4>Semantic Code Analysis</h4>
      <p>AI embeddings are run over abstract syntax, not text. The engine understands code structure the way a human reviewer would.</p>
    </div>
    <div class="feat">
      <div class="feat-icon">⬢</div>
      <h4>Zero-overhead Inference</h4>
      <p>Rust handles vector sorting at native speed. Free-tier Gemini / Groq APIs eliminate any model-hosting costs entirely.</p>
    </div>
    <div class="feat">
      <div class="feat-icon">✦</div>
      <h4>Enterprise-grade Pipeline</h4>
      <p>Mirrors how SonarQube and Snyk operate under the hood — built from scratch on free-tier cloud infrastructure.</p>
    </div>
    <div class="feat">
      <div class="feat-icon">◌</div>
      <h4>Real-time Dashboard</h4>
      <p>Fine-grained SolidJS reactivity + WebSocket streaming = vulnerable lines appear red within seconds of a push event.</p>
    </div>
    <div class="feat">
      <div class="feat-icon">⊕</div>
      <h4>Free-tier Deployable</h4>
      <p>Go's minimal idle memory and containerized Python /tmp parsing means the full stack runs within Render or Fly.io free limits.</p>
    </div>
  </div>
</div>

<div class="limebox fade-in">
  <h4>Why this stands out</h4>
  <p>
    You are not building a text chatbot wrapper. You are using AI embeddings to run
    semantic search over abstract code syntax — detecting logical vulnerabilities that
    basic scanners miss. That is an engineering problem, not a prompt-engineering problem.
  </p>
</div>

<!-- ═══════════════════ FOOTER ═══════════════════ -->
<div class="footer fade-in">
  <div class="lime-badge">Star Project 2025</div>
  <div class="footer-logo">VulnScan AI</div>
  <div class="footer-sub">Real-Time Automated Security Code Audit &amp; Vulnerability Detection Pipeline</div>
  <div class="footer-stack">Go · Rust · Python · gRPC · Qdrant · Supabase · SolidJS</div>
  <div class="footer-links">
    <a href="#">Docs</a>
    <a href="#">API Reference</a>
    <a href="#">Architecture</a>
    <a href="#">GitHub</a>
  </div>
</div>

</div><!-- /container -->

<script>
const els = document.querySelectorAll('.fade-in');
const obs = new IntersectionObserver(entries => {
  entries.forEach((e, i) => {
    if (e.isIntersecting) {
      setTimeout(() => e.target.classList.add('vis'), i * 70);
      obs.unobserve(e.target);
    }
  });
}, { threshold: 0.08 });
els.forEach(el => obs.observe(el));
</script>
</body>
</html>
