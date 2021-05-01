package main

import (
	"database/sql"
	"fmt"
)

// struktur tabel
type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// menampilkan data untuk single row
func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name FROM user WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name)
}

// memperbaharui data
func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE user SET name='%s' WHERE id=%d", u.Name, u.ID)
	_, err := db.Exec(statement)
	return err
}

// menghapus data
func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM user WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

// membuat data baru
func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO user(name) VALUES('%s')", u.Name)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

// menampilkan data dalam multiple rows
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT id, name FROM user LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
