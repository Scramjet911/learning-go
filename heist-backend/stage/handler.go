package stage

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/Scramjet911/learning-go/go-books/constant"
	s "github.com/Scramjet911/learning-go/go-books/server"
	"github.com/Scramjet911/learning-go/go-books/util"
)

type StageHandler struct {
	server *s.Server
}

func NewStageHandler(server *s.Server) *StageHandler {
	return &StageHandler{server: server}
}

func partialSha256(input string) string {
	hasher := sha256.New()

	hasher.Write([]byte(input))

	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	return hashString[:10]
}

func (s *StageHandler) FirstStage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	input := &FirstStageInput{}
	if !util.ParseBody(r, input) {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	specificValue := "HTC Dream"

	if input.Solution == specificValue {
		hashedValue := partialSha256(input.Solution)
		stagePassed := StageResponse{
			Hint:   "Different variants clue",
			Code:   hashedValue,
			Status: constant.Success,
		}

		res, _ := json.Marshal(stagePassed)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		stageFailed := ErrorResponse{
			Status:  constant.Failure,
			Message: util.GetRandomMeanResponse(),
		}
		res, _ := json.Marshal(stageFailed)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
	}
}
