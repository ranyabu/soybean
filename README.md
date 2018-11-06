# soybean

SOYBEAN 提供多播、广播两种P2P通信接口，同时提供拦截器给上层做收发拦截、处理器给上层做逻辑处理。

具体：
```go
type SoybeanNet interface {
	Multicast(*PeerMessage, ...string)
	Broadcast(*PeerMessage, ...string)

	Reg(Handler)
	Add(Interceptor)

	Startup(*viper.Viper)
}
```

使用示例：
```go
type MyHandler struct {
}

func (self *MyHandler) Name() string {
	return "MyHandler"
}

func (self *MyHandler) Handle(stream api.Stream, msg *api.PeerMessage) {
	fmt.Println(fmt.Sprintf("handle message from %s with %s", stream.Meta()[api.Remote], string(msg.Body)))
}

func (self *MyHandler) Type() int32 {
	return 1
}

type MyInterceptor struct {
}

func (self *MyInterceptor) Name() string {
	return "MyInterceptor"
}

func (self *MyInterceptor) Intercept(stream api.Stream, msg *api.PeerMessage) bool {
	if msg.Meta["REQ"] != nil {
		fmt.Println(fmt.Sprintf("intercept message from %s with %s", stream.Meta()[api.Remote], string(msg.Body)))
		GetNet().Multicast(&api.PeerMessage{Type: 1, Meta: map[string][]byte{"ACK": {0, 1, 2}}, Body: []byte("World")}, "0.0.0.0")
		return false
	}
	return true
}

func (self *MyInterceptor) Index() int32 {
	return 1
}

func init() {
	GetNet().Reg(&MyHandler{})
	GetNet().Add(&MyInterceptor{})
}

func TestAll(t *testing.T) {
	v := viper.GetViper()
	v.Set(api.ListenIp, "0.0.0.0")
	v.Set(api.ListenPort, 12181)
	GetNet().Startup(v)
	GetNet().Multicast(&api.PeerMessage{Type: 1, Meta: map[string][]byte{"REQ": {0, 1, 2}}, Body: []byte("Hello")}, "0.0.0.0")
	time.Sleep(time.Second * 5)
}
```