package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func decodeJWT(token string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Некорректный формат JWT")
		return
	}

	// Раскодируем header
	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		fmt.Println("Ошибка декодирования header:", err)
		return
	}
	var header map[string]interface{}
	json.Unmarshal(headerJSON, &header)
	fmt.Println("Header:", header)

	// Раскодируем payload
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Ошибка декодирования payload:", err)
		return
	}
	var payload map[string]interface{}
	json.Unmarshal(payloadJSON, &payload)
	fmt.Println("Payload:", payload)

	// Подпись (не декодируем, просто покажем)
	fmt.Println("Signature (base64):", parts[2])
}

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImEzNTYzQGEucnUifQ.juUlZQ5GRZ-mJnXFAbuWfGosUwdMQsHndHEPTI8ioDk"
	decodeJWT(token)
}
