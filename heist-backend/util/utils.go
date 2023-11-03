package util

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"

	"github.com/Scramjet911/learning-go/go-books/constant"
)

func ParseBody(r *http.Request, x interface{}) bool {
	if body, err := io.ReadAll(r.Body); err == nil {
		unmarshalErr := json.Unmarshal([]byte(body), x)
		if unmarshalErr == nil {
			return true
		}
	}
	return false
}

func GetRandomMeanResponse() string {
	randomIndex := rand.Intn(len(constant.MeanResponses))
	return constant.MeanResponses[randomIndex]
}
