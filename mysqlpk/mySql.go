package mysqlpk

import (
	"database/sql"
	"fmt"

	"github.com/ThomasHamilton2/todo-list/schema"
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
	query := `
        INSERT INTO Todo (Title, Complete)
        VALUES (?, ?)
        RETURNING ID;
    `

	rows, err := m.DB.Query(query, todo.Title, todo.Complete)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

//Delete deletes Todo object from database
func (m *MySQL) Delete(id int) error {
	query := `
        DELETE FROM Todo
        WHERE ID = ?;
    `

	if _, err := m.DB.Exec(query, id); err != nil {
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
		// panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	fmt.Println("connected to mySql")

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("pinged mySql")

	return &MySQL{db}, nil
}
