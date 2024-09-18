package pkg

import (
	"fmt"
	"regexp"

	asp "github.com/AkihiroSuda/aspectgo/aspect"
)

// IAspect will be woven
type IAspect struct {
}

func (a *IAspect) Pointcut() asp.Pointcut {
	s := regexp.QuoteMeta("(github.com/AkihiroSuda/aspectgo/example/nonmainpkg/pkg.I).Foo")
	return asp.NewCallPointcutFromRegexp(s)
}
func (a *IAspect) Advice(ctx asp.Context) []interface{} {
	return advice("IAspect", ctx)
}

func advice(name string, ctx asp.Context) []interface{} {
	args, recv := ctx.Args(), ctx.Receiver()
	fmt.Printf("%s BEFORE call (args=%+v, recv=%+v)\n",
		name, args, recv)
	res := ctx.Call(args)
	fmt.Printf("%s AFTER call (res=%+v)\n",
		name, res)
	return res
}
