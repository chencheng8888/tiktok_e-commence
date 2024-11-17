package token

import "testing"

func TestJwt(t *testing.T) {
	j := NewJWTer()
	var (
		UserID int32 = 111
		Secret       = "happy"
	)
	tokenString, err := j.GenerateJwtToken(UserID, Secret)
	if err != nil {
		t.Error(err)
	}
	res, err := j.VerifyJwtToken(tokenString, Secret)
	if err != nil {
		t.Error(err)
	}
	if res != UserID {
		t.Error("verify failed")
	}
}