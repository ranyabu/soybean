package grpc_net

import (
	"context"
	"fmt"
	"github.com/soybean/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"sort"
	"strings"
	"sync"
)

const (
	GrpcInfect = "GRPC-INFECT"
)

type ConnectManager interface {
	Add(string, GrpcStream)
	Remove(string) GrpcStream
	Get(string) GrpcStream
	Foreach(func(string, GrpcStream))
}

func newConnectManager() ConnectManager {
	return &connectManager{mapping: make(map[string]GrpcStream)}
}

type connectManager struct {
	mapping map[string]GrpcStream
	lock    sync.RWMutex
}

func (self *connectManager) Add(remote string, stream GrpcStream) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.mapping[remote] = stream
}

func (self *connectManager) Remove(remote string) GrpcStream {
	self.lock.Lock()
	defer self.lock.Unlock()
	stream := self.mapping[remote]
	delete(self.mapping, remote)
	return stream
}

func (self *connectManager) Get(remote string) GrpcStream {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.mapping[remote]
}

func (self *connectManager) Foreach(f func(string, GrpcStream)) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	for _, v := range self.mapping {
		f(v.Meta()[api.Remote], v)
	}
}

type GrpcStream interface {
	api.Stream
	send(*api.PeerMessage, byte) error
}

type grpcStream struct {
	close bool
	meta  map[string]string
}

type SyNetGrpc struct {
	handlers       map[int32]api.Handler
	interceptors   []api.Interceptor
	connectManager ConnectManager
	ser            []grpc.ServerOption
	cli            []grpc.DialOption
	once           sync.Once
	viper          *viper.Viper
	lock           sync.RWMutex
}

func (self *SyNetGrpc) Multicast(msg *api.PeerMessage, remotes ...string) {
	for _, remote := range remotes {
		stream := self.get(remote)
		if stream != nil {
			stream.send(msg, 0)
		}
	}
}

func (self *SyNetGrpc) Broadcast(msg *api.PeerMessage, remotes ...string) {
	if len(remotes) != 0 {
		for _, remote := range remotes {
			stream := self.get(remote)
			if stream != nil {
				stream.send(msg, 0)
			}
		}
	} else {
		self.connectManager.Foreach(func(remote string, stream GrpcStream) { stream.send(msg, 0) })
	}
}

func (self *SyNetGrpc) Reg(handler api.Handler) {
	fmt.Println(fmt.Sprintf("--reg handler %s", handler.Name()))
	self.handlers[handler.Type()] = handler
}

func (self *SyNetGrpc) Add(interceptor api.Interceptor) {
	fmt.Println(fmt.Sprintf("--add interceptor %s", interceptor.Name()))
	self.interceptors = append(self.interceptors, interceptor)
	sort.Slice(syNet.interceptors, func(i, j int) bool {
		if syNet.interceptors[i].Index() < syNet.interceptors[j].Index() {
			return true
		}
		return false
	})
}

func (self *SyNetGrpc) Startup(viper *viper.Viper) {
	self.once.Do(func() {
		self.viper = viper
		go func() {
			serve(viper.GetString(api.ListenIp), viper.GetInt(api.ListenPort), self.ser...)
		}()
	})
}

func (self *SyNetGrpc) get(remote string) GrpcStream {
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.connectManager.Get(remote) == nil {
		stream, err := cli(viper.GetString(api.ListenIp), viper.GetInt(api.ListenPort), self.cli...)
		if err != nil {
			return nil
		}
		return stream
	}
	return self.connectManager.Get(remote)
}

var syNet = &SyNetGrpc{handlers: make(map[int32]api.Handler), connectManager: newConnectManager(), cli: []grpc.DialOption{grpc.WithInsecure()}}

func GetSyNet() *SyNetGrpc {
	return syNet
}

var parseRemote = func(context context.Context) string {
	grpcPeer, ok := peer.FromContext(context)

	if !ok {
		fmt.Println("warn! parse grpc context error")
		return ""
	}

	if grpcPeer.Addr == net.Addr(nil) {
		fmt.Println("warn! peer addr  nil")
		return ""
	} else {
		return strings.Split(grpcPeer.Addr.String(), ":")[0]
	}
}

var invoke = func(message *api.PeerMessage, handler api.Handler, stream GrpcStream) {
	fmt.Println(fmt.Sprintf("handler %s handle message", handler.Name()))

	for _, interceptor := range GetSyNet().interceptors {
		result := interceptor.Intercept(stream, message)
		if !result {
			fmt.Println(fmt.Sprintf("interceptor %s intercept message", interceptor.Name()))
			return
		}
	}
	handler.Handle(stream, message)
}
