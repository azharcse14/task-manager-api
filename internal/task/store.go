package task

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// ListFilter: task-এর লিস্ট চাওয়ার সময় কী কী শর্ত আছে
type ListFilter struct {
	Done   *bool // nil মানে "done দিয়ে বাছাই করবো না"
	Limit  int   // একবারে কয়টা task দেবো
	Offset int   // শুরু থেকে কয়টা task বাদ দেবো
}

// List: শর্ত মেনে এক পাতার task, আর মোট কয়টা task আছে সেটা ফেরত দেয়
func (s *Store) List(ctx context.Context, f ListFilter) ([]Task, int, error) {
	where := ""
	args := []any{}
	if f.Done != nil {
		where = " WHERE done = ?"
		args = append(args, *f.Done)
	}

	// ১. মোট কয়টা task আছে গুনে নাও
	var total int
	err := s.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM tasks"+where, args...,
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// ২. শুধু এই পাতার task গুলো আনো
	args = append(args, f.Limit, f.Offset)
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, title, done FROM tasks"+where+" ORDER BY id LIMIT ? OFFSET ?",
		args...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, t)
	}
	return tasks, total, rows.Err()
}

func (s *Store) GetByID(id int) (Task, error) {
	var t Task
	err := s.db.QueryRow(
		"SELECT id, title, done FROM tasks WHERE id = ?", id,
	).Scan(&t.ID, &t.Title, &t.Done)
	return t, err
}

func (s *Store) Create(t *Task) error {
	result, err := s.db.Exec(
		"INSERT INTO tasks (title, done) VALUES (?, ?)",
		t.Title, t.Done,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = int(id)
	return nil
}

func (s *Store) Update(t Task) error {
	result, err := s.db.Exec(
		"UPDATE tasks SET title = ?, done = ? WHERE id = ?",
		t.Title, t.Done, t.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
