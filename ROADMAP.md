# Go REST API প্রজেক্ট রোডম্যাপ

এই ফাইলটা `task-manager-api` প্রজেক্ট দেখে বানানো। নতুন প্রজেক্ট বানানোর সময় এটা উপর থেকে নিচে পড়বে আর একটা একটা ধাপ করবে।

**কীভাবে ব্যবহার করবে:**

- প্রতিটা ধাপের পাশে `- [ ]` আছে। কাজ শেষ হলে সেটাকে `- [x]` করে দেবে।
- নতুন প্রজেক্টে `task-manager-api` আর `task`-এর জায়গায় নিজের নাম বসাবে। যেমন বই-এর প্রজেক্ট হলে `book-api` আর `book`।
- একদম নিচে "আমার নোট" অংশ আছে। সেখানে নিজের কথা লিখে রাখবে।

---

## ১. এক লাইনে পুরো ক্রম

ভুলে গেলে শুধু এই লাইনটা মনে রাখো:

> **মডেল → টেবিল → স্টোর → হেল্পার → হ্যান্ডলার → রাউট → মেইন → টেস্ট**

---

## ২. আসল নিয়ম একটাই

**যার উপর অন্যরা নির্ভর করে, সেটা আগে লেখো। যে অন্যের উপর নির্ভর করে, সেটা পরে লেখো।**

এই প্রজেক্টে কে কার উপর নির্ভর করে:

```
main.go  ──►  internal/handler  ──►  internal/task  ──►  ডাটাবেস
```

- `internal/task` শুধু ডাটাবেস ব্যবহার করে। তাই এটা **সবার আগে** লিখবে।
- `internal/handler` ব্যবহার করে `task`-কে। তাই এটা **task-এর পরে** লিখবে।
- `main.go` ব্যবহার করে `handler` আর `task` দুটোকেই। তাই এর পুরো কোড **সবার শেষে** লিখবে।

একটা জিনিস জেনে রাখো: একই ফোল্ডারের (package-এর) ভেতরে ফাংশন কোনটা উপরে আর কোনটা নিচে, Go তাতে কিছু মনে করে না। এই ক্রমটা শুধু **তোমার মাথা পরিষ্কার রাখার জন্য**। কারণ যে জিনিস এখনো লেখোনি, সেটা ব্যবহার করতে গেলে গুলিয়ে যাবে।

---

## ৩. পুরো প্রজেক্টের ম্যাপ

```
task-manager-api/
├── go.mod                  ← প্রজেক্টের নাম আর Go ভার্সন
├── go.sum                  ← লাইব্রেরির হিসাব (Go নিজে বানায়, হাতে লিখবে না)
├── .gitignore              ← কোন ফাইল Git-এ যাবে না
├── README.md               ← প্রজেক্টের পরিচয়
├── main.go                 ← সব জোড়া লাগিয়ে সার্ভার চালু করে
├── api.http                ← হাতে হাতে API চেক করার request
└── internal/
    ├── task/               ← ডেটার অংশ (ডাটাবেসের সাথে কথা বলে)
    │   ├── task.go         ← Task দেখতে কেমন (struct)
    │   ├── schema.go       ← ডাটাবেসে টেবিল বানায়
    │   └── store.go        ← ডাটাবেসে পড়া-লেখা করে (SQL)
    └── handler/            ← HTTP-র অংশ (request নেয়, উত্তর দেয়)
        ├── response.go     ← JSON উত্তর পাঠানোর helper
        ├── query.go        ← URL থেকে মান পড়ার helper
        ├── task.go         ← প্রতিটা endpoint কী করবে
        └── routes.go       ← কোন URL কোন ফাংশনে যাবে
```

**`internal` ফোল্ডার কেন?** Go-র নিয়ম হলো, `internal`-এর ভেতরের কোড বাইরের অন্য প্রজেক্ট import করতে পারে না। মানে এটা শুধু এই প্রজেক্টের নিজের জিনিস।

**`task` আর `handler` আলাদা কেন?** `task` জানে ডাটাবেস কীভাবে চালাতে হয়, কিন্তু HTTP চেনে না। `handler` জানে HTTP কীভাবে চালাতে হয়, কিন্তু SQL লেখে না। এক কাজ এক জায়গায় থাকলে খুঁজে পাওয়া আর বদলানো সহজ।

---

## ৪. ভাগ ১: প্রজেক্ট শুরু করা

### - [ ] ধাপ ১: GitHub-এ repo বানাও

1. GitHub-এ নতুন repository বানাও।
2. **Add a README file**-এ টিক দাও।
3. **.gitignore template**-এ **Go** বেছে নাও।
4. তারপর কম্পিউটারে নামাও:

```bash
git clone https://github.com/azharcse14/task-manager-api.git
cd task-manager-api
```

**কেন সবার আগে:** শুরু থেকেই Git থাকলে প্রতিটা ধাপ সেভ করে রাখা যায়। কিছু ভুল হলে আগের জায়গায় ফেরা যায়।

### - [ ] ধাপ ২: `go.mod` বানাও

```bash
go mod init github.com/azharcse14/task-manager-api
```

**কেন:** `go.mod` ছাড়া Go জানে না এটা একটা প্রজেক্ট। import-এর পথ (`github.com/azharcse14/task-manager-api/internal/task`) এই নাম থেকেই আসে।

### - [ ] ধাপ ৩: `.gitignore`-এ নিজের জিনিস যোগ করো

ফাইলের একদম শেষে লেখো:

```gitignore
# Compiled binary for this project
task-manager-api

# SQLite database
*.db
```

**কেন:** ডাটাবেস ফাইল আর build করা প্রোগ্রাম Git-এ রাখার দরকার নেই। এগুলো প্রত্যেকের কম্পিউটারে নিজে নিজে তৈরি হয়।

### - [ ] ধাপ ৪: `main.go` (ছোট ভার্সন)

প্রথমে শুধু দেখো সার্ভার চলে কিনা:

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

চালাও: `go run .` তারপর ব্রাউজারে `http://localhost:8080/health` খোলো। `ok` দেখালে ঠিক আছে।

**কেন এত আগে:** শুরুতেই নিশ্চিত হও যে Go আর সার্ভার ঠিকমতো চলছে। পরে কিছু ভাঙলে বুঝবে সমস্যা নতুন কোডে, সেটআপে না।

👉 **commit করো:** `git commit -m "Add HTTP server with health endpoint"`

---

## ৫. ভাগ ২: কঙ্কাল বানানো (শুধু কাঠামো, এখনো কোনো feature না)

এই ভাগে প্রতিটা ফাইলের শুধু "খোলস" বানাবে। আসল কাজ (feature) আসবে ভাগ ৩-এ।

### - [ ] ধাপ ৫: `internal/task/task.go` (মডেল)

```go
package task

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
```

**কেন সবার আগে:** বাকি প্রায় সব ফাইল এই `Task` ব্যবহার করে। এটা না থাকলে কেউ জানে না একটা task-এর ভেতরে কী কী থাকে।

**`json:"id"` কী:** JSON-এ এই field-এর নাম কী হবে সেটা বলে দেয়। Go-তে `ID`, কিন্তু JSON-এ `id`।

### - [ ] ধাপ ৬: SQLite লাইব্রেরি আনো

```bash
go get modernc.org/sqlite
```

এতে `go.mod` আপডেট হবে আর `go.sum` নিজে থেকে তৈরি হবে।

**কেন `modernc.org/sqlite`:** এটা পুরোটা Go-তে লেখা। তাই আলাদা করে C compiler ইনস্টল করতে হয় না।

### - [ ] ধাপ ৭: `internal/task/schema.go` (টেবিল)

```go
package task

import "database/sql"

const createTasksTable = `
CREATE TABLE IF NOT EXISTS tasks (
	id    INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT    NOT NULL,
	done  BOOLEAN NOT NULL DEFAULT 0
)`

func CreateSchema(db *sql.DB) error {
	_, err := db.Exec(createTasksTable)
	return err
}
```

**কেন store-এর আগে:** টেবিল না থাকলে store-এর কোনো SQL চলবে না।

**`IF NOT EXISTS` কেন:** সার্ভার যতবার চালু হবে ততবার এটা চলবে। টেবিল আগে থেকে থাকলে আর নতুন করে বানাবে না।

### - [ ] ধাপ ৮: `internal/task/store.go` (শুধু খোলস)

```go
package task

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
```

**কেন:** `Store` হলো ডাটাবেসের সাথে কথা বলার একমাত্র জায়গা। এখনো কোনো method লিখবে না, সেগুলো ভাগ ৩-এ একটা একটা করে আসবে।

### - [ ] ধাপ ৯: `internal/handler/response.go` (helper)

```go
package handler

import (
	"encoding/json/v2"
	"log"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.MarshalWrite(w, data); err != nil {
		log.Println("failed to encode response:", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return 0, false
	}
	return id, true
}
```

**কেন handler-এর আগে:** প্রতিটা handler উত্তর পাঠাতে এগুলো ব্যবহার করবে। আগে না লিখলে প্রতিটা handler-এ একই কোড বারবার লিখতে হবে।

**ভেতরের ক্রম:** `writeJSON` → `writeError` → `parseID`। কারণ `writeError` ব্যবহার করে `writeJSON`-কে, আর `parseID` ব্যবহার করে `writeError`-কে।

**মনে রাখো:** import-এ অবশ্যই `"encoding/json/v2"` লিখবে, শুধু `"encoding/json"` না। নাহলে `MarshalWrite` খুঁজে পাবে না।

### - [ ] ধাপ ১০: `internal/handler/task.go` (শুধু খোলস)

```go
package handler

import "github.com/azharcse14/task-manager-api/internal/task"

type TaskHandler struct {
	store *task.Store
}

func NewTaskHandler(store *task.Store) *TaskHandler {
	return &TaskHandler{store: store}
}
```

**কেন:** handler-এর হাতে store থাকতে হবে, যাতে সে ডাটাবেস থেকে জিনিস আনতে পারে। এখানেই `handler` আর `task` জোড়া লাগে।

### - [ ] ধাপ ১১: `internal/handler/routes.go`

```go
package handler

import "net/http"

func (h *TaskHandler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	// নতুন route এখানে একটা একটা করে যোগ হবে

	return mux
}
```

**কেন handler-এর পরে:** route শুধু বলে দেয় কোন URL কোন handler-এ যাবে। handler না থাকলে route কোথাও পাঠাতে পারবে না।

### - [ ] ধাপ ১২: `main.go` (পুরো ভার্সন)

ধাপ ৪-এর ছোট `main.go` মুছে এটা লেখো:

```go
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/azharcse14/task-manager-api/internal/handler"
	"github.com/azharcse14/task-manager-api/internal/task"
)

func main() {
	// ১. ডাটাবেস খোলো
	db, err := sql.Open("sqlite", "tasks.db")
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}
	defer db.Close()

	// ২. ডাটাবেস সত্যিই চলছে কিনা দেখো
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// ৩. টেবিল বানাও
	if err := task.CreateSchema(db); err != nil {
		log.Fatal("Cannot create schema:", err)
	}

	// ৪. store → handler জোড়া লাগাও
	store := task.NewStore(db)
	taskHandler := handler.NewTaskHandler(store)

	// ৫. সার্ভার চালু করো
	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", taskHandler.Routes()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
```

**কেন সবার শেষে:** `main.go` সবাইকে ব্যবহার করে। তাই সবাই তৈরি হওয়ার পরেই এটা লেখা যায়।

**`_ "modernc.org/sqlite"` কী:** এটা SQLite driver-কে চালু করে। আমরা সরাসরি এর কোনো ফাংশন ডাকি না, তাই নামের জায়গায় `_` দিই। এটা না লিখলে `sql.Open("sqlite", ...)` কাজ করবে না।

চালাও: `go run .` তারপর `http://localhost:8080/health` খোলো।

👉 **commit করো:** `git commit -m "Add project skeleton with SQLite"`

---

## ৬. ভাগ ৩: প্রতিটা feature বানানোর রেসিপি

কঙ্কাল তৈরি। এবার প্রতিটা endpoint **একটা একটা করে** যোগ করবে। প্রতিবার এই ৫টা ধাপ, এই ক্রমেই:

| ক্রম | কোথায় | কী লিখবে |
|---|---|---|
| ১ | `internal/task/store.go` | ডাটাবেসের method (SQL) |
| ২ | `internal/handler/task.go` | handler method |
| ৩ | `internal/handler/routes.go` | একটা route লাইন |
| ৪ | `api.http` | একটা request |
| ৫ | টার্মিনাল | `go run .` → চেক → commit |

**কেন এই ক্রম:** সেই একই নিয়ম। handler ব্যবহার করে store-এর method-কে। route ব্যবহার করে handler-কে। তাই store আগে, handler পরে, route তার পরে।

**কেন একটা একটা করে:** একসাথে পাঁচটা endpoint লিখে চালালে কোনটায় ভুল আছে বোঝা কঠিন। একটা লিখে চালালে ভুল হলে সঙ্গে সঙ্গে বুঝবে।

### কোন feature আগে, কোনটা পরে

- [ ] **১. `POST /tasks`** (নতুন task বানানো)
  - store: `Create` · handler: `Create` · route: `mux.HandleFunc("POST /tasks", h.Create)`
  - **কেন প্রথমে:** ডেটা না থাকলে বাকি endpoint চেক করার কিছু থাকবে না।
- [ ] **২. `GET /tasks`** (সব task দেখা)
  - store: `List` · handler: `List` · route: `mux.HandleFunc("GET /tasks", h.List)`
  - **কেন দ্বিতীয়:** এখন দেখতে পাবে ধাপ ১-এ বানানো task সত্যিই ডাটাবেসে গেছে কিনা।
- [ ] **৩. `GET /tasks/{id}`** (একটা task দেখা)
  - store: `GetByID` · handler: `GetByID` · route: `mux.HandleFunc("GET /tasks/{id}", h.GetByID)`
  - **নতুন শেখা:** URL-এর `{id}` অংশ `r.PathValue("id")` দিয়ে পড়া।
- [ ] **৪. `PUT /tasks/{id}`** (task বদলানো)
  - store: `Update` · handler: `Update` · route: `mux.HandleFunc("PUT /tasks/{id}", h.Update)`
  - **নতুন শেখা:** `RowsAffected()` দিয়ে দেখা task সত্যিই ছিল কিনা। না থাকলে 404।
- [ ] **৫. `DELETE /tasks/{id}`** (task মোছা)
  - store: `Delete` · handler: `Delete` · route: `mux.HandleFunc("DELETE /tasks/{id}", h.Delete)`
  - **নতুন শেখা:** মুছে ফেলার পর 204 (No Content) পাঠানো।
- [ ] **৬. `GET /tasks`-এ pagination আর filter**
  - নতুন ফাইল `internal/handler/query.go` (intQuery, boolQuery)
  - store: `List`-এ `ListFilter` যোগ · handler: `List`-এ page, per_page, done পড়া
  - **কেন শেষে:** আগে সাধারণ List কাজ করুক, তারপর সেটাকে উন্নত করবে।
- [ ] **৭. পরের ধাপ (এখনো শিখিনি)**
  - [ ] টেস্ট (`*_test.go` ফাইল)
  - [ ] user আর login (JWT)
  - [ ] Docker আর PostgreSQL

প্রতিটার পুরো কোড তোমার `task-manager-api` রিপোতে আছে। আটকে গেলে সেখান থেকে দেখে নেবে।

---

## ৭. ফাইলের ভেতরে কোন কোড আগে

| ফাইল | ভেতরের ক্রম |
|---|---|
| `task/task.go` | package → struct |
| `task/schema.go` | package → import → SQL লেখা (`const`) → `CreateSchema` |
| `task/store.go` | `Store` struct → `NewStore` → `Create` → `List` → `GetByID` → `Update` → `Delete` |
| `handler/response.go` | `writeJSON` → `writeError` → `parseID` |
| `handler/query.go` | `intQuery` → `boolQuery` |
| `handler/task.go` | `TaskHandler` struct → `NewTaskHandler` → উত্তরের struct (যেমন `listResponse`) → handler method গুলো |
| `handler/routes.go` | `NewServeMux` → `HandleFunc` লাইনগুলো → `return mux` |
| `main.go` | ডাটাবেস খোলা → `Ping` → `CreateSchema` → `NewStore` → `NewTaskHandler` → `ListenAndServe` |

**একটা handler method-এর ভেতরের ক্রম** (প্রায় সব handler এভাবেই চলে):

1. request থেকে মান পড়ো (path, query, header, body)
2. মান ঠিক আছে কিনা চেক করো (ভুল হলে 400)
3. store-কে ডাকো
4. store-এর error চেক করো (না পেলে 404, অন্য সমস্যা হলে 500)
5. উত্তর পাঠাও (`writeJSON`)

---

## ৮. প্রতিটা feature-এর Git রুটিন

```bash
# ১. main আপডেট করো
git switch main
git pull

# ২. নতুন branch বানাও
git switch -c feature/create-task

# ৩. কোড লেখো, তারপর চেক করো
go vet ./...
go run .

# ৪. সেভ করো
git add .
git commit -m "Add POST endpoint to create tasks"

# ৫. GitHub-এ পাঠাও
git push -u origin feature/create-task
```

তারপর GitHub-এ **PR খোলো → merge করো**। শেষে আবার ধাপ ১ দিয়ে main আপডেট করো।

**branch-এর নাম:**

- নতুন জিনিস হলে `feature/...`
- ভুল ঠিক করলে `fix/...`
- কাজ না বদলে কোড সাজালে `refactor/...`

**commit message:** ইংরেজিতে, "Add", "Fix", "Replace" দিয়ে শুরু। কী করলে সেটা ছোট করে লেখো।

**একটা commit = একটা কাজ।** দুইটা আলাদা কাজ করলে দুইটা commit করো।

---

## ৯. দরকারি কমান্ড

| কমান্ড | কী করে |
|---|---|
| `go mod init <নাম>` | নতুন প্রজেক্ট শুরু করে, `go.mod` বানায় |
| `go get <লাইব্রেরি>` | লাইব্রেরি আনে, `go.sum` আপডেট করে |
| `go mod tidy` | অদরকারি লাইব্রেরি সরায়, দরকারিগুলো যোগ করে |
| `go run .` | প্রজেক্ট চালায় |
| `go build ./...` | শুধু compile করে দেখে কোনো ভুল আছে কিনা (কিছু না দেখালে সব ঠিক) |
| `go vet ./...` | সাধারণ ভুল খোঁজে |
| `go fmt ./...` | কোড নিয়ম মেনে সাজায় |
| `go version` | কোন Go ভার্সন চলছে দেখায় |

---

## ১০. প্রথম প্রজেক্টে তুমি যে ক্রমে বানিয়েছিলে

তোমার `task-manager-api`-এর commit দেখে (পুরোনো থেকে নতুন):

1. Initial commit (README, .gitignore)
2. Go module আর `main.go` বানানো
3. health endpoint দিয়ে সার্ভার
4. home আর health endpoint
5. সব task JSON-এ দেখানো
6. ID দিয়ে একটা task দেখা (path parameter)
7. JSON উত্তরের helper বানানো
8. POST দিয়ে task বানানো
9. PUT আর DELETE
10. `parseID` helper আলাদা করা (একই কোড বারবার না লেখার জন্য)
11. mutex যোগ করা (একসাথে অনেক request এলে ডেটা যেন নষ্ট না হয়)
12. memory থেকে SQLite ডাটাবেসে যাওয়া
13. কোড `task` আর `handler` প্যাকেজে ভাগ করা
14. pagination, filter আর JSON v2

**নতুন প্রজেক্টের ক্রম এটা থেকে আলাদা কেন?**

প্রথমবার তুমি শিখছিলে। তাই আগে সব `main.go`-তে লিখেছ, পরে ভাগ করেছ। শেখার জন্য এটাই ঠিক ছিল।

এখন তুমি জানো প্রজেক্ট শেষে দেখতে কেমন হবে। তাই নতুন প্রজেক্টে শুরু থেকেই ভাগ করা ফোল্ডারে লিখবে। এতে পরে সব সরানোর ঝামেলা থাকবে না।

**mutex এখন নেই কেন?** memory-তে ডেটা রাখার সময় mutex লাগত। SQLite-এ যাওয়ার পর আর লাগে না, কারণ Go-র `database/sql` নিজেই একসাথে অনেক request সামলাতে পারে।

---

## ১১. আমার নোট

এখানে নিজের কথা লিখে রাখো।

### তারিখ:

**কী শিখলাম:**

-

**কোথায় আটকে গেলাম:**

-

**কীভাবে ঠিক করলাম:**

-

---

### তারিখ:

**কী শিখলাম:**

-

**কোথায় আটকে গেলাম:**

-

**কীভাবে ঠিক করলাম:**

-

---

### যেসব ভুল বারবার করি

- [ ] import-এ `"encoding/json/v2"`-এর জায়গায় `"encoding/json"` লিখে ফেলি
- [ ] header `writeJSON`-এর পরে লিখে ফেলি (সবসময় আগে লিখতে হবে)
-
