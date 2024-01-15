package main

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strings"

	pb "github.com/endpoints/examples/bookstore/chat"
	"github.com/endpoints/examples/bookstore/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func getOpenAPIHandler() http.Handler {
	_ = mime.AddExtensionType(".svg", "image/svg+xml")
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	mux := runtime.NewServeMux()
	err := pb.RegisterChatServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50052", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	oa := getOpenAPIHandler()

	gatewayAddr := "localhost:8082"
	server := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				mux.ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
	log.Fatalln(fmt.Errorf("serving gRPC-Gateway server: %w", server.ListenAndServe()))
}
