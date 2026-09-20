package task

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAll() ([]Task, error) {
	rows, err := s.db.Query("SELECT id, title, done FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
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

func (s *Store) GetDeleted() ([]Task, error) {
	rows, err := s.db.Query(
		"SELECT id, title, done FROM tasks WHERE deleted_at IS NOT NULL ORDER BY deleted_at DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (s *Store) Restore(id int) error {
	result, err := s.db.Exec(
		"UPDATE tasks SET deleted_at = NULL WHERE id = ? AND deleted_at IS NOT NULL",
		id,
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
