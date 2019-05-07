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
	roleMap map[int64]role
}

func (h *Heirarchy) setRoles(r []role) {
	h.roles = r
	h.roleMap = make(map[int64]role)
	for i := 0; i < len(r); i++ {
		h.roleMap[r[i].ID] = r[i]
	}
}

func (h *Heirarchy) setUsers(u []user) {
	h.users = u
	h.userMap = make(map[int64]user)
	for i := 0; i < len(u); i++ {
		h.userMap[u[i].ID] = u[i]
	}
}

func (h *Heirarchy) getUserByID(userID int64) (user, bool) {
	u, ok := h.userMap[userID]
	return u, ok
}

func (h *Heirarchy) getSubordinates(u int64) []user {
	return h.users
}
