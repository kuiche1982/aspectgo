package main

import (
	"fmt"
	"regexp"

	asp "github.com/AkihiroSuda/aspectgo/aspect"
)

// ExampleAspect implements interface asp.Aspect
type ExampleAspect struct {
}

// Executed on compilation-time
func (a *ExampleAspect) Pointcut() asp.Pointcut {
	s := "fmt\\.Print*"
	return asp.NewCallPointcutFromRegexp(s)
}

// Executed ONLY on runtime
func (a *ExampleAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()
	// nop
	res := ctx.Call(args)
	return res
}

// ExampleAspect implements interface asp.Aspect
type ExampleAspect2 struct {
}

// Executed on compilation-time
func (a *ExampleAspect2) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("github.com/AkihiroSuda/aspectgo/example/hello2")
	s := pkg + ".*"
	return asp.NewCallPointcutFromRegexp(s)
}

// Executed ONLY on runtime
func (a *ExampleAspect2) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()
	fmt.Println("BEFORE hello")
	res := ctx.Call(args)
	fmt.Println("AFTER hello")
	return res
}
