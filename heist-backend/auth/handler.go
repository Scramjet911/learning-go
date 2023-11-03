package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"net/http"

	s "github.com/Scramjet911/learning-go/go-books/server"
	"github.com/Scramjet911/learning-go/go-books/util"
)

type AuthHandler struct {
	server *s.Server
}

func NewAuthHandler(server *s.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

func (a *AuthHandler) GetNameHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newAuth := AuthInput{}
	util.ParseBody(r, newAuth)

	h := sha256.New()
	h.Write([]byte(newAuth.TeamName))

	nameHash := h.Sum(nil)
	hashString := hex.EncodeToString(nameHash)

	hashRes := AuthResponse{Hash: hashString}

	res, _ := json.Marshal(hashRes)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
