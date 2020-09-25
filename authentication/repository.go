package authentication

import (
	"github.com/51103220/microbon/security"
)

var usersHolder = make(map[string]*SugarUser)
var authenticators = make(map[string]string)

func AddUsers(users ...*SugarUser) {
	for _, user := range users {
		usersHolder[user.Id] = user

		challenge := security.MakeBasicChallenge(user.Username, user.Password)

		authenticators[challenge] = user.Id
	}
}

func GetUser(id string) *SugarUser {
	user, found := usersHolder[id]
	if found {
		return user
	}

	return nil
}

func Authenticate(challenge string) *SugarUser {
	id, authenticated := authenticators[challenge]

	if authenticated {
		user, found := usersHolder[id]

		if found {
			return user
		}
	}

	return nil
}

func Authorize(user *SugarUser, role string) bool {
	return user.HasRole(role)
}
