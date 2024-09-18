//go:build generate

package main

//go:generate ./aspectgo -w ./ -t github.com/AkihiroSuda/aspectgo/example/hello2 ./main_aspect.go
