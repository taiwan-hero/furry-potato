package heirarchy

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
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

func TestGetUserByID(t *testing.T) {
	h := Heirarchy{}
	u := parseUsers("./users1.json")
	h.setUsers(u)
	u1, exists := h.getUserByID(5)
	if !exists {
		t.Errorf("u1: %s \n", u1.Name)
	}
	expected := user{5, "Steve Trainer", 5}
	if !reflect.DeepEqual(u1, expected) {
		t.Errorf("got %v ; expected %v", u1, expected)
	}
}

func TestGetSubordinateRoles(t *testing.T) {
	h := Heirarchy{}
	r := parseRoles("./roles1.json")
	h.setRoles(r)

	r1 := h.getSubordinateRoles(2)

	expected := []int64{3, 4, 5}
	if !reflect.DeepEqual(r1, expected) {
		t.Errorf("got %v ; expected %v", r1, expected)
	}
}

func TestGetUsersByRole(t *testing.T) {
	h := Heirarchy{}
	r := parseRoles("./roles1.json")
	h.setRoles(r)
	u := parseUsers("./users1.json")
	h.setUsers(u)

	actual := h.getUsersByRole(5)
	expected := []user{{5, "Steve Trainer", 5}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v ; expected %v", actual, expected)
	}
}
func TestGetSubordinates(t *testing.T) {
	h := Heirarchy{}
	r := parseRoles("./roles1.json")
	h.setRoles(r)
	u := parseUsers("./users1.json")
	h.setUsers(u)

	actual := h.getSubordinates(3)
	expected := []user{{2, "Emily Employee", 4}, {5, "Steve Trainer", 5}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v ; expected %v", actual, expected)
	}
}
