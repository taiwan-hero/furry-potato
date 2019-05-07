package heirarchy

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

//Heirarchy implements exported context for user role heirarchy structure
type Heirarchy struct {
	userMap map[int64]user
	roleMap map[int64][]user
	roleSub map[int64][]int64
}

//populates the heirarchy context with a slice of roles
func (h *Heirarchy) setRoles(r []role) {
	h.roleSub = make(map[int64][]int64)
	for i := 0; i < len(r); i++ {
		h.roleSub[r[i].Parent] = append(h.roleSub[r[i].Parent], r[i].ID)
	}
}

//populates the heirarchy context with a slice of users
func (h *Heirarchy) setUsers(u []user) {
	h.userMap = make(map[int64]user)
	for i := 0; i < len(u); i++ {
		h.userMap[u[i].ID] = u[i]
	}
	h.roleMap = make(map[int64][]user)
	for i := 0; i < len(u); i++ {
		h.roleMap[u[i].Role] = append(h.roleMap[u[i].Role], u[i])
	}
}

//given a userID, returns the user document
func (h *Heirarchy) getUserByID(userID int64) (user, bool) {
	u, ok := h.userMap[userID]
	return u, ok
}

//given a role ID, returns a slice of users with that role
func (h *Heirarchy) getUsersByRole(roleID int64) []user {
	out := h.roleMap[roleID]
	return out
}

//given a role ID, returns all role ID's subordinate to it (and its subordinates)
func (h *Heirarchy) getSubordinateRoles(roleID int64) []int64 {
	var out, todo []int64
	r, _ := h.roleSub[roleID]
	out = r
	todo = r
	//iterate through the todo queue for more subordinate roles
	for len(todo) > 0 {
		r0 := todo[0]
		todo = todo[1:]
		r1, ok := h.roleSub[r0]
		if ok {
			out = append(out, r1...)
			todo = append(todo, r1...)
		}
	}
	return out
}

//given a user ID, returns all users that have a subordinate role (and its subordinates)
func (h *Heirarchy) getSubordinates(userID int64) []user {
	var out []user

	u, ok := h.getUserByID(userID)
	if ok {
		subRoles := h.getSubordinateRoles(u.Role)
		for i := 0; i < len(subRoles); i++ {
			u := h.getUsersByRole(subRoles[i])
			out = append(out, u...)
		}
	}
	return out
}
