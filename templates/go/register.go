package golang

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"text/template"

	"github.com/suhli/protoc-gen-validate/templates/goshared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	goshared.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
}
