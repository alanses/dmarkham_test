package responses

import (
	"encoding/json"
	"net/http"
	"test_dmarkham/blueprint/jokes/signatures"
)

type JokeResponse struct {
	Data *Data `json:"data"`
}

type Data struct {
	FullName   signatures.FullNameInfo `json:"full_name"`
	RandomJoke signatures.RandomJoke   `json:"random_joke"`
}

func NewJokeResponse(fullName *signatures.FullNameInfo, randomJoke *signatures.RandomJoke) *JokeResponse {
	return &JokeResponse{
		Data: &Data{
			FullName:   *fullName,
			RandomJoke: *randomJoke,
		},
	}
}

func (jokeResponse *JokeResponse) ToJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(*jokeResponse)
}
