package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ClientMiddleWare

// ClientMiddleWare client middleware print server address、rpc timeout and conn timeout
func ClientMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get server info
		klog.Infof("server address: %v, rpc timeout: %v, connection timeout: %v\n",
			ri.To().Address(), ri.Config().RPCTimeout(), ri.Config().ConnectTimeout(),
		)
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
