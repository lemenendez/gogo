{{ define "primary-key-auto" }}
ALTER TABLE `{{.Name}}` 
    MODIFY COLUMN {{ template "field" .PK}} 
    AUTO_INCREMENT;
{{- end -}}