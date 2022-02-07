{{ define "comma" }}{{if eq .IsLast false}},{{end}}{{ end }}
{{ define "req" }}{{if eq .Required true}} not null{{end}}{{ end }}
{{ define "sign"}}{{if .Unsigned}} unsigned{{end}}{{ end }}
{{ define "size"}}{{if .Size}}({{.Size}}){{end}}{{ end }}
{{ define "f_attr"}}{{ template "size" . }}{{ template "sign" . }}{{ template "req" . }}{{ end }}
{{ define "field"}}`{{.Name}}` {{.MappedType}}{{template "f_attr" . }}{{end}}
{{ define "charset"}}{{ if .charset }}DEFAULT CHARACTER SET {{.charset}}{{end}}{{end}}
{{ define "collate"}}{{ if .collate }}COLLATE={{.collate}}{{end}}{{end}}
{{ define "create_table"}}
CREATE TABLE `{{.Name}}`
(
	{{- range $index, $ele := .Fields }}
        {{template "field" $ele}}{{ template "comma" $ele }}
	{{- end }}
) 
ENGINE=InnoDB
{{ template "charset" .Props }}
{{ template "collate" .Props }}
{{ if .Comment}}COMMENT '{{.Comment}}'{{end}};
{{- end -}}