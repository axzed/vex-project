package jwts

import "testing"

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzczMTM5ODksInRva2VuIjoiMTAwMCJ9.Eh_27dHot2ZBXbhwpb1ycPBxLcTxyWQ_hhavfQcUdYw"
	ParseToken(tokenString, "access_secret")
}
