package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/kitex_gen/user/userservice"
	"go-easy-note/pkg/constant"
	"go-easy-note/pkg/errno"
	"go-easy-note/pkg/middleware"
	"time"
)

var UserClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}

	cli, err := userservice.NewClient(
		constant.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleWare),
		client.WithMiddleware(middleware.ClientMiddleWare),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // connection timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	UserClient = cli
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) error {
	resp, err := UserClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) error {
	resp, err := UserClient.CheckUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return nil
}
