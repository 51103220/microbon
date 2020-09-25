package authentication

type SugarUserFunctions map[string]string

type SugarUser struct {
	Id        string
	Username  string
	Password  string
	Secret    string
	PublicKey string
	Roles     SugarUserFunctions
}

func (user *SugarUser) AddRole(roles ...string) {
	if user.Roles == nil {
		user.Roles = make(map[string]string)
	}

	for _, role := range roles {
		user.Roles[role] = role
	}
}

func (user *SugarUser) HasRole(role string) bool {
	_, hasRole := user.Roles[role]

	return hasRole
}
