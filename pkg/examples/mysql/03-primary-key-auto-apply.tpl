{{ define "primary-key-auto" }}/*********************** PRIMARY KEY AUTO ***********************/
ALTER TABLE `{{.Name}}` 
    MODIFY COLUMN {{ template "field" .PK}}
    AUTO_INCREMENT;
/******************************* {{ c1gi }} ******************************/
/****************************************************************/{{ end }}