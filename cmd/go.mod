module github.com/mobile-core/rcmd/cmd

go 1.16

replace (
	github.com/mobile-core/rcmd/pkg/cfg => ../pkg/cfg
	github.com/mobile-core/rcmd/pkg/fileutil => ../pkg/fileutil
	github.com/mobile-core/rcmd/pkg/k8s => ../pkg/k8s
	github.com/mobile-core/rcmd/pkg/log => ../pkg/log
	github.com/mobile-core/rcmd/pkg/ssh => ../pkg/ssh
)

require (
	github.com/mobile-core/rcmd/pkg/cfg v0.0.0-00010101000000-000000000000
	github.com/mobile-core/rcmd/pkg/k8s v0.0.0-00010101000000-000000000000
	github.com/mobile-core/rcmd/pkg/ssh v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.2.1
)
