package db

import (
	"database/sql"
	"fmt"

	"github.com/ThomasHamilton2/todo-list-6/schema"
	_ "github.com/go-sql-driver/mysql" //handles MySQL database connection
)

//MySQL has a pointer to a MySQL database object
type MySQL struct {
	DB *sql.DB
}

//Close closes the database connection
func (m *MySQL) Close() {
	m.DB.Close()
}

//Insert inserts a new Todo object into the database
func (m *MySQL) Insert(todo *schema.Todo) (int, error) {
	res, err := m.DB.Exec(`INSERT INTO Todo (Title, Complete) VALUES (?,?)`, todo.Title, todo.Complete)

	if err != nil {
		println("Exec err:", err.Error())
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		println("Error:", err.Error())
		return -1, err
	}

	println("LastInsertId:", id)
	return int(id), nil
}

//Update updatesTodo object
func (m *MySQL) Update(todo *schema.Todo) error {
	var stmt = `UPDATE Todo SET Title = ?, Complete = ? WHERE ID = ?`
	_, err := m.DB.Exec(stmt, todo.Title, todo.Complete, todo.ID)

	if err != nil {
		println("Exec err:", err.Error())
		return err
	}

	return nil
}

//Delete deletes Todo object from database
func (m *MySQL) Delete(id int) error {
	var stmt = `DELETE FROM Todo WHERE ID = ?`
	if _, err := m.DB.Exec(stmt, id); err != nil {
		println("Exec err:", err.Error())
		return err
	}

	return nil
}

//GetAll gets all Todo objects from database
func (m *MySQL) GetAll() ([]schema.Todo, error) {
	query := `
        SELECT *
        FROM Todo
        ORDER BY ID;
    `

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var todoList []schema.Todo
	for rows.Next() {
		var t schema.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Complete); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}

	return todoList, nil
}

//ConnectMySQL connects to the MySQL database
func ConnectMySQL() (*MySQL, error) {
	fmt.Println("made it to ConnectMySQL")
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/Todo_db")
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to mySql")

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("pinged mySql")

	return &MySQL{db}, nil
}
