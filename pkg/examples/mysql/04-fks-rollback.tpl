{{ define "fks-rollback" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if gt (len $ele.FTable) 0 -}}
/*************************** FK CONST ***************************/
ALTER TABLE `{{$tablename}}`
    DROP FOREIGN KEY (`{{$ele.Name}}`);
/****************************************************************/
{{ end -}}{{- end -}}{{- end -}}