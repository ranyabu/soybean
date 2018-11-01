package api

import "github.com/spf13/viper"

const (
	Remote     = "r"
	ListenIp   = "ip"
	ListenPort = "port"
)

type SoybeanNet interface {
	Multicast(*PeerMessage, ...string)
	Broadcast(*PeerMessage, ...string)

	Reg(Handler)
	Add(Interceptor)

	Startup(*viper.Viper)
}

type Stream interface {
	Meta() map[string]string
	Close()
}

type Plugin interface {
	Name() string
}

type Handler interface {
	Plugin
	Handle(Stream, *PeerMessage)
	Type() int32
}

type Interceptor interface {
	Plugin
	Intercept(Stream, *PeerMessage) bool
	Index() int32
}
