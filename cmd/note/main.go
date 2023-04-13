package note

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"go-easy-note/cmd/note/dal"
	"go-easy-note/cmd/note/rpc"
	"go-easy-note/kitex_gen/note/noteservice"
	"go-easy-note/pkg/bound"
	"go-easy-note/pkg/constants"
	"go-easy-note/pkg/middleware"
	mytracer "go-easy-note/pkg/tracer"
	"net"
)

func Init() {
	mytracer.InitJaeger(constants.NoteServiceName)
	rpc.InitRPC()
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()
	svr := noteservice.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.NoteServiceName}),
		server.WithMiddleware(middleware.CommonMiddleWare),
		server.WithMiddleware(middleware.ServerMiddleWare),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxQPS: 100, MaxConnections: 1000}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
