package auth_models

import (
	"github.com/krishak-fiem/db/go/cassandra"
	"github.com/scylladb/gocqlx/v2/table"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const tableName = "User"

var UserTable table.Table

func init() {
	userMetadata := cassandra.CreateMetadata(tableName, []string{"email", "password"}, []string{"email"}, []string{""})
	userTable := cassandra.CreateTable(userMetadata)

	UserTable = *userTable
}

func (u *User) CreateUser() error {
	err := cassandra.InsertRow(UserTable, &u)
	if err != nil {
		return nil
	}
	return nil
}

func (u *User) GetUser() error {
	err := cassandra.GetRow(UserTable, &u)
	if err != nil {
		return nil
	}
	return nil
}
