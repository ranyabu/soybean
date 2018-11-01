package grpc_net

import (
	"context"
	"fmt"
	"github.com/soybean/api"
	"google.golang.org/grpc"
	"io"
)

type clientStream struct {
	handleClient PeerService_HandleClient
	grpcStream
}

func (stream *clientStream) send(message *api.PeerMessage, infect byte) error {
	message.Meta[GrpcInfect] = []byte{infect}
	return stream.handleClient.Send(message)
}

func (stream *clientStream) Meta() map[string]string {
	return stream.meta
}

func (stream *clientStream) Close() {
	defer GetSyNet().connectManager.Remove(stream.meta[api.Remote])
	stream.handleClient.CloseSend()
}

func cli(remote string, port int, dialOpts ...grpc.DialOption) (GrpcStream, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", remote, port), dialOpts...)

	if err != nil {
		return nil, err
	}

	handleClient, err := NewPeerServiceClient(conn).Handle(context.Background())

	if err != nil {
		return nil, err
	}

	stream := &clientStream{handleClient: handleClient}
	stream.meta = map[string]string{api.Remote: remote}
	stream.close = false

	GetSyNet().connectManager.Add(remote, stream)

	go func() {
		defer GetSyNet().connectManager.Remove(stream.meta[api.Remote])
		for {
			if stream.close {
				fmt.Println(fmt.Sprintf("warn! cli err %v", err))
				break
			}

			message, err := handleClient.Recv()

			if err == io.EOF {
				continue
			}

			if err != nil {
				fmt.Println(fmt.Sprintf("warn! cli err %v", err))
				break
			}

			if handle := GetSyNet().handlers[message.GetType()]; handle != nil {
				go func() { invoke(message, handle, stream) }()
			}
		}
	}()

	return stream, nil
}
