{{ define "create_table_rollback"}}/*************************** DROP TABLE *************************/
DROP TABLE `{{.Name}}`;
/****************************************************************/
{{- end -}}