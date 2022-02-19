{{ define "fks" }}
{{ $tablename := .Name }}
{{- range $idx, $ele := .Fields }}
{{- if and ($ele.TextProps) ( index $ele.TextProps "ftable" ) ( index $ele.TextProps "fkey" ) -}}
/*************************** FK CONST ***************************/
ALTER TABLE `{{$tablename}}`
    ADD CONSTRAINT fk_const_{{$tablename}}_{{$ele.Name}}
    FOREIGN KEY (`{{$ele.Name}}`)
    REFERENCES `{{$ele.TextProps.ftable }}` (`{{$ele.TextProps.fkey }}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/
{{ end }}{{- end -}}{{- end -}}