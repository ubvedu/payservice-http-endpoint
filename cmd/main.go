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
    confPath := os.Args[1]
    conf, err := coreConfig.Parse(confPath)
    if err != nil {
        log.Fatalln(err)
    }
    config.BuildDI(conf)

    port := "8080"
    log.Printf("Serving at: http://localhost:%s\n", port)
    log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), endpoint.NewRouter()))
}
