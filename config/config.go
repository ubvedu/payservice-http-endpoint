package config

import (
    "core-payment-lesson/server"
    "github.com/golobby/container/v3"
    "google.golang.org/grpc"
    "log"
)

func BuildDI(conn *grpc.ClientConn) {
    if err := container.Singleton(func() server.PayServiceClient {
        return server.NewPayServiceClient(conn)
    }); err != nil {
        log.Fatalln(err)
    }
}
