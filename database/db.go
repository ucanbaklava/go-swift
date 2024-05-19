package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	// Create users table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        role TEXT DEFAULT 'user',
        is_active BOOLEAN DEFAULT TRUE,
        last_login DATETIME
    );
	
	create table if not exists posts(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT not null unique,
		content TEXT not null,
		excerpt TEXT,
		slug TEXT not null,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_deleted BOOLEAN default false,
		user_id integer not null,
		FOREIGN KEY (user_id) references users(id)
	);

	create table if not exists comments(
		id integer PRIMARY KEY autoincrement,
		content text not null,
		user_id integer not null,
		post_id integer not null,
		parent_id integer,
		foreign key (user_id) references users(id),
		foreign key (post_id) references posts(id),
		foreign key (parent_id) references comments(id)
	);
	
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
		return nil, err
	}

	return db, nil
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = SetupDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
