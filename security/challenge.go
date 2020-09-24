package security

import (
	b64 "encoding/base64"
	"fmt"
)

func MakeBasicChallenge(username string, password string) string {
	secret := fmt.Sprintf("%s:%s", username, password)
	challenge := b64.StdEncoding.EncodeToString([]byte(secret))

	return challenge
}
