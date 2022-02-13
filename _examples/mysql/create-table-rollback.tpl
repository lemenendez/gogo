{{ define "create_table_rollback"}}
DROP TABLE `{{.Name}}`;
{{- end -}}