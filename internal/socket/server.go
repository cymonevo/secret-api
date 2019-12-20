package socket

import (
	"github.com/cymon1997/go-backend/internal/log"
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
