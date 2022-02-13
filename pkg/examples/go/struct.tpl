{{ define "imports" }}
{{ if .goimports }}
imports (
	"{{.goimports}}"
)
{{- end }}
{{- end }}
{{ define "package" }}package {{.package}}{{- end }}
{{ define "struct" }}
{{ template "package" .Props }}
{{ template "imports" .Props }}
// {{ .Comment }}
type {{ .Name | title }} struct {
	{{- range $index, $ele := .Fields }}
	{{if $ele.Comment}}// {{ $ele.Comment }}{{end}}
	{{ $ele.Name | SnakeToInitCaps }} {{ $ele.MappedType }}
	{{- end }}
}
{{ end -}}
{{ template "entity" . }}