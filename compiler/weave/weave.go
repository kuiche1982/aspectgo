// Package weave provides the weaver.
package weave

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/types"
	"log"
	"strings"

	"golang.org/x/tools/go/loader"

	"github.com/AkihiroSuda/aspectgo/aspect"
	"github.com/AkihiroSuda/aspectgo/compiler/parse"
	"github.com/AkihiroSuda/aspectgo/compiler/util"
	"github.com/AkihiroSuda/aspectgo/compiler/weave/match"
)

// Weave weaves aspect files to the target package and emit the woven files to wovenGOPATH.
func Weave(wovenGOPATH string, target string, afs []*parse.AspectFile) ([]string, error) {
	rewriters := []*rewriter{}
	for _, af := range afs {
		_, prog, err := loadTarget(target)
		if err != nil {
			return nil, err
		}
		matched, pointcutsByIdent, err := findMatchedThings(prog, af.Pointcuts)
		if err != nil {
			return nil, err
		}
		if util.DebugMode {
			log.Printf("Found %d matches", len(matched))
		}
		if len(matched) != len(pointcutsByIdent) {
			log.Fatal("impl error")
		}
		if len(matched) == 0 {
			return []string{}, nil
		}

		// rewrittenFnames1, err := rewriteAspectFile(wovenGOPATH, af)
		// if err != nil {
		// 	return nil, err
		// }

		rw := &rewriter{
			Program:          prog,
			Matched:          matched,
			Aspects:          pointcutMapToAspectMap(af.Pointcuts),
			PointcutsByIdent: pointcutsByIdent,
		}
		rewriters = append(rewriters, rw)
	}
	rewrittenFnames2, err := rewriteProgram(wovenGOPATH, rewriters[0])
	if err != nil {
		return nil, err
	}
	rewrittenFnames1 := []string{}
	return append(rewrittenFnames1, rewrittenFnames2...), nil
}

func pointcutMapToAspectMap(pointcuts map[*types.Named]aspect.Pointcut) map[aspect.Pointcut]*types.Named {
	aspects := make(map[aspect.Pointcut]*types.Named)
	for asp, pc := range pointcuts {
		x, ok := aspects[pc]
		if ok {
			log.Printf("pointcut conflict: %s vs %s", x, asp)
		}
		aspects[pc] = asp
	}
	return aspects
}

func findMatchedThings(prog *loader.Program, pointcuts map[*types.Named]aspect.Pointcut) (map[*ast.Ident]types.Object, map[*ast.Ident]aspect.Pointcut, error) {
	objs := make(map[*ast.Ident]types.Object)
	pointcutsByIdent := make(map[*ast.Ident]aspect.Pointcut)
	for _, pkgInfo := range prog.InitialPackages() {
		for id, obj := range pkgInfo.Uses {
			posn := prog.Fset.Position(id.Pos())
			if strings.HasSuffix(posn.Filename, "_aspect.go") {
				continue
			}
			for _, pointcut := range pointcuts {
				matched := match.ObjMatchPointcut(prog, id, obj, pointcut)
				if !matched {
					continue
				}
				if util.DebugMode {
					log.Printf("MATCHED %s:%d:%d: %s, pointcut=%s",
						posn.Filename, posn.Line, posn.Column,
						obj, pointcut)
				}
				objs[id] = obj
				xpt, ok := pointcutsByIdent[id]
				if ok {
					log.Printf("OVERRIDE %s:%d:%d: %s, pointcut=%s vs old=%s",
						posn.Filename, posn.Line, posn.Column,
						obj, pointcut, xpt)
				}
				pointcutsByIdent[id] = pointcut
			}
		}
	}
	return objs, pointcutsByIdent, nil
}

func loadTarget(target string) (*loader.Config, *loader.Program, error) {
	buldCtx := build.Default
	buldCtx.ToolTags = append(buldCtx.ToolTags, "aspect")
	conf := loader.Config{
		ParserMode: parser.ParseComments,
		Build:      &buldCtx,
	}
	conf.Import(target)
	prog, err := conf.Load()
	if err != nil {
		return nil, nil, err
	}
	return &conf, prog, nil
}
