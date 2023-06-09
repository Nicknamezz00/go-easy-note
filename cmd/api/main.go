package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"go-easy-note/cmd/api/handlers"
	"go-easy-note/cmd/api/rpc"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/pkg/constants"
	"go-easy-note/pkg/errno"
	"go-easy-note/pkg/tracer"
	"time"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()

	r := server.New(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if val, ok := data.(int64); ok {
				return jwt.MapClaims{constants.IdentityKey: val}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"code":   errno.SuccessCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVal handlers.UserParam
			if err := c.Bind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVal.UserName) == 0 || len(loginVal.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{
				UserName: loginVal.UserName,
				Password: loginVal.Password,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", authMiddleware.LoginHandler)
	user1.POST("/register", handlers.Register)

	note1 := v1.Group("/note")
	note1.Use(authMiddleware.MiddlewareFunc())
	note1.GET("/query", handlers.QueryNote)
	note1.POST("", handlers.CreateNote)
	note1.PUT("/:note_id", handlers.UpdateNote)
	note1.DELETE("/:note_id", handlers.DeleteNote)

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	r.Spin()
}
