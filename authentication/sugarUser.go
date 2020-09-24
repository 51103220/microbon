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

func (sf SugarUserFunctions) AddRole(roles ...string) SugarUserFunctions {
	if sf == nil {
		sf = make(map[string]string)
	}
	for _, role := range roles {
		sf[role] = role
	}

	return sf
}

func (sf SugarUserFunctions) HasRole(role string) bool {
	_, hasRole := sf[role]

	return hasRole
}
