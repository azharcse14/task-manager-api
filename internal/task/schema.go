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
