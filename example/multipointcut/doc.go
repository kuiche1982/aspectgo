//go:build generate

package main

//go:generate ./aspectgo -w ./ -t github.com/AkihiroSuda/aspectgo/example/multipointcut ./main_aspect.go
