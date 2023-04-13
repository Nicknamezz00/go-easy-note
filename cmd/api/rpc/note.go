package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"go-easy-note/kitex_gen/note"
	"go-easy-note/kitex_gen/note/noteservice"
	"go-easy-note/pkg/constants"
	"go-easy-note/pkg/errno"
	"go-easy-note/pkg/middleware"
	"time"
)

var NoteClient noteservice.Client

func initNoteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	cli, err := noteservice.NewClient(
		constants.NoteServiceName,
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
	NoteClient = cli
}

func CreateNote(ctx context.Context, req *note.CreateNoteRequest) error {
	resp, err := NoteClient.CreateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return nil
}

func UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) error {
	resp, err := NoteClient.UpdateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return nil
}

func QueryNote(ctx context.Context, req *note.QueryNoteRequest) ([]*note.Note, int64, error) {
	resp, err := NoteClient.QueryNote(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return resp.Notes, resp.Total, nil
}

func DeleteNote(ctx context.Context, req *note.DeleteNoteRequset) error {
	resp, err := NoteClient.DeleteNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return nil
}
