{{ define "unique-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if $ele.Unique -}}
/***************************** UNIQUE ***************************/
ALTER TABLE `{{$tablename}}`
    DROP INDEX `{{$tablename}}_{{$ele.Name}}_idx`;
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{- end -}}
{{- end -}} 
{{- end -}}