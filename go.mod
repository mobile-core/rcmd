module github.com/mobile-core/rcmd

go 1.16

replace (
	local.packages/cmd => ./cmd
	local.packages/k8s => ./pkg/k8s
	local.packages/ssh => ./pkg/ssh
)

require (
	github.com/spf13/cobra v1.2.1 // indirect
	k8s.io/client-go v0.22.2 // indirect
	local.packages/cmd v0.0.0-00010101000000-000000000000
	local.packages/k8s v0.0.0-00010101000000-000000000000 // indirect
	local.packages/ssh v0.0.0-00010101000000-000000000000 // indirect
)
