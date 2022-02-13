{{ define "primary-key" }}
ALTER TABLE `{{.Name}}` 
    ADD CONSTRAINT pk_{{.Name}} 
    PRIMARY KEY (`{{.PK.Name}}`);
{{- end -}}