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
    if err := config.BuildDI(conf); err != nil {
        log.Fatalln(err)
    }

    server := &http.Server{
        Addr:    ":" + conf.Http.Port,
        Handler: endpoint.NewRouter(),
    }
    log.Printf("Serving at: http://localhost:%s\n", conf.Http.Port)
    log.Fatalln(server.ListenAndServe())
}
