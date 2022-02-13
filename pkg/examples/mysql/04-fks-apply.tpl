{{ define "fks" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if gt (len $ele.FTable) 0 -}}
/*************************** FK CONST ***************************/
ALTER TABLE `{{$tablename}}`
    ADD CONSTRAINT {{$tablename}}_const_fk_{{$ele.Name}}
    FOREIGN KEY (`{{$ele.Name}}`)
    REFERENCES `{{$ele.FTable}}` (`{{$ele.FKey}}`);
/****************************************************************/
{{ end }}{{- end -}}{{- end -}}