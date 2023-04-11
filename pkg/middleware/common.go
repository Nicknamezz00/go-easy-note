package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleWare

// CommonMiddleWare common middleware will print rpc info、real request and real response
func CommonMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// print real request
		klog.Infof("real request: %+v\n", req)
		// print remote service info
		klog.Infof("remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// print real response if everything is fine
		klog.Infof("real response: %+v\n", resp)
		return nil
	}
}
