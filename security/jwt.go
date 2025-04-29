package gekotools

import "github.com/golang-jwt/jwt/v5"

type JwtBuilder interface {
	CreateToken(sigM jwt.SigningMethod) (string, error)
	CreateTokenWithClaims(sigM jwt.SigningMethod) (string, error)
}

type jwtBuilder struct {
	secret []byte
}

// Constructor
func NewJwtBuilder(secret []byte) *jwtBuilder {
	return &jwtBuilder{
		secret: secret,
	}
}

// Create token
func (jb *jwtBuilder) CreateToken(sigM jwt.SigningMethod) (string, error) {
	newToken := jwt.New(sigM)
	signToken, err := newToken.SignedString(jb.secret)
	if err != nil {
		return "", err
	}
	return signToken, nil
}
