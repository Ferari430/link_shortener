package auth

type LoginResponce struct { // Это то что отправляется в ответ
	Token string `json:"token"`
}

type LoginRequest struct { // Это то что приходит по апи
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
 
type RegisterRequest struct { // Это то что приходит по апи
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponce struct { // Это то что отправляется в ответ
	Token string `json:"token"`
}
