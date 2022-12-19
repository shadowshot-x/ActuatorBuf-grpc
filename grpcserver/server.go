package grpcserver

import (
	"context"
	"fmt"
	"shadowshot-x/actuatorbuf/protobufs"
)

type PingServer struct {
	protobufs.UnimplementedPingRPCServer
}

type ActuatorServer struct {
	protobufs.UnimplementedActuatorServer
}

func (s *PingServer) PingCheck(ctx context.Context, pm *protobufs.PingMessage) (*protobufs.PingResponse, error) {
	return &protobufs.PingResponse{Ping: "server alive"}, nil
}

func (s *ActuatorServer) ContractStateCheck(ctx context.Context, pm *protobufs.ContractVariableState) (*protobufs.ContractVariableStateCheck, error) {
	fmt.Println(pm)
	response := &protobufs.ContractVariableStateCheck{StateCheck: true}
	return response, nil
}
