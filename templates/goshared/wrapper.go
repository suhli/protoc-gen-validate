package goshared

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}{{ $e := .ErrorMessage}}

	if wrapper := {{ accessor . }}; wrapper != nil {
		{{ render (unwrap . "wrapper") }}
	} {{ if .MessageRules.GetRequired }} else {
		err := {{ errCauseMessage . $e "value is required and must not be nil." }}
		if !all { return err }
		errors = append(errors, err)
	} {{ end }}
`
