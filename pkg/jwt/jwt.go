package jwt

import (
	"fmt"
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
	claims := jwt.MapClaims{
		"email": data.Email, // Передаем строку, а не структуру или карту
	}

	log.Printf("Creating token for email: %v (type: %T)", data.Email, data.Email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	log.Println("[JWT] Parsing token:", token)

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("[JWT ERROR] Unexpected signing method:", t.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		log.Println("[JWT ERROR] Failed to parse token:", err)
		return false, nil
	}

	if !t.Valid {
		log.Println("[JWT ERROR] Token is not valid")
		return false, nil
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("[JWT ERROR] Failed to convert claims to MapClaims")
		return false, nil
	}

	log.Printf("[JWT] Claims parsed: %v", claims)

	emailRaw, exists := claims["email"]
	if !exists {
		log.Println("[JWT ERROR] 'email' field is missing in claims")
		return false, nil
	}

	emailMap, ok := emailRaw.(map[string]interface{})
	if !ok {
		log.Println("[JWT ERROR] 'email' field is not a map:", emailRaw)
		return false, nil
	}

	log.Println(emailMap)

	email, ok := emailMap["Email"].(string)
	if !ok {
		log.Println("[JWT ERROR] 'Email' field in map is not a string:", emailMap["Email"])
		return false, nil
	}

	if !ok {
		log.Println("[JWT ERROR] 'email' field is not a string:", emailRaw)
		return false, nil
	}

	log.Println("[JWT SUCCESS] Token is valid for email:", email)

	return true, &JWTData{
		Email: email,
	}
}
