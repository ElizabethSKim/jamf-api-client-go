package classic

import (
	"time"
)

// ExpirationBuffer is the buffer duration added to the expiration check.
var ExpirationBuffer = 30 * time.Second

type AuthToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

func (t *AuthToken) IsExpired() (bool, error) {
	if t.Expires == "" {
		return true, nil
	}
	expiration, err := time.Parse(time.RFC3339, t.Expires)
	if err != nil {
		return true, err
	}
	// Add buffer to prevent race conditions where the token is valid
	// during this check but expires before the HTTP request is made
	return expiration.Before(time.Now().Add(ExpirationBuffer)), nil
}
