{{ define "primary-key-auto-rollback" }}/*********************** PRIMARY KEY AUTO ***********************/
ALTER TABLE `{{.Name}}` 
    MODIFY COLUMN {{ template "field" .PK}};
/****************************************************************/{{ end }}