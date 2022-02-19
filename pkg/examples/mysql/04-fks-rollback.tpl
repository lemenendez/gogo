{{ define "fks-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if and ($ele.TextProps) ( index $ele.TextProps "ftable" ) ( index $ele.TextProps "fkey" ) -}}
/*************************** FK CONST ***************************/
ALTER TABLE `{{$tablename}}`
    DROP FOREIGN KEY (`{{$ele.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{ end -}}{{- end -}}{{- end -}}