package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
)

type Auth struct {
	*jwtauth.JWTAuth
	devMode bool
}

// todo: bu fonksiyon ve struct aslında gereksiz !
func NewAuth(secret string, devMode bool) *Auth {
	return &Auth{jwtauth.New("HS256", []byte(secret), nil), devMode}
}

type TokenClaims struct {
	UserID       int64  `json:"userID"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserFullName string `json:"userFullName"`
}

func TokenClaimsFromRequest(r *http.Request) *TokenClaims {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tc := TokenClaims{
		UserID:       int64(claims["userID"].(float64)),
		UserName:     claims["userName"].(string),
		UserEmail:    claims["userEmail"].(string),
		UserFullName: claims["userFullName"].(string),
	}

	return &tc
}

func NewTokenClaims() *TokenClaims {
	t := &TokenClaims{}

	return t
}

func NewTokenClaimsForUser(userID int64, userName string, userEmail string, userFullName string) (*TokenClaims, error) {
	t := NewTokenClaims()

	t.UserID = userID
	t.UserName = userName
	t.UserEmail = userEmail
	t.UserFullName = userFullName

	return t, nil
}

func (tc *TokenClaims) asMapClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"userID":       tc.UserID,
		"userName":     tc.UserName,
		"userEmail":    tc.UserEmail,
		"userFullName": tc.UserFullName,
	}
}

func (tc *TokenClaims) Encode(t *jwtauth.JWTAuth) string {
	mc := tc.asMapClaims()
	_, tokenString, _ := t.Encode(mc)
	return tokenString
}
