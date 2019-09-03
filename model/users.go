package model

import (
	"database/sql"
	"fmt"
	"strings"
)

// Users need to explain
type Users struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	ID     int    `json:"id"`
}

// GetUsers will return an users with the given ID
func (m *Users) GetUsers(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT Name, Gender, Age, ID FROM users WHERE id=%d", m.ID)
	return db.QueryRow(statement).Scan(&m.Name, &m.Gender, &m.Age, &m.ID)
}

// UpdateUsers will update the users info with the given ID
func (m *Users) UpdateUsers(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET Name='%s', Gender='%s', Age=%d, ID=%d WHERE id=%d", m.Name, m.Gender, m.Age, m.ID, m.ID)
	_, err := db.Exec(statement)
	return err
}

// DeleteUsers will detele the users with the given id
func (m *Users) DeleteUsers(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", m.ID)
	_, err := db.Exec(statement)
	return err
}

// CreateUsers will create users with given info
func (m *Users) CreateUsers(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users (Name, Gender, Age, ID) VALUES('%s', '%s', '%d', '%d')", m.Name, m.Gender, m.Age, m.ID)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&m.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers will give all  the users info
func GetUsers(db *sql.DB, start, count int) ([]Users, error) {
	statement := fmt.Sprintf("SELECT id, Name, Gender, Age, ID FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := []Users{}
	for rows.Next() {
		var m Users
		if err := rows.Scan(&m.ID, &m.Name, &m.Gender, &m.Age, &m.ID); err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}

// DeleteUsers will delete given users with id from databse
func DeleteUsers(db *sql.DB, arrayUserID []Users) (err error) {
	var ids string
	for _, u := range arrayUserID {
		u.GetUsers(db)
		fmt.Println(u)
		ids = fmt.Sprintf("%v, %v", ids, u.ID)
	}

	ids = strings.Trim(ids, "! ,")
	fmt.Printf("> ids=%v", ids)
	statement := fmt.Sprintf("DELETE FROM users WHERE id in (%v)", ids)
	_, err = db.Exec(statement)
	return
}
