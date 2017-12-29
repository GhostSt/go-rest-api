package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	"log"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserModel struct {
	*registry
}

func (c UserController) AddUser(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error  {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	defer  r.Body.Close()

	if err != nil {
		return err
	}

    err = json.Unmarshal(body, &user)

    if err != nil {
    	return err
	}

	userService := UserModel{registry: c.registry}

	if !userService.IsUniqueEmail(user) {
		c.registry.render.JSON(rw, 400, map[string]string{"message": "User exists"})

		return nil
	}

	userService.addUser(user)

	return nil
}

func (c UserController) getList(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	userModel := UserModel{registry: c.registry}

	list := userModel.getList()

	rw.Header().Add("Content-Type", "applicant/json111")
	json.NewEncoder(rw).Encode(list)

	return nil
}

func (c UserController) getUser(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	id := params.ByName("id")

	userModel := UserModel{registry: c.registry}

	user, err = userModel.getUser(id)

	if err != nil {

	}
}

// Checks if user email exist in database
func (u *UserModel) IsUniqueEmail(user User) bool  {
	err := u.registry.db.QueryRow("SELECT id FROM user WHERE email = ?", user.Email).Scan()

	if (err != nil) {
		return true
	}

	return false
}

// Adds user to database
func (u *UserModel) addUser(user User)  {
	db := u.registry.db

	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := transaction.Prepare("INSERT INTO user(name, email) VALUES(?, ?)")

	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	transaction.Commit()
}

// Gets list of users
func (u *UserModel) getList() []User {
	db := u.registry.db

	var users []User

	rows, err := db.Query("SELECT * FROM user")

	if (err != nil) {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id, &user.Name, &user.Email)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}

// Gets user by id
func (u *UserModel) getUser(id string) User {
	db := u.registry.db

	var user User

	row, err := db.Query("SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		panic(nil)
	}

	err = row.Scan(&user.Id, &user.Email, &user.Name)

	if err != nil {
		panic(err)
	}

	return User{}
}