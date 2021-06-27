package main

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type userStorage interface {
	get(id int) (User, error)
	insert(u User) (int64, error)
	delete(id int) (int64, error)
}

func (storage *dataBaseStorage) get(id int) (User, error) {
	row := storage.db.QueryRow("select id, firstname, lastname, name, email from users where id = $1", id)

	user := User{}

	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Name, &user.Email)

	if err != nil {
		return user, errors.New("user not exist")
	}

	return user, nil
}

func (storage *dataBaseStorage) insert(u User) (int64, error) {
	result, err := storage.db.Exec(
		"insert into users (firstname, lastname, name, email) values ($1, $2, $3, $4)",
		u.Firstname, u.Lastname, u.Name, u.Email)

	if err != nil {
		return 0, nil
	}

	fmt.Print(result)

	return result.LastInsertId()
}

func (storage *dataBaseStorage) update(u User) error {
	return nil
}

func (storage *dataBaseStorage) delete(id int) (int64, error) {
	result, err := storage.db.Exec("delete from users where id = $1", id)

	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}
