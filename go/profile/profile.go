package profilemodels

import (
	"github.com/krishak-fiem/db/go/cassandra"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"log"
)

type Profile struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	StreetAddress string `json:"streetAddress"`
	Pincode       uint32 `json:"pincode"`
	State         string `json:"state"`
	City          string `json:"city"`
	PhoneNumber   string `json:"phoneNumber"`
}

const tableName = "public.profiles"

var ProfileTable table.Table
var Session gocqlx.Session

func init() {
	session := cassandra.Init(9043)
	err := session.ExecStmt(`CREATE TABLE IF NOT EXISTS public.profiles (
		name text,
		email text PRIMARY KEY,
		street_address text,
		phone_number text,
		pincode int,
		city text,
		state text,
	)`)

	if err != nil {
		log.Fatalf(err.Error())
	}

	profileMetadata := cassandra.CreateMetadata(tableName, []string{"email", "name", "street_address", "phone_number", "pincode", "city", "state"}, []string{"email"})
	profileTable := cassandra.CreateTable(profileMetadata)

	ProfileTable = *profileTable
	Session = session
}

func (p *Profile) CreateProfile() error {
	err := cassandra.InsertRow(Session, ProfileTable, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) GetProfile() error {
	err := cassandra.GetRow(Session, ProfileTable, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateProfile() error {
	err := cassandra.InsertRow(Session, ProfileTable, p)
	if err != nil {
		return err
	}
	return nil
}
