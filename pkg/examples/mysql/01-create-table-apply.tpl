{{ define "comma" }}{{if eq .IsLast false}},{{end}}{{ end }}
{{ define "req" }}{{if eq .Required true}} NOT NULL{{end}}{{ end }}
{{ define "sign"}}{{if .Unsigned}} UNSIGNED{{end}}{{ end }}
{{ define "size"}}{{if .Size}}({{.Size}}){{end}}{{ end }}
{{ define "comment"}}{{ if .Comment}} COMMENT '{{.Comment}}'{{end}}{{end}}
{{ define "f_attr"}}{{ template "size" . }}{{ template "sign" . }}{{ template "req" . }}{{template "comment" .}}{{ end }}
{{ define "field"}}`{{.Name}}` {{.MappedType | upper }}{{template "f_attr" . }}{{end}}
{{ define "charset"}}{{ if .charset }}DEFAULT CHARACTER SET {{.charset}}{{end}}{{end}}
{{ define "collate"}}{{ if .collate }}COLLATE={{.collate}}{{end}}{{end}}
{{ define "create_table"}}/************************** CREATE TABLE ************************/
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
/******************************* {{ c1gi }} ******************************/
/****************************************************************/{{ end }}