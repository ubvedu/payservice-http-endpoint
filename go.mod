module payservice-http-endpoint

go 1.18

require github.com/gorilla/mux v1.8.0

replace payservice-core v0.0.0-20220606165212-a8211899bbad => ./../payservice-core

require (
	github.com/golobby/container/v3 v3.2.1
	payservice-core v0.0.0-20220606165212-a8211899bbad
)

require (
	github.com/kr/pretty v0.3.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
)
