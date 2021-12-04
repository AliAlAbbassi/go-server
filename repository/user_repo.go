package repository

import (
	"container/list"
	"database/sql"
	"log"

	"github.com/alialabbassi/go-server/model"
	util "github.com/lakshanwd/db-helper/mysql"
)

//UserRepo - User repository
type UserRepo struct {
	Name string
}

//GetUserRepository - returns User repository
func GetUserRepository() UserRepo {
	return UserRepo{Name: "User"}
}

//Select - Select Users from db
func (repo UserRepo) Select() (*list.List, error) {
	reader := func(rows *sql.Rows, collection *list.List) {
		var User model.User
		err := rows.Scan(&User.UserID, &User.Name)
		collection.PushBack(User)
		log.Fatal(err)
	}
	return util.ExecuteReader(DbConnection, "select User_id, User_name, User_age, street_address, street_address_2, city, state, zip_code, country from tbl_User", reader)
}

//Insert - Insert User to db
func (repo UserRepo) Insert(doc interface{}) (int64, error) {
	User := doc.(model.User)
	return util.ExecuteInsert(DbConnection, "insert into tbl_User (user_name, user_password) values (?,?)", User.Name, User.Password)
}

//Update - Update User
func (repo UserRepo) Update(doc interface{}) (int64, error) {
	User := doc.(model.User)
	return util.ExecuteUpdateDelete(DbConnection, "update tbl_User set user_name=?, user_password=?", User.Name, User.Password)
}

//Remove - Delete User from db
func (repo UserRepo) Remove(doc interface{}) (int64, error) {
	User := doc.(model.User)
	return util.ExecuteUpdateDelete(DbConnection, "delete from tbl_User where User_id=?", User.UserID)
}
