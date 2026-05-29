<div align="center">

# 🛡️ Codesentinel 
### Real-Time Automated Security Code Audit & Vulnerability Detection Pipeline

> An automated security engineer that detects vulnerabilities in your code the moment you push — powered by vector semantics over AST, not regex.

<br/>

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Rust](https://img.shields.io/badge/Rust-1.78-DE7A38?style=for-the-badge&logo=rust&logoColor=white)
![Python](https://img.shields.io/badge/Python-3.12-3776AB?style=for-the-badge&logo=python&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-Protobuf-6B4FA0?style=for-the-badge&logo=grpc&logoColor=white)
![Qdrant](https://img.shields.io/badge/Qdrant-VectorDB-1A7A4A?style=for-the-badge)
![Supabase](https://img.shields.io/badge/Supabase-PostgreSQL-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)

<br/>

![Star Project](https://img.shields.io/badge/⭐_STAR_PROJECT-Demo_Day_2025-D4F542?style=for-the-badge&labelColor=2A3800&color=BFE320)
![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

</div>

---

## 📌 What is this?

Every time a developer pushes code, this pipeline wakes up. It:

1. **Receives** the GitHub / GitLab webhook push event
2. **Clones** the commit diff and parses it into Abstract Syntax Trees
3. **Tokenizes** the AST into function and class blocks
4. **Embeds** each block using a free AI API (Gemini / Groq)
5. **Queries** Qdrant via cosine similarity against known-vulnerable code vectors
6. **Flags** any match above **90% similarity** as a critical security risk — live, in seconds

Vulnerable functions flash red on a SolidJS dashboard with exact **file name and line number** — before the code ever reaches production.

> **Core innovation:** Instead of fragile regex rules, we use AI vector embeddings over AST tokens. A function structurally equivalent to a SQL-injectable query is caught — even if variable names, formatting, and syntax look completely different on the surface.

---

## 📊 Key Metrics

| Metric | Value |
|---|---|
| Cosine similarity threshold for critical flag | **≥ 90%** |
| Microservices in the pipeline | **4** |
| Languages in one gRPC stream | **3** (Go, Python, Rust) |
| Supported source languages (tree-sitter) | **40+** |
| Vulnerability stream delivery | **Real-time via WebSocket** |
| Infrastructure cost at free tier | **$0** |

---

## 🏗️ System Architecture

```
┌─────────────────────────────────────────────────────────┐
│   [ GitHub / GitLab Webhook ]  →  push event on commit  │
└───────────────────┬─────────────────────────────────────┘
                    │  HTTP POST
                    ▼
╔════════════════════════════════════════════════════════╗
║  SERVICE 1 — Go Webhook Receiver                       ║
║  Validates payload · goroutines per webhook · low idle ║
╚════════════════════════════════════════════════════════╝
                    │  gRPC Unary RPC
                    ▼
╔════════════════════════════════════════════════════════╗
║  SERVICE 2 — Python Static Analyzer                    ║
║  Git clone · AST parse (tree-sitter) · tokenize blocks ║
╚════════════════════════════════════════════════════════╝
                    │  gRPC Bi-directional Stream
                    ▼
╔════════════════════════════════════════════════════════╗
║  SERVICE 3 — Rust Semantic Engine                      ║
║  AI embeddings · Qdrant query · cosine similarity ≥90% ║
╚════════════════════════════════════════════════════════╝
          │  WebSocket stream          │  Async DB write
          ▼                            ▼
  ┌──────────────────────┐   ┌──────────────────────┐
  │  SolidJS Dashboard   │   │  Supabase PostgreSQL  │
  │  live red flags +    │   │  vuln history · repos │
  │  file · line number  │   │  severity metrics     │
  └──────────────────────┘   └──────────────────────┘
```

---

## 🚀 Microservice Breakdown

### Service 01 — Go Webhook Receiver *(The Ingress Gate)*

Exposes a public HTTP endpoint for GitHub / GitLab push events. Validates HMAC signatures, extracts repository URLs and commit hashes, and fans out to goroutine workers for parallel processing.

**Why Go?** Goroutines handle dozens of simultaneous webhooks at near-zero idle memory — fits comfortably inside free-tier Render / Fly.io memory caps.

---

### Service 02 — Python Static Analyzer *(The Parser)*

Receives repository metadata via gRPC, clones the target commit into `/tmp`, and uses Python's built-in `ast` module and `tree-sitter` to decompose source files into function and class blocks for downstream embedding.

**Why Python?** `tree-sitter` supports 40+ languages — the engine analyzes Python, JavaScript, Go, Rust, and more without additional adapters.

---

### Service 03 — Rust Semantic Engine *(The Vulnerability Scanner)*

Receives code blocks over a persistent bi-directional gRPC stream. Calls a free AI API (Gemini / Groq) to generate embeddings, then queries Qdrant with cosine similarity. Any block matching a known-vulnerable pattern above 90% is immediately flagged as a critical risk.

**Why Rust?** Sub-millisecond vector score sorting, zero-cost network abstractions, and memory safety without garbage collection pauses.

---

### Layer 04 — Storage & Live Dashboard

**Supabase** (PostgreSQL) stores vulnerability histories, severity metrics, and the registry of linked repositories.

**SolidJS** connects directly to the Rust engine via WebSocket — flagged lines flash red in real-time showing exact file name and line number.

**Why SolidJS?** Fine-grained reactivity means only the newly flagged line re-renders — no full virtual DOM diffing on every incoming WebSocket event.

---

## ⚡ Live Output Example

This is what gets flagged the moment a developer pushes the following pattern:

```python
# src/api/users.py — line 47

def get_user_by_name(name):
    conn = db.connect()
    cursor = conn.cursor()
    query = f"SELECT * FROM users WHERE name = '{name}'"  # ⚠ UNSAFE
    cursor.execute(query)
    return cursor.fetchall()
```

```
┌─────────────────────────────────────────────────────────────┐
│  ⚠  CRITICAL   SQL INJECTION   src/api/users.py   line 47  │
│                                                             │
│  Unsanitized f-string interpolation in SQL query            │
│  Vector similarity: 96.2%                    CVSS: 9.1     │
└─────────────────────────────────────────────────────────────┘
```

---

## 🛠️ Technology Stack

| Layer | Technology | Role |
|---|---|---|
| Gateway | **Go 1.22** | High-throughput HTTP ingestion, goroutine concurrency |
| Analysis | **Python 3.12** | AST parsing, tree-sitter, code block tokenization |
| Inference | **Rust 1.78** | Vector math, gRPC server, Qdrant client, WebSocket emitter |
| RPC | **gRPC / Protobuf** | Strongly-typed, bi-directional streaming across all services |
| Vector DB | **Qdrant** | Cosine similarity search over vulnerability embedding vectors |
| Database | **Supabase** | PostgreSQL for audit logs, repo registry, severity scores |
| AI APIs | **Gemini / Groq** | Free-tier embedding generation for code snippets |
| Frontend | **SolidJS** | Fine-grained reactive dashboard with live WebSocket binding |

---

## 🆚 Why This Beats Regex Scanners

| | Traditional Regex Scanners | VulnScan AI (Vector Semantic) |
|---|---|---|
| Detection method | String pattern matching | Structural vector similarity |
| Variable renaming | ❌ Bypasses the scanner | ✅ Still detected |
| Language support | Per-language rule sets | ✅ 40+ via tree-sitter |
| Code structure understanding | ❌ None | ✅ Full AST context |
| Obfuscated injection | ❌ Missed | ✅ Caught at semantic level |
| False negative rate | High | Low (90% threshold tuned) |

---

## 🏁 Quickstart

```bash
# Clone the repository
git clone https://github.com/your-org/vulnscan-pipeline
cd vulnscan-pipeline

# Start all four services with Docker Compose
docker compose up --build
```

```
✓ go-gateway      started on :8080
✓ py-analyzer     started on :50051
✓ rust-engine     started on :50052 + ws :9000
✓ qdrant          started on :6333
✓ solidjs-dash    started on :3000
```

```bash
# Register a repository to monitor
curl -X POST localhost:8080/repos \
     -d '{"url": "https://github.com/you/myapp"}'

# {"status": "registered", "webhook_url": "https://your-host/webhook/abc123"}

# Add the webhook URL to your GitHub repo settings
# Then open the dashboard and push a commit — watch it scan live
open http://localhost:3000
```

---

## 🎯 What This Project Demonstrates

| Capability | Detail |
|---|---|
| **Polyglot gRPC Pipelines** | Complex payloads (AST tokens, source code, vector arrays) flow across Go, Python, and Rust over Protobuf |
| **Semantic Code Analysis** | AI embeddings over abstract syntax — not text — so the engine understands code structure |
| **Zero-overhead Inference** | Rust sorts vectors at native speed; free-tier AI APIs eliminate model-hosting costs |
| **Enterprise-grade Pipeline** | Mirrors how SonarQube and Snyk operate under the hood |
| **Real-time Dashboard** | Fine-grained SolidJS reactivity + WebSocket = vulnerable lines appear within seconds of push |
| **Free-tier Deployable** | Go's minimal idle memory + containerized Python /tmp = full stack within Render / Fly.io free limits |

> This is not a chatbot wrapper. This uses AI embeddings to run semantic search over abstract code syntax — detecting logical vulnerabilities that basic scanners miss. That is an engineering problem, not a prompt-engineering problem.

---

## 📁 Project Structure

```
vulnscan-pipeline/
├── go-gateway/          # Service 1 — HTTP webhook receiver
│   ├── main.go
│   └── proto/
├── py-analyzer/         # Service 2 — AST parser & tokenizer
│   ├── analyzer.py
│   └── proto/
├── rust-engine/         # Service 3 — Semantic vector engine
│   ├── src/main.rs
│   └── proto/
├── dashboard/           # SolidJS live frontend
│   └── src/
├── proto/               # Shared Protobuf definitions
│   └── pipeline.proto
├── docker-compose.yml
└── README.md
```


---

<div align="center">

**Go · Rust · Python · gRPC · Qdrant · Supabase · SolidJS**

⭐ Star this repo if you found it useful

</div>
