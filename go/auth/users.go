package authmodels

import (
	"github.com/krishak-fiem/db/go/cassandra"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"log"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const tableName = "public.users"

var UserTable table.Table
var Session gocqlx.Session

func init() {
	session := cassandra.Init(9042)
	err := session.ExecStmt(`CREATE TABLE IF NOT EXISTS public.users (
		email text PRIMARY KEY,
		password text,
	)`)

	if err != nil {
		log.Fatalf(err.Error())
	}

	userMetadata := cassandra.CreateMetadata(tableName, []string{"email", "password"}, []string{"email"})
	userTable := cassandra.CreateTable(userMetadata)

	UserTable = *userTable
	Session = session
}

func (u *User) CreateUser() error {
	err := cassandra.InsertRow(Session, UserTable, u)
	if err != nil {
		return nil
	}
	return nil
}

func (u *User) GetUser() error {
	err := cassandra.GetRow(Session, UserTable, u)
	if err != nil {
		return nil
	}
	return nil
}
