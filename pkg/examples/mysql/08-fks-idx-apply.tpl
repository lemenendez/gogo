{{ define "fks-idx" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if gt (len $ele.FTable) 0 -}}
/*************************** FK INDEX ***************************/
CREATE INDEX `{{$tablename}}_{{ $ele.Name}}_idx`
    ON `{{$tablename}}` (`{{$ele.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{ end -}} 
{{- end -}}
{{- end -}}