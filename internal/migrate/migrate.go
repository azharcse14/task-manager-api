package migrate

import (
	"database/sql"
	"log"
)

type Migration struct {
	Name string
	SQL  string
}

const createMigrationsTable = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	name       TEXT     PRIMARY KEY,
	applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

var migrations = []Migration{
	{
		Name: "001_create_tasks_table",
		SQL: `
		CREATE TABLE IF NOT EXISTS tasks (
			id    INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT    NOT NULL,
			done  BOOLEAN NOT NULL DEFAULT 0
		)`,
	},
	{
		Name: "002_add_deleted_at_to_tasks",
		SQL:  `ALTER TABLE tasks ADD COLUMN deleted_at DATETIME DEFAULT NULL`,
	},
}

func Run(db *sql.DB) error {
	if _, err := db.Exec(createMigrationsTable); err != nil {
		return err
	}

	for _, m := range migrations {
		var count int
		err := db.QueryRow(
			"SELECT COUNT(*) FROM schema_migrations WHERE name = ?",
			m.Name,
		).Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		if _, err := db.Exec(m.SQL); err != nil {
			return err
		}

		if _, err := db.Exec(
			"INSERT INTO schema_migrations (name) VALUES (?)",
			m.Name,
		); err != nil {
			return err
		}

		log.Println("migration applied:", m.Name)
	}

	return nil
}
