{{ define "fks-idx-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if and ($ele.TextProps) ( index $ele.TextProps "ftable" ) ( index $ele.TextProps "fkey" ) -}}
/*************************** FK INDEX ***************************/
DROP INDEX `{{$tablename}}_{{ $ele.Name}}_idx`
    ON `{{$tablename}}`;
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{ end -}} 
{{- end -}}
{{- end -}}