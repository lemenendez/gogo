{{ define "unique" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if $ele.Unique -}}/***************************** UNIQUE ***************************/
CREATE UNIQUE INDEX `{{$tablename}}_{{$ele.Name}}_idx`
    ON `{{$tablename}}` (`{{$ele.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/    
{{- end -}}{{- end -}}{{- end -}}