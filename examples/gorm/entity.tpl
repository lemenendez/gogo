{{ define "package" }}package {{.package}}{{- end -}}
{{ define "entity" }}
{{ template "package" .Props }}
type {{ .Name | title }} struct {
	{{- range $index, $ele := .Fields }}
	{{ $ele.Name | SnakeToInitCaps }}   {{ $ele.MappedType }}
	{{- end }}
}

func ({{.Name | title}}) TableName() {
    return "{{.Name}}"
}

{{- end -}}
{{ template "entity" . }}