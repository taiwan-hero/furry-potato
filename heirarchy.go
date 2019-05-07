package heirarchy

import "fmt"

type role struct {
	ID     int64  `json:"Id"`
	Name   string `json:"Name"`
	Parent int64  `json:"Parent"`
}

type user struct {
	ID   int64  `json:"Id"`
	Name string `json:"Name"`
	Role int64  `json:"Role"`
}

type Heirarchy struct {
}

func setRoles(r []role) {
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d: %s %d\n", r[i].ID, r[i].Name, r[i].Parent)
	}
}

func setUsers(u []user) {
	for i := 0; i < len(u); i++ {
		fmt.Printf("%d: %s %d\n", u[i].ID, u[i].Name, u[i].Role)
	}
}

func getSubordinates(u int64) []user {

}
