module github.com/mobile-core/rcmd/pkg/cfg

replace github.com/mobile-core/rcmd/pkg/fileutil => ../fileutil

go 1.17

require (
	github.com/mobile-core/rcmd/pkg/fileutil v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)
