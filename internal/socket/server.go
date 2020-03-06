package socket

import (
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/googollee/go-socket.io"
)

type Server struct {
	engine *socketio.Server
}

func New() *Server {
	engine, err := socketio.NewServer(nil)
	if err != nil {
		log.FatalDetail(log.TagSocket, "error init server", err)
	}
	return &Server{
		engine: engine,
	}
}
