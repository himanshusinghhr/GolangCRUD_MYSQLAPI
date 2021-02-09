package model

import (
	"github.com/todoapi/view"
)

func ReadAll() ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * from todo")
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)

	}
	return todos, nil

}
func ReadByName(name string) ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * from todo where name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)

	}
	return todos, nil

}
