package templates

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/suhli/protoc-gen-validate/templates/cc"
	"text/template"

	"github.com/suhli/protoc-gen-validate/templates/ccnop"
	"github.com/suhli/protoc-gen-validate/templates/go"
	"github.com/suhli/protoc-gen-validate/templates/java"
	"github.com/suhli/protoc-gen-validate/templates/shared"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}

func Template(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"cc":    {makeTemplate("h", cc.RegisterHeader, params), makeTemplate("cc", cc.RegisterModule, params)},
		"ccnop": {makeTemplate("h", ccnop.RegisterHeader, params), makeTemplate("cc", ccnop.RegisterModule, params)},
		"go":    {makeTemplate("go", golang.Register, params)},
		"java":  {makeTemplate("java", java.Register, params)},
	}
}

func FilePathFor(tpl *template.Template) FilePathFn {
	switch tpl.Name() {
	case "h":
		return cc.CcFilePath
	case "ccnop":
		return cc.CcFilePath
	case "cc":
		return cc.CcFilePath
	case "java":
		return java.JavaFilePath
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(".validate." + tpl.Name())
			return &out
		}
	}
}
