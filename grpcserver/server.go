package grpcserver

import (
	"context"
	"fmt"
	"shadowshot-x/actuatorbuf/protobufs"
	"sync/atomic"
)

type PingServer struct {
	protobufs.UnimplementedPingRPCServer
}

type ActuatorServer struct {
	protobufs.UnimplementedActuatorServer
	P *atomic.Value
}

func (s *PingServer) PingCheck(ctx context.Context, pm *protobufs.PingMessage) (*protobufs.PingResponse, error) {
	return &protobufs.PingResponse{Ping: "server alive"}, nil
}

func (s *ActuatorServer) ContractStateCheck(ctx context.Context, pm *protobufs.ContractVariableState) (*protobufs.ContractVariableStateCheck, error) {
	fmt.Println(pm)
	fmt.Println(s.P.Load())
	response := &protobufs.ContractVariableStateCheck{StateCheck: true}
	return response, nil
}
