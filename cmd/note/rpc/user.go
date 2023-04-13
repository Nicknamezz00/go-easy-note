package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/kitex_gen/user/userservice"
	"go-easy-note/pkg/constants"
	"go-easy-note/pkg/errno"
	"go-easy-note/pkg/middleware"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	cli, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleWare),
		client.WithInstanceMW(middleware.ClientMiddleWare),
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
	userClient = cli
}

func GetUsers(ctx context.Context, req *user.GetUsersRequest) (map[int64]*user.User, error) {
	resp, err := userClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	ret := make(map[int64]*user.User)
	for _, u := range resp.Users {
		ret[u.UserId] = u
	}
	return ret, nil
}
