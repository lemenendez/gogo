
{{ define "fks-idx-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if gt (len $ele.FTable) 0 -}}
DROP INDEX `{{$tablename}}_{{ $ele.Name}}_idx`
    ON `{{$tablename}}`;
{{ end -}} 
{{- end -}}
{{- end -}}