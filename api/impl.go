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

func (s *server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := Pong{
		Ping: s.service.Echo("pong"),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
