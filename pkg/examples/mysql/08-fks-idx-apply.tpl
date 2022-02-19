{{ define "fks-idx" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if and ($ele.TextProps) ( index $ele.TextProps "ftable" ) ( index $ele.TextProps "fkey" ) -}}
/*************************** FK INDEX ***************************/
CREATE INDEX `{{$tablename}}_{{ $ele.Name}}_idx`
    ON `{{$tablename}}` (`{{$ele.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{ end -}} 
{{- end -}}
{{- end -}}