package heirarchy

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type roles struct {
	Roles []role `json:"Roles"`
}

type users struct {
	Users []user `json:"Users"`
}

func parseRoles(fileName string) []role {
	data, _ := ioutil.ReadFile(fileName)

	var r roles

	json.Unmarshal(data, &r)

	return r.Roles
}

func parseUsers(fileName string) []user {
	data, _ := ioutil.ReadFile(fileName)

	var u users

	json.Unmarshal(data, &u)

	return u.Users
}

func TestGetSubordinates(t *testing.T) {
	r := parseRoles("./roles1.json")

	setRoles(r)

	u := parseUsers("./users1.json")

	setUsers(u)
}
