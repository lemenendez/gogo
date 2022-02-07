{{ define "unique" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if $ele.Unique -}}
CREATE UNIQUE INDEX `{{$tablename}}_{{$ele.Name}}_idx`
    ON `{{$tablename}}` (`{{$ele.Name}}`);
{{ end }}
{{- end -}} 
{{- end -}}