package task_test

import (
	"context"
	"database/sql"
	"errors"
	"slices"
	"testing"

	_ "modernc.org/sqlite" // test main.go চালায় না, তাই driver এখানেও import করতে হয়

	"github.com/azharcse14/task-manager-api/internal/task"
)

// newTestStore: প্রতিটা test-এর জন্য একদম নতুন, খালি ডাটাবেস বানায়।
// ডাটাবেসটা থাকে শুধু RAM-এ (:memory:), বন্ধ করলেই মুছে যায়।
func newTestStore(t *testing.T) *task.Store {
	t.Helper() // ভুল হলে এই লাইন না দেখিয়ে, যে test ডেকেছে তার লাইন দেখাও

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("ডাটাবেস খোলা গেল না: %v", err)
	}

	// :memory:-এ প্রতিটা connection আলাদা খাতা পায়।
	// তাই একটাই connection রাখছি, যেন সবাই একই খাতায় লেখে আর পড়ে।
	db.SetMaxOpenConns(1)

	// test শেষ হলে (pass হোক বা fail) নিজে থেকেই বন্ধ হবে
	t.Cleanup(func() { db.Close() })

	if err := task.CreateSchema(db); err != nil {
		t.Fatalf("table বানানো গেল না: %v", err)
	}
	return task.NewStore(db)
}

// createTask: test-এর প্রস্তুতির জন্য তাড়াতাড়ি একটা task বানিয়ে দেয়
func createTask(t *testing.T, s *task.Store, title string, done bool) task.Task {
	t.Helper()

	tk := task.Task{Title: title, Done: done}
	if err := s.Create(&tk); err != nil {
		t.Fatalf("task বানানো গেল না: %v", err)
	}
	return tk
}

// গল্প: একটা task বানাও, তারপর ID দিয়ে খুঁজে আনো, একই জিনিস ফেরত আসা উচিত
func TestCreateAndGetByID(t *testing.T) {
	s := newTestStore(t)

	created := createTask(t, s, "Learn testing", true)

	if created.ID == 0 {
		t.Fatal("Create-এর পর ID সেট হওয়ার কথা, কিন্তু ID এখনো 0")
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID error: %v", err)
	}
	if got != created {
		t.Errorf("GetByID() = %+v, আশা করেছিলাম %+v", got, created)
	}
}

// গল্প: যে task নেই, সেটা খুঁজলে "পাওয়া যায়নি" বলা উচিত
func TestGetByIDNotFound(t *testing.T) {
	s := newTestStore(t)

	_, err := s.GetByID(999)

	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("sql.ErrNoRows আশা করেছিলাম, পেলাম: %v", err)
	}
}

// গল্প: task বানাও, বদলাও, আবার খুঁজে আনো, নতুন মান দেখা উচিত
func TestUpdate(t *testing.T) {
	s := newTestStore(t)
	created := createTask(t, s, "Old title", false)

	updated := task.Task{ID: created.ID, Title: "New title", Done: true}
	if err := s.Update(updated); err != nil {
		t.Fatalf("Update error: %v", err)
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID error: %v", err)
	}
	if got != updated {
		t.Errorf("Update-এর পর পেলাম %+v, আশা করেছিলাম %+v", got, updated)
	}
}

// গল্প: যে task নেই, সেটা বদলাতে চাইলে "পাওয়া যায়নি" বলা উচিত
func TestUpdateNotFound(t *testing.T) {
	s := newTestStore(t)

	err := s.Update(task.Task{ID: 999, Title: "Ghost"})

	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("sql.ErrNoRows আশা করেছিলাম, পেলাম: %v", err)
	}
}

// গল্প: task বানাও, মুছে ফেলো, তারপর আর খুঁজে পাওয়ার কথা না
func TestDelete(t *testing.T) {
	s := newTestStore(t)
	created := createTask(t, s, "Delete me", false)

	if err := s.Delete(created.ID); err != nil {
		t.Fatalf("Delete error: %v", err)
	}

	// মোছার পর আর খুঁজে পাওয়ার কথা না
	if _, err := s.GetByID(created.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("মোছার পরও task পাওয়া গেল, err = %v", err)
	}

	// একই task দ্বিতীয়বার মুছলে "পাওয়া যায়নি" আসা উচিত
	if err := s.Delete(created.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("দ্বিতীয়বার মোছায় sql.ErrNoRows আশা করেছিলাম, পেলাম: %v", err)
	}
}

// List: একই প্রশ্ন ("এই শর্তে কোন task আসবে?") অনেক রকমভাবে,
// তাই এখানে table-driven test
func TestList(t *testing.T) {
	s := newTestStore(t)

	// ৫টা task: ১ আর ৩ শেষ হয়েছে, ২, ৪, ৫ বাকি
	createTask(t, s, "Task 1", true)
	createTask(t, s, "Task 2", false)
	createTask(t, s, "Task 3", true)
	createTask(t, s, "Task 4", false)
	createTask(t, s, "Task 5", false)

	yes, no := true, false

	tests := []struct {
		name      string
		filter    task.ListFilter
		wantTotal int   // শর্ত মেনে মোট কয়টা task আছে
		wantIDs   []int // এই পাতায় কোন কোন task আসবে, এই ক্রমে
	}{
		{name: "প্রথম পাতা", filter: task.ListFilter{Limit: 2, Offset: 0}, wantTotal: 5, wantIDs: []int{1, 2}},
		{name: "দ্বিতীয় পাতা", filter: task.ListFilter{Limit: 2, Offset: 2}, wantTotal: 5, wantIDs: []int{3, 4}},
		{name: "শেষ পাতায় একটাই", filter: task.ListFilter{Limit: 2, Offset: 4}, wantTotal: 5, wantIDs: []int{5}},
		{name: "পাতার বাইরে গেলে খালি", filter: task.ListFilter{Limit: 2, Offset: 10}, wantTotal: 5, wantIDs: []int{}},
		{name: "শুধু শেষ হওয়া", filter: task.ListFilter{Done: &yes, Limit: 10}, wantTotal: 2, wantIDs: []int{1, 3}},
		{name: "শুধু বাকি থাকা", filter: task.ListFilter{Done: &no, Limit: 10}, wantTotal: 3, wantIDs: []int{2, 4, 5}},
		{name: "filter আর পাতা একসাথে", filter: task.ListFilter{Done: &no, Limit: 2, Offset: 2}, wantTotal: 3, wantIDs: []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, total, err := s.List(context.Background(), tt.filter)
			if err != nil {
				t.Fatalf("List error: %v", err)
			}

			if total != tt.wantTotal {
				t.Errorf("total = %d, আশা করেছিলাম %d", total, tt.wantTotal)
			}

			// পুরো task না মিলিয়ে শুধু ID-গুলো মিলাচ্ছি, এতে পড়তে সহজ
			gotIDs := []int{}
			for _, tk := range got {
				gotIDs = append(gotIDs, tk.ID)
			}
			if !slices.Equal(gotIDs, tt.wantIDs) {
				t.Errorf("IDs = %v, আশা করেছিলাম %v", gotIDs, tt.wantIDs)
			}
		})
	}
}
