package helpers

import (
	"fmt"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func NewGrpcWeb(s *grpc.Server, port string, origin string) *http.Server {
	wrapped := grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(or string) bool {
		return origin == or
	}))

	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrapped.ServeHTTP(resp, req)
	}

	httpsServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(handler),
	}

	return &httpsServer
}
