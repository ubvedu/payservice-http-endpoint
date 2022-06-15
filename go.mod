module payservice-http-endpoint

go 1.18

require github.com/gorilla/mux v1.8.0

replace diLesson v0.0.0-20220606165212-a8211899bbad => github.com/TalismanFR/core-payment-lesson v0.0.0-20220606165212-a8211899bbad

require (
	diLesson v0.0.0-20220606165212-a8211899bbad
	github.com/golobby/container/v3 v3.2.1
)

require (
	github.com/kr/pretty v0.3.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
