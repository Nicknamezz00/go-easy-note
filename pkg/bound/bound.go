package bound

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/shirou/gopsutil/cpu"
	"go-easy-note/pkg/constants"
	"go-easy-note/pkg/errno"
	"net"
)

var _ remote.InboundHandler = &cpuLimitHandler{}

type cpuLimitHandler struct {
}

func NewCpuLimitHandler() remote.InboundHandler {
	return &cpuLimitHandler{}
}

// OnActive implements remote.InboundHandler interface
func (c *cpuLimitHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	return ctx, nil
}

// OnInactive implements remote.InboundHandler interface
func (c *cpuLimitHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	return ctx
}

//OnRead implements remote.InboundHandler interface
func (c *cpuLimitHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	percent := cpuPercent()
	klog.CtxInfof(ctx, "current cpu is %.2g", percent)

	if percent > constants.CPURateLimit {
		return ctx, errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}
	return ctx, nil
}

// OnMessage implements remote.InboundHandler interface
func (c *cpuLimitHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	return ctx, nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	return percent[0]
}
