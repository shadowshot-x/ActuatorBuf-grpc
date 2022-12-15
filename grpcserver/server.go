package grpcserver

import (
	"context"
	"shadowshot-x/actuatorbuf/protobufs"
)

type Server struct {
	protobufs.UnimplementedPingRPCServer
}

func (s *Server) PingCheck(ctx context.Context, pm *protobufs.PingMessage) (*protobufs.PingResponse, error) {
	return &protobufs.PingResponse{Ping: "server alive"}, nil
}
