package rest

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}
