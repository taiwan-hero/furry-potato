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

type Heirarchy struct {
	roles   []role
	users   []user
	userMap map[int64]user
	roleMap map[int64][]user
	roleSub map[int64][]int64
}

func (h *Heirarchy) setRoles(r []role) {
	h.roles = r
	h.roleSub = make(map[int64][]int64)
	for i := 0; i < len(r); i++ {
		h.roleSub[r[i].Parent] = append(h.roleSub[r[i].Parent], r[i].ID)
	}
}

func (h *Heirarchy) setUsers(u []user) {
	h.users = u
	h.userMap = make(map[int64]user)
	for i := 0; i < len(u); i++ {
		h.userMap[u[i].ID] = u[i]
	}
	h.roleMap = make(map[int64][]user)
	for i := 0; i < len(u); i++ {
		h.roleMap[u[i].Role] = append(h.roleMap[u[i].Role], u[i])
	}
}

func (h *Heirarchy) getUserByID(userID int64) (user, bool) {
	u, ok := h.userMap[userID]
	return u, ok
}

func (h *Heirarchy) getUsersByRole(roleID int64) []user {
	out := h.roleMap[roleID]
	return out
}

func (h *Heirarchy) getSubordinateRoles(roleID int64) []int64 {
	var out, todo []int64
	r, _ := h.roleSub[roleID]
	out = r
	todo = r
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
