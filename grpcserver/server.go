package grpcserver

import (
	"context"
	"fmt"
	"sync/atomic"

	pkg "github.com/shadowshot-x/ActuatorBuf-grpc/pkg/simpleVariableActuate"
	"github.com/shadowshot-x/ActuatorBuf-grpc/protobufs"
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
	desiredState := s.P.Load().(pkg.SimpleVariable)
	currentState := pkg.ConvertToSimpleVariable(pm.Var1, pm.Var2)
	check, msg, err := desiredState.StateCheck(*currentState)
	if err != nil {
		return nil, err
	}
	if !check {
		response := &protobufs.ContractVariableStateCheck{StateCheck: "false", StateMessage: msg}
		return response, nil
	}
	response := &protobufs.ContractVariableStateCheck{StateCheck: "true"}
	return response, nil
}
