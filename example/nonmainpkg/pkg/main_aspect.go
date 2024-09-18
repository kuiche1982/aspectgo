package pkg

import (
	"regexp"

	asp "github.com/AkihiroSuda/aspectgo/aspect"
)

// SAspect won't be woven, because it's not an "execution" pointcut
type SAspect struct {
}

func (a *SAspect) Pointcut() asp.Pointcut {
	s := regexp.QuoteMeta("(*github.com/AkihiroSuda/aspectgo/example/nonmainpkg/pkg.S).Foo")
	return asp.NewCallPointcutFromRegexp(s)
}
func (a *SAspect) Advice(ctx asp.Context) []interface{} {
	return advice("SAspect", ctx)
}
