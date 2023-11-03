package auth

type AuthInput struct {
	TeamName string `json:"teamName"`
}

type AuthResponse struct {
	Hash string `json:"hash"`
}
