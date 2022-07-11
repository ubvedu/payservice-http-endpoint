package main

import (
    coreConfig "core-payment-lesson/config"
    _ "embed"
    "fmt"
    "log"
    "net/http"
    "os"
    endpoint "payservice-http-endpoint"
    "payservice-http-endpoint/config"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: go run ./cmd <config path>")
        os.Exit(1)
    }
    confPath := os.Args[1]
    conf, err := coreConfig.Parse(confPath)
    if err != nil {
        log.Fatalln(err)
    }
    config.BuildDI(conf)

    server := &http.Server{
        Addr:         ":" + conf.Http.Port,
        Handler:      endpoint.NewRouter(),
        ReadTimeout:  conf.Http.ReadTimeout,
        WriteTimeout: conf.Http.WriteTimeout,
    }
    log.Printf("Serving at: http://localhost:%s\n", conf.Http.Port)
    log.Fatalln(server.ListenAndServe())
}
