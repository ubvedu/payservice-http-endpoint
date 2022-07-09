package config

import (
    "core-payment-lesson/config"
    "core-payment-lesson/server"
    "github.com/golobby/container/v3"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "log"
)

func BuildDI(config *config.Config) {
    if err := container.Singleton(func() (server.PayServiceClient, error) {
        conn, err := grpc.Dial(":"+config.Grpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            return nil, err
        }
        return server.NewPayServiceClient(conn), nil
    }); err != nil {
        log.Fatalln(err)
    }
}