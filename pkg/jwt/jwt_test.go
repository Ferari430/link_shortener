package jwt_test

import (
	"my_project/pkg/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {

	const email = "a123@a.ru"

	jwtService := jwt.NewJWT("1223123213")

	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatal("Cant create jwt token", err.Error())
	}

	isValid, jwtData := jwtService.ParseToken(token)

	if !isValid {
		t.Fatal("Cant parse jwt token")

	}

	if jwtData.Email != email {
		t.Fatalf("Email %s not equal test email  %s", jwtData.Email, email)
	}

}
