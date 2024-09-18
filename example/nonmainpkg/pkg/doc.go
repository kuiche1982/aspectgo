//go:build generate

package pkg

//go:generate ./aspectgo -w ./ -t github.com/AkihiroSuda/aspectgo/example/nonmainpkg/pkg ./main_aspect.go ./main2_aspect.go
