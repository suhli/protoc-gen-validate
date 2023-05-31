package goshared

const constTpl = `{{ $f := .Field }}{{ $r := .Rules }}{{ $e := .ErrorMessage}}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			{{- if isEnum $f }}
			err := {{ errCauseMessage . $e "value must equal " (enumVal $f $r.GetConst) }}
			{{- else }}
			err := {{ errCauseMessage . $e "value must equal " $r.GetConst }}
			{{- end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
