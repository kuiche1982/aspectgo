//go:build generate

package main

//go:generate ./aspectgo -w ./ -t github.com/AkihiroSuda/aspectgo/example/receiver ./main_aspect.go
