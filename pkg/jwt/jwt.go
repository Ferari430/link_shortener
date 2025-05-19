package jwt

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})

	log.Printf("Creating token for email: %v (type: %T)", data.Email, data.Email)

	token, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return token, nil

}

func (j *JWT) ParseToken(token string) (bool, *JWTData) {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return false, nil
	}

	email := t.Claims.(jwt.MapClaims)["email"]
	return t.Valid, &JWTData{
		Email: email.(string),
	}

}
