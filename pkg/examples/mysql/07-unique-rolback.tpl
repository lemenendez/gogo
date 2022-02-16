{{ define "unique-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if $ele.Unique -}}
/***************************** UNIQUE ***************************/
ALTER TABLE `{{$tablename}}_{{$ele.Name}}_idx`
    DROP INDEX `{{$tablename}}` (`{{$ele.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{- end -}}
{{- end -}} 
{{- end -}}