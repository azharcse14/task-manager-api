package handler_test

import (
	"database/sql"
	"encoding/json/v2"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/azharcse14/task-manager-api/internal/handler"
	"github.com/azharcse14/task-manager-api/internal/task"
)

// ---------- সাহায্যকারী ফাংশন ----------

// newTestServer: নতুন খালি ডাটাবেস দিয়ে পুরো router বানায়।
// main.go যা করে প্রায় তাই, শুধু ফাইলের বদলে :memory:, আর ListenAndServe নেই।
func newTestServer(t *testing.T) http.Handler {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("ডাটাবেস খোলা গেল না: %v", err)
	}
	db.SetMaxOpenConns(1)
	t.Cleanup(func() { db.Close() })

	if err := task.CreateSchema(db); err != nil {
		t.Fatalf("table বানানো গেল না: %v", err)
	}

	h := handler.NewTaskHandler(task.NewStore(db))
	return h.Routes()
}

// doRequest: একটা নকল request বানিয়ে router-এর হাতে দেয়, আর recorder ফেরত দেয়
func doRequest(t *testing.T, srv http.Handler, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(method, path, strings.NewReader(body)) // নকল চিঠি
	rec := httptest.NewRecorder()                                     // টেপ রেকর্ডার
	srv.ServeHTTP(rec, req)                                           // চিঠিটা router-কে দাও
	return rec
}

// decodeJSON: response-এর body থেকে JSON পড়ে v-তে রাখে
func decodeJSON(t *testing.T, rec *httptest.ResponseRecorder, v any) {
	t.Helper()

	if err := json.Unmarshal(rec.Body.Bytes(), v); err != nil {
		t.Fatalf("response JSON পড়া গেল না: %v\nbody ছিল: %s", err, rec.Body.String())
	}
}

// createViaAPI: POST /tasks দিয়ে একটা task বানায় (প্রস্তুতির জন্য)
func createViaAPI(t *testing.T, srv http.Handler, body string) task.Task {
	t.Helper()

	rec := doRequest(t, srv, "POST", "/tasks", body)
	if rec.Code != http.StatusCreated {
		t.Fatalf("task বানানো গেল না: status %d, body %s", rec.Code, rec.Body.String())
	}
	var tk task.Task
	decodeJSON(t, rec, &tk)
	return tk
}

// বাইরের একজন client-এর চোখে response দেখতে কেমন।
// handler-এর listResponse আর pageInfo ছোট হাতের (unexported), তাই এখান থেকে দেখা যায় না।
// তাই নিজেদের মতো করে বানিয়ে নিচ্ছি, ঠিক যেমন একজন client বানাতো।
type meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type listBody struct {
	Data []task.Task `json:"data"`
	Meta meta        `json:"meta"`
}

type errorBody struct {
	Error string `json:"error"`
}

// ---------- POST /tasks ----------

// গল্প: ঠিকঠাক task পাঠালে 201, সঠিক header, আর নতুন task ফেরত আসবে
func TestCreateTask(t *testing.T) {
	srv := newTestServer(t)

	rec := doRequest(t, srv, "POST", "/tasks", `{"title":"Learn httptest","done":false}`)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, আশা করেছিলাম %d", rec.Code, http.StatusCreated)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, আশা করেছিলাম application/json", got)
	}
	if got := rec.Header().Get("Location"); got != "/tasks/1" {
		t.Errorf("Location = %q, আশা করেছিলাম /tasks/1", got)
	}

	var got task.Task
	decodeJSON(t, rec, &got)
	want := task.Task{ID: 1, Title: "Learn httptest", Done: false}
	if got != want {
		t.Errorf("body = %+v, আশা করেছিলাম %+v", got, want)
	}
}

// ভুল body পাঠালে 400 আর একটা error message আসা উচিত
func TestCreateTaskBadInput(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{name: "JSON-ই না", body: `hello`},
		{name: "অর্ধেক JSON", body: `{"title":`},
		{name: "খালি body", body: ``},
		{name: "title নেই", body: `{"done":true}`},
		{name: "title খালি", body: `{"title":""}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newTestServer(t)

			rec := doRequest(t, srv, "POST", "/tasks", tt.body)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, আশা করেছিলাম %d", rec.Code, http.StatusBadRequest)
			}
			var e errorBody
			decodeJSON(t, rec, &e)
			if e.Error == "" {
				t.Error("400-এর সাথে একটা error message থাকার কথা")
			}
		})
	}
}

// ---------- GET /tasks/{id} ----------

func TestGetTask(t *testing.T) {
	srv := newTestServer(t)
	created := createViaAPI(t, srv, `{"title":"Find me","done":true}`)

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{name: "আছে এমন task", path: "/tasks/1", wantStatus: http.StatusOK},
		{name: "নেই এমন task", path: "/tasks/999", wantStatus: http.StatusNotFound},
		{name: "ID সংখ্যা না", path: "/tasks/abc", wantStatus: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, srv, "GET", tt.path, "")
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, আশা করেছিলাম %d", rec.Code, tt.wantStatus)
			}
		})
	}

	// 200-এর বেলায় শুধু status না, body-ও মিলিয়ে দেখি
	rec := doRequest(t, srv, "GET", "/tasks/1", "")
	var got task.Task
	decodeJSON(t, rec, &got)
	if got != created {
		t.Errorf("GET body = %+v, আশা করেছিলাম %+v", got, created)
	}
}

// ---------- PUT /tasks/{id} ----------

// গল্প: বানাও, PUT দিয়ে বদলাও, তারপর GET দিয়ে দেখো সত্যিই বদলেছে কিনা
func TestUpdateTask(t *testing.T) {
	srv := newTestServer(t)
	createViaAPI(t, srv, `{"title":"Old","done":false}`)

	rec := doRequest(t, srv, "PUT", "/tasks/1", `{"title":"New","done":true}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("PUT status = %d, আশা করেছিলাম %d", rec.Code, http.StatusOK)
	}

	rec = doRequest(t, srv, "GET", "/tasks/1", "")
	var got task.Task
	decodeJSON(t, rec, &got)
	want := task.Task{ID: 1, Title: "New", Done: true}
	if got != want {
		t.Errorf("PUT-এর পর GET = %+v, আশা করেছিলাম %+v", got, want)
	}
}

func TestUpdateTaskErrors(t *testing.T) {
	srv := newTestServer(t)
	createViaAPI(t, srv, `{"title":"Exists"}`)

	tests := []struct {
		name       string
		path       string
		body       string
		wantStatus int
	}{
		{name: "নেই এমন task", path: "/tasks/999", body: `{"title":"X"}`, wantStatus: http.StatusNotFound},
		{name: "ID সংখ্যা না", path: "/tasks/abc", body: `{"title":"X"}`, wantStatus: http.StatusBadRequest},
		{name: "title খালি", path: "/tasks/1", body: `{"title":""}`, wantStatus: http.StatusBadRequest},
		{name: "JSON ভুল", path: "/tasks/1", body: `oops`, wantStatus: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, srv, "PUT", tt.path, tt.body)
			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, আশা করেছিলাম %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

// ---------- DELETE /tasks/{id} ----------

// গল্প: বানাও, মুছে ফেলো, তারপর আর খুঁজে পাওয়ার কথা না
func TestDeleteTask(t *testing.T) {
	srv := newTestServer(t)
	createViaAPI(t, srv, `{"title":"Delete me"}`)

	rec := doRequest(t, srv, "DELETE", "/tasks/1", "")
	if rec.Code != http.StatusNoContent {
		t.Fatalf("DELETE status = %d, আশা করেছিলাম %d", rec.Code, http.StatusNoContent)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("204-এর body খালি থাকার কথা, পেলাম %q", rec.Body.String())
	}

	if rec := doRequest(t, srv, "GET", "/tasks/1", ""); rec.Code != http.StatusNotFound {
		t.Errorf("মোছার পর GET status = %d, আশা করেছিলাম 404", rec.Code)
	}
	if rec := doRequest(t, srv, "DELETE", "/tasks/1", ""); rec.Code != http.StatusNotFound {
		t.Errorf("দ্বিতীয়বার DELETE status = %d, আশা করেছিলাম 404", rec.Code)
	}
}

// ---------- GET /tasks ----------

func TestListTasks(t *testing.T) {
	srv := newTestServer(t)
	createViaAPI(t, srv, `{"title":"A","done":true}`)
	createViaAPI(t, srv, `{"title":"B","done":false}`)
	createViaAPI(t, srv, `{"title":"C","done":false}`)

	tests := []struct {
		name     string
		path     string
		wantLen  int  // এই পাতায় কয়টা task আসবে
		wantMeta meta // meta অংশে কী থাকবে
	}{
		{name: "কিছু না দিলে default", path: "/tasks", wantLen: 3,
			wantMeta: meta{Page: 1, PerPage: 10, Total: 3, TotalPages: 1}},
		{name: "পাতায় ২টা করে", path: "/tasks?per_page=2", wantLen: 2,
			wantMeta: meta{Page: 1, PerPage: 2, Total: 3, TotalPages: 2}},
		{name: "দ্বিতীয় পাতা", path: "/tasks?per_page=2&page=2", wantLen: 1,
			wantMeta: meta{Page: 2, PerPage: 2, Total: 3, TotalPages: 2}},
		{name: "শুধু শেষ হওয়া", path: "/tasks?done=true", wantLen: 1,
			wantMeta: meta{Page: 1, PerPage: 10, Total: 1, TotalPages: 1}},
		{name: "page=0 দিলে 1 ধরে", path: "/tasks?page=0", wantLen: 3,
			wantMeta: meta{Page: 1, PerPage: 10, Total: 3, TotalPages: 1}},
		{name: "per_page বেশি দিলে 100-এ আটকায়", path: "/tasks?per_page=500", wantLen: 3,
			wantMeta: meta{Page: 1, PerPage: 100, Total: 3, TotalPages: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, srv, "GET", tt.path, "")

			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, আশা করেছিলাম %d", rec.Code, http.StatusOK)
			}

			wantCount := strconv.Itoa(tt.wantMeta.Total)
			if got := rec.Header().Get("X-Total-Count"); got != wantCount {
				t.Errorf("X-Total-Count = %q, আশা করেছিলাম %q", got, wantCount)
			}

			var body listBody
			decodeJSON(t, rec, &body)
			if len(body.Data) != tt.wantLen {
				t.Errorf("data-তে %d টা task, আশা করেছিলাম %d টা", len(body.Data), tt.wantLen)
			}
			if body.Meta != tt.wantMeta {
				t.Errorf("meta = %+v, আশা করেছিলাম %+v", body.Meta, tt.wantMeta)
			}
		})
	}
}

func TestListTasksBadQuery(t *testing.T) {
	srv := newTestServer(t)

	tests := []struct {
		name string
		path string
	}{
		{name: "page অক্ষর", path: "/tasks?page=abc"},
		{name: "per_page অক্ষর", path: "/tasks?per_page=x"},
		{name: "done ভুল", path: "/tasks?done=maybe"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, srv, "GET", tt.path, "")
			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, আশা করেছিলাম %d", rec.Code, http.StatusBadRequest)
			}
		})
	}
}

// ---------- Router ----------

// এই আচরণের কোড তুমি লেখোনি, Go-র router নিজেই দেয়: ভুল method দিলে 405
func TestMethodNotAllowed(t *testing.T) {
	srv := newTestServer(t)

	rec := doRequest(t, srv, "PATCH", "/tasks/1", `{"title":"X"}`)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, আশা করেছিলাম %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
