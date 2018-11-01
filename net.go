package soybean

import (
	"github.com/soybean/api"
	"github.com/soybean/grpc_net"
)

func GetNet() api.SoybeanNet {
	return grpc_net.GetSyNet()
}
