package grpc_net

import (
	"errors"
	"fmt"
	"github.com/soybean/api"
	"google.golang.org/grpc"
	"io"
	"net"
)

type serverStream struct {
	handleServer PeerService_HandleServer
	grpcStream
}

func (stream *serverStream) send(message *api.PeerMessage, infect byte) error {
	message.Meta[GrpcInfect] = []byte{infect}
	return stream.handleServer.Send(message)
}

func (stream *serverStream) Meta() map[string]string {
	return stream.meta
}

func (stream *serverStream) Close() {
	stream.close = true
}

type peerServiceServer struct {
}

func (peerServer *peerServiceServer) Handle(handleServer PeerService_HandleServer) error {
	remote := parseRemote(handleServer.Context())
	if remote == "" {
		fmt.Println("warn! parse ip null")
		return errors.New("parse ip null")
	}

	stream := &serverStream{handleServer: handleServer}
	stream.meta = map[string]string{api.Remote: remote}
	stream.close = false

	defer GetSyNet().connectManager.Remove(stream.meta[api.Remote])

	for {
		if stream.close {
			break
		}

		message, err := handleServer.Recv()

		if err == io.EOF {
			continue
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("ser err %v", err))
			stream.Close()
			break
		}

		if handle := GetSyNet().handlers[message.Type]; handle != nil {
			go func() { invoke(message, handle, stream) }()
		}
	}
	return errors.New("ser close")
}

func serve(ip string, port int, opts ...grpc.ServerOption) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))

	if err != nil {
		panic(fmt.Sprintf("failed to listen: %s:%d, %v", ip, port, err))
	}

	fmt.Println(fmt.Sprintf("success to listen: %s:%d", ip, port))
	grpcServer := grpc.NewServer(opts...)

	RegisterPeerServiceServer(grpcServer, &peerServiceServer{})
	grpcServer.Serve(lis)
}
