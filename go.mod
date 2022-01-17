module github.com/mobile-core/rcmd

go 1.16

replace (
	github.com/mobile-core/rcmd/cmd => ./cmd
	github.com/mobile-core/rcmd/pkg/cfg => ./pkg/cfg
	github.com/mobile-core/rcmd/pkg/k8s => ./pkg/k8s
	github.com/mobile-core/rcmd/pkg/log => ./pkg/log
	github.com/mobile-core/rcmd/pkg/ssh => ./pkg/ssh
)

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/mobile-core/rcmd/cmd v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
