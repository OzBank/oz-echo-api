package api

import (
	"encoding/json"
	"net/http"

	"github.com/OzBank/oz-echo-core/echo"
)

type server struct {
	service *echo.EchoService
}

func NewServer() *server {
	return &server{
		service: echo.NewEchoService(),
	}
}

func (s *server) Echo(w http.ResponseWriter, r *http.Request) {

	var phrase Phrase
	if err := json.NewDecoder(r.Body).Decode(&phrase); err != nil {
		panic(err)
	}

	resp := Echo{
		Response: phrase.Phrase,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
