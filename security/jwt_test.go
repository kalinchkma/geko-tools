package gekotools

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func Test_jwtBuilder(t *testing.T) {
	newJwtBuilder := NewJwtBuilder([]byte("super secret"))

	newToken, err := newJwtBuilder.CreateToken(jwt.SigningMethodHS256)
	if err != nil {
		t.Fatal(err)

	}
	fmt.Println(newToken)
	t.Log("success")
}
