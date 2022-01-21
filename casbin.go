package casbin

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2/middleware"
)

func AuthMiddleware(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Println("auth middleware in", req)
			reply, err = handler(ctx, req)
			fmt.Println("auth middleware out", reply)
			return
		}
	}
}
