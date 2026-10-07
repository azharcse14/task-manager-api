# 📚 আমার শেখার গাইড

> **AI মেন্টরের জন্য:** আমি যখন এই repo-র লিংক দেব, তখন উত্তর দেওয়ার আগে এই ফাইলটা পুরো পড়বে, আর নিচের নিয়ম মেনে আমাকে শেখাবে। আমাকে আলাদা করে আবার নির্দেশনা দিতে হবে না।
>
> `ROADMAP.md` আর এই ফাইল আলাদা জিনিস। `ROADMAP.md` বলে **কীভাবে বানাতে হয়**। এই ফাইল বলে **কীভাবে আমাকে শেখাতে হবে, আর কী শেখা বাকি**।

Repo: https://github.com/azharcse14/task-manager-api

---

## ০. শুরুতে AI যা করবে (এই ক্রমে)

1. repo clone করবে বা পড়বে। দেখবে: `git log --oneline --graph --all`, `git branch -a`, খোলা PR, merge না হওয়া branch।
2. কোড, `*_test.go` ফাইল আর `ROADMAP.md`-এর "আমার নোট" অংশ পড়বে।
3. কোডের সাথে এই ফাইলের **৬ নম্বর অংশ (বর্তমান অবস্থা)** আর **৭ নম্বর অংশ (শেখার পথ)** মিলিয়ে দেখবে।
4. তারপর ছোট করে বলবে:
    - শেষ আপডেটের পর আমি নতুন কী শিখেছি (কোন commit-এ তার প্রমাণ)
    - কোডে কোন ভুল বা উন্নতির জায়গা আছে
    - পরের **একটা** টপিক কী, আর কেন সেটাই
5. আমি নিজে কোনো টপিক চাইলে সেটা দিয়েই শুরু করবে।

---

## ১. আমি কে, আমার লক্ষ্য কী

- এই একটা প্রজেক্ট দিয়ে আমি **Go**, **Git/GitHub**, আর **industry-র নিয়ম মেনে backend** শিখছি।
- ধীরে ধীরে advanced database আর product বানাতে যা যা লাগে, সব এই প্রজেক্টেই যোগ করে শিখতে চাই। যেমন Docker, PostgreSQL, auth, CI/CD, caching, observability, deployment।
- আমি **একদম শিশুর মতো, গোড়া থেকে** বুঝতে চাই। আমি কিছু জানি ধরে নিয়ে লাফিয়ে যাবে না।

---

## ২. কোন ভাষায়, কোন ঢঙে বোঝাবে

- **চলিত বাংলা।** সাধু ভাষা না।
- technical শব্দ **ইংরেজিতেই** রাখবে, যেমন context, pointer, transaction, middleware। জোর করে বাংলা বানাবে না। তবে কোনো শব্দ প্রথমবার এলে বাংলায় সহজ মানে বলে দেবে।
- ছোট ছোট বাক্যে বলবে। প্রতিদিনের জীবনের উদাহরণ দেবে, যেমন দোকান, চিঠি, লাইব্রেরি, রান্নাঘর।
- **এক সময়ে একটাই টপিক।** একসাথে অনেক নতুন জিনিস দেবে না।
- কোড দেখালে **আমার repo-র আসল ফাইল আর ফাংশন** ধরে দেখাবে, যেমন `internal/task/store.go`-এর `List`।
- কোডের comment বাংলায় লিখবে। variable আর function-এর নাম, আর commit message ইংরেজিতে লিখবে, আমার repo যেভাবে আছে।
- দরকার হলে ASCII ছবি বা ছোট diagram দিয়ে বোঝাবে।

---

## ৩. প্রতিটা টপিক এই ৭টা প্রশ্নের উত্তর দিয়ে বোঝাবে

| # | প্রশ্ন | কী থাকবে |
|---|---|---|
| ১ | **এটা কী?** | আগে শিশুকে বলার মতো এক লাইন, তারপর একটু বড় করে |
| ২ | **কেন লাগে?** | এটা না থাকলে আমার প্রজেক্টে কী সমস্যা হতো। আগে সমস্যাটা দেখাবে, তারপর সমাধান |
| ৩ | **কীভাবে ব্যবহার করব?** | আমার প্রজেক্টের কোন ফাইলে কী বদলাবে, ছোট ছোট ধাপে |
| ৪ | **ভেতরে আর মেমোরিতে কী হয়?** | এটা stack-এ থাকে না heap-এ, কত byte নেয়, কী কপি হয় আর কী হয় না। দরকার হলে pointer, goroutine, OS, disk, network পর্যন্ত, যতটা গভীরে যাওয়া যায় |
| ৫ | **বিকল্প কী?** | অন্য উপায়, লাইব্রেরি বা টুল। কোনটা কখন ভালো, industry-তে কোনটা বেশি চলে |
| ৬ | **সাধারণ ভুল** | নতুনরা কোথায় আটকায়, কোন ফাঁদে পড়ে |
| ৭ | **নিজে করো** | ছোট একটা কাজ। কোন test বা command চালিয়ে যাচাই করব। branch-এর নাম আর commit message কী হবে |

---

## ৪. শেখানোর নিয়ম

- **আগে বোঝাবে, তারপর আমাকে লিখতে দেবে।** আটকে গেলে আগে hint দেবে। আমি "পুরো কোড দাও" বললে পুরো কোড দেবে। তখনও প্রতিটা লাইন কেন লেখা হলো, সেটা বোঝাবে।
- **প্রতিটা নতুন কাজের জন্য নতুন branch আর নতুন PR।** branch-এর নাম শুরু হবে `feature/`, `fix/`, `refactor/`, `test/`, `docs/` বা `ci/` দিয়ে।
- **নতুন কোডের সাথে test লিখতে হবে।** test ছাড়া কোনো কাজ শেষ ধরবে না।
- আমার কোডে ভুল বা খারাপ অভ্যাস দেখলে **সোজাসুজি বলবে**, আর কেন সেটা খারাপ তা বোঝাবে। প্রশংসা দিয়ে ভুলটা ঢেকে দেবে না।
- আমি **Go 1.27 বা তার নতুন ভার্সন** ব্যবহার করি। তাই `encoding/json/v2`, নতুন `net/http` routing, `min`/`max`, `slog` এসব ধরে নিয়ে উত্তর দেবে। পুরনো Go-র নিয়মে উত্তর দেবে না। নিশ্চিত না হলে release notes দেখে নেবে।
- সেশনের শেষে বলে দেবে এই ফাইলের **৬ আর ৭ নম্বর অংশে** কী আপডেট করতে হবে, যাতে আমি commit করে রাখতে পারি।

---

## ৫. পরের বার AI-কে যা লিখব

```text
আমার repo: https://github.com/azharcse14/task-manager-api
প্রথমে LEARNING_GUIDE.md পড়ো, তারপর সেই নিয়ম মেনে চলো।
আজ শিখতে চাই: ________   (ফাঁকা রাখলে তুমিই পরের ধাপ ঠিক করো)
```

---

## ৬. বর্তমান অবস্থা

**শেষ আপডেট:** ৮ অক্টোবর ২০২৬
**main-এর শেষ commit:** `aa19982` (PR #10, HTTP tests)

### ✅ যা শিখেছি (কোডে প্রমাণ আছে)

- **Go-র ভিত্তি:** module (`go.mod`), package, struct, struct tag, method, pointer (`*bool`), error handling, `errors.Is`
- **HTTP:**
    - `net/http` সার্ভার
    - Go 1.22+ routing (`GET /tasks/{id}`, `r.PathValue`)
    - status code: 200, 201, 204, 400, 404, 405, 500
    - header: `Content-Type`, `Location`, `X-Total-Count`
- **JSON:** `encoding/json/v2` (`MarshalWrite`, `UnmarshalRead`)। শিখেছি যে field-এর নাম case-sensitive।
- **Concurrency (শুরু):** মেমোরিতে ডেটা রাখার সময় `sync.Mutex` লাগে, আর database-এ যাওয়ার পর কেন আর লাগে না।
- **Database:**
    - SQLite (`modernc.org/sqlite`) আর `database/sql`
    - `?` দিয়ে parameterized query, যেটা SQL injection আটকায়
    - `LastInsertId`, `RowsAffected`, `sql.ErrNoRows`
    - `LIMIT/OFFSET` দিয়ে pagination, `COUNT(*)`, শর্ত অনুযায়ী বদলানো `WHERE`
- **Structure:**
    - `internal/` package
    - data layer (`task`) আর HTTP layer (`handler`) আলাদা
    - constructor দিয়ে dependency injection (`NewStore`, `NewTaskHandler`)
- **Logging:** `log/slog`, তবে এখনো শুধু একটা জায়গায়।
- **Testing:**
    - unit test, store-এর integration test, `httptest` দিয়ে HTTP test
    - table-driven test, `t.Run`, `t.Helper`, `t.Cleanup`
    - `:memory:` SQLite সাথে `SetMaxOpenConns(1)`
- **Git/GitHub:** feature branch, PR, merge, ছোট commit, ইংরেজিতে কাজের ভাষায় লেখা commit message, README, ROADMAP

### 🟡 অর্ধেক হয়ে আছে

**`feature/migrations` branch** (commit `3579292`, WIP, merge হয়নি)।

এখানে যা করা হয়েছে:
- `internal/migrate` প্যাকেজ আর `schema_migrations` টেবিল
- `deleted_at` column
- `GET /tasks/trash` আর `POST /tasks/{id}/restore`

যা বাকি:
- `Delete` এখনো আসল `DELETE` করে। মানে soft delete এখনো হয়নি।
- `List`, `GetByID` আর `Update` এখনো `deleted_at IS NULL` দেখে না।
- migration একটা transaction-এর ভেতরে চলে না।
- কোনো test নেই।
- branch-টা পুরনো `main` থেকে বানানো, তাই এখনকার `main`-এর সাথে মেলাতে হবে।

### 🔧 কোড রিভিউ থেকে যা ঠিক করা বাকি

- [ ] `handler` প্যাকেজ `database/sql` import করে `sql.ErrNoRows` চেক করে। মানে HTTP layer database-এর কথা জানে। → `task` প্যাকেজে নিজের `ErrNotFound` বানাও।
- [ ] `context` আছে শুধু `List`-এ। `GetByID`, `Create`, `Update`, `Delete`-এ নেই। → সবগুলোতে `ctx` নাও, আর `QueryRowContext`/`ExecContext` ব্যবহার করো।
- [ ] 500 error হলে আসল কারণ কোথাও log হয় না। → `slog.Error` দিয়ে log করো। client-কে শুধু সাধারণ একটা message দাও।
- [ ] request body কত বড় হতে পারবে, তার কোনো সীমা নেই। কেউ বিশাল body পাঠালে সেটা মেমোরি খেয়ে ফেলতে পারে। → `http.MaxBytesReader` ব্যবহার করো।
- [ ] title-এ শুধু space (`"   "`) দিলেও পাস করে যায়। title কত লম্বা হতে পারবে তারও কোনো সীমা নেই। → `strings.TrimSpace` দিয়ে space ছেঁটে ফেলো। দৈর্ঘ্য গুনতে `utf8.RuneCountInString` ব্যবহার করো, কারণ একটা বাংলা অক্ষর ৩ byte নেয়, তাই `len()` ভুল সংখ্যা দেয়।
- [ ] `/health` endpoint refactor করার সময় হারিয়ে গেছে। ROADMAP-এ আছে, কিন্তু `routes.go`-তে নেই।
- [ ] `main.go`-তে `log` ব্যবহার হচ্ছে, `handler`-এ `slog`। → সব জায়গায় `slog` ব্যবহার করো।
- [ ] সার্ভার চলে সরাসরি `http.ListenAndServe` দিয়ে। কোনো timeout নেই, graceful shutdown-ও নেই।
- [ ] port (`:8080`) আর database ফাইলের path (`tasks.db`) কোডের ভেতরে সরাসরি লেখা।
- [ ] SQLite-এ WAL mode আর `busy_timeout` সেট করা নেই। একসাথে অনেক write এলে `database is locked` error আসতে পারে।
- [ ] `List`-এ `COUNT` আর `SELECT` দুটো আলাদা query। মাঝখানে কেউ task যোগ করলে `total` আর `data` না-ও মিলতে পারে। (transaction শেখার সময় এটা দেখবে)
- [ ] README-র Roadmap-এ "Automated tests" এখনো টিক দেওয়া হয়নি, অথচ test লেখা হয়ে গেছে।
- [ ] **Git অভ্যাস:** PR #6 merge হওয়ার পরেও একই `feature/filtering` branch-এ README আর ROADMAP commit করে আবার PR #7 খোলা হয়েছে। → merge হয়ে গেলে branch মুছে ফেলো। নতুন কাজের জন্য নতুন branch খোলো (যেমন `docs/readme`)। GitHub-এ "Automatically delete head branches" চালু করে দাও।

---

## ৭. শেখার পথ (ধাপে ধাপে)

> নিয়ম: উপরের ধাপ শেষ না করে নিচের ধাপে যাবে না। প্রতিটা ধাপ এক বা একাধিক ছোট PR দিয়ে করবে।

### ধাপ ০: ভিত্তি ✅
- [x] HTTP server, routing, JSON
- [x] CRUD
- [x] প্রথমে মেমোরিতে ডেটা আর mutex, তারপর SQLite
- [x] কোড package-এ ভাগ করা
- [x] pagination আর filter
- [x] test (unit, integration, HTTP)

### ধাপ ১: CI, GitHub Actions দিয়ে ⬅️ **এখন এটা**
> কেন এখন: test লেখা হয়ে গেছে। এখন প্রতিটা PR-এ সেগুলো নিজে থেকেই চলুক, যাতে ভুল কোড `main`-এ ঢুকতে না পারে।
- [ ] CI কী। workflow, job, step আর runner কী।
- [ ] `.github/workflows/ci.yml` ফাইলে `gofmt` চেক, `go vet` আর `go test -race ./...`
- [ ] `-race` কী ধরে। এটা data race ধরে, মানে যখন দুটো goroutine একই মেমোরিতে একসাথে লেখে।
- [ ] branch protection চালু করা, যাতে CI pass না করলে merge করা না যায়।
- [ ] `govulncheck` দিয়ে dependency-তে নিরাপত্তার ত্রুটি খোঁজা।
- [ ] README-তে CI badge লাগানো।

### ধাপ ২: কোড পরিষ্কার করা
> কেন: সামনের বড় কাজের আগে ভিত শক্ত করা। ৬ নম্বর অংশের "কোড রিভিউ" তালিকা থেকে একটা একটা করে ঠিক করবে, প্রতিটা আলাদা ছোট PR-এ।
- [ ] সব store method-এ `context`
- [ ] `task.ErrNotFound`। এখানে শিখবে domain error, sentinel error আর `%w` দিয়ে error wrapping।
- [ ] 500 error-এর log
- [ ] `http.MaxBytesReader`
- [ ] title validation
- [ ] `/health` ফেরত আনা

### ধাপ ৩: Database migrations (`feature/migrations` শেষ করা)
- [ ] migration কী, আর কেন শুধু `CREATE TABLE IF NOT EXISTS` যথেষ্ট না
- [ ] প্রতিটা migration একটা transaction-এর ভেতরে চালানো
- [ ] `created_at` আর `updated_at` column
- [ ] soft delete: `deleted_at`, trash, restore, আর সব query-তে `deleted_at IS NULL`
- [ ] migration আর soft delete-এর test
- [ ] বিকল্পগুলো জানা: `goose`, `golang-migrate`, `atlas`

### ধাপ ৪: Production-এর মতো server
- [ ] config environment variable থেকে পড়া (`PORT`, `DB_PATH`), আর 12-factor app কী
- [ ] `http.Server`-এর timeout: `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`
- [ ] graceful shutdown (`signal.NotifyContext`, `srv.Shutdown`)
- [ ] SQLite-এ WAL আর `busy_timeout`, আর connection pool (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`)

### ধাপ ৫: Middleware আর Logging
- [ ] middleware কী (`func(http.Handler) http.Handler`)
- [ ] প্রতিটা request-এর log: method, path, status, আর কত সময় লাগল
- [ ] panic recovery
- [ ] request ID, আর `context`-এ মান রাখা
- [ ] সব জায়গায় `slog`, আর JSON format-এ log

### ধাপ ৬: API আরও ভালো করা
- [ ] `PATCH` দিয়ে partial update। pointer field দিয়ে বোঝা যায় কোনো field "দেওয়া হয়নি"।
- [ ] validation-এর কোড আলাদা জায়গায় রাখা
- [ ] error-এর standard format (RFC 9457 Problem Details)
- [ ] অজানা JSON field বাতিল করা (`json.RejectUnknownMembers`)
- [ ] sorting (`?sort=-created_at`) আর search (`?q=`)
- [ ] idempotency কী, আর PUT, POST, PATCH-এ এটা কীভাবে আলাদা

### ধাপ ৭: Interface আর testing গভীরে
- [ ] interface কী, আর কেন interface-টা যে ব্যবহার করবে তার দিকে রাখতে হয় ("accept interfaces, return structs")
- [ ] নকল (fake) store দিয়ে handler test
- [ ] test coverage (`go test -cover`) আর benchmark (`testing.B`)

### ধাপ ৮: Docker আর PostgreSQL
- [ ] container আর image কী, VM-এর সাথে পার্থক্য কী
- [ ] multi-stage `Dockerfile` আর `.dockerignore`
- [ ] `docker compose` দিয়ে api আর postgres একসাথে চালানো
- [ ] `pgx` driver, `$1` placeholder, আর SQLite থেকে PostgreSQL-এ যাওয়া
- [ ] PostgreSQL দিয়ে integration test (`testcontainers-go`)
- [ ] raw SQL, `sqlc` আর ORM (`GORM`)-এর তুলনা

### ধাপ ৯: User আর Authentication
- [ ] `users` টেবিল, foreign key, task-এর মালিক, আর `JOIN`
- [ ] password hashing (`bcrypt` বা `argon2id`), আর কেন password plain text-এ রাখা যায় না
- [ ] login: session cookie বনাম JWT
- [ ] authorization, যাতে কেউ অন্যের task দেখতে বা বদলাতে না পারে
- [ ] CORS

### ধাপ ১০: Advanced Database
- [ ] index, B-tree, আর `EXPLAIN ANALYZE`
- [ ] transaction, ACID, isolation level, lock আর deadlock
- [ ] N+1 problem
- [ ] keyset (cursor) pagination, আর বড় `OFFSET` কেন ধীর
- [ ] full-text search
- [ ] constraint (`UNIQUE`, `CHECK`) আর normalization
- [ ] backup, replication আর read replica-র ধারণা

### ধাপ ১১: Concurrency আর Performance
- [ ] goroutine, channel, `select`, `sync.WaitGroup`, `errgroup`
- [ ] worker pool আর background job
- [ ] `pprof` দিয়ে CPU আর memory profile, আর escape analysis
- [ ] load test (`k6` বা `hey`)

### ধাপ ১২: Cache আর Rate limit
- [ ] Redis, cache-aside pattern, আর cache invalidation
- [ ] rate limiting (token bucket)

### ধাপ ১৩: Observability
- [ ] liveness check বনাম readiness check
- [ ] metrics (Prometheus)
- [ ] tracing (OpenTelemetry)

### ধাপ ১৪: Docs আর Delivery
- [ ] OpenAPI দিয়ে API docs
- [ ] API versioning (`/v1`)
- [ ] release: git tag, semantic versioning, CHANGELOG
- [ ] CD আর cloud-এ deploy করা
- [ ] message queue-র ধারণা

---

## ৮. পাশাপাশি চলবে: Go-র ভেতরের জিনিস আর মেমোরি

> এগুলো আলাদা কোনো ধাপ না। উপরের কোনো ধাপে সম্পর্কিত কোড এলে AI তখনই এগুলো ধরে ভেতরের গল্পটা বলবে।

- [ ] variable কোথায় থাকে: stack না heap, আর escape analysis (`go build -gcflags=-m`)
- [ ] pointer, value receiver বনাম pointer receiver, আর কখন কপি হয়
- [ ] slice-এর ভেতরে কী থাকে (pointer, len আর cap মিলে ২৪ byte), `append` কীভাবে বড় হয়, আর `make([]T, 0, n)`
- [ ] map-এর ভেতরে কী থাকে, আর কেন একসাথে অনেক goroutine একটা map-এ লিখলে সমস্যা হয়
- [ ] string, byte, rune আর UTF-8 (একটা বাংলা অক্ষর কেন ৩ byte)
- [ ] interface-এর ভেতরে কী থাকে (type আর data মিলে দুটো word), আর nil interface-এর ফাঁদ
- [ ] goroutine কীভাবে ছোট stack দিয়ে শুরু হয় আর দরকারে বাড়ে, আর scheduler (G-M-P)
- [ ] `net/http` কেন প্রতিটা connection-এর জন্য আলাদা goroutine চালায়
- [ ] garbage collector, `GOGC`, `GOMEMLIMIT`
- [ ] `io.Reader` দিয়ে একটু একটু করে পড়া (streaming) বনাম সব একবারে মেমোরিতে আনা
- [ ] `database/sql`-এর connection pool ভেতরে কীভাবে কাজ করে
- [ ] OS-এর দিক: process, file descriptor, socket, আর disk-এ লেখা (`fsync`)

---

## ৯. আমার প্রশ্নের খাতা

> সেশনে যে প্রশ্ন মাথায় আসে কিন্তু তখন উত্তর জানা হয়নি, সেগুলো এখানে লিখে রাখব। AI পরের বার এগুলো দেখে উত্তর দেবে।

-