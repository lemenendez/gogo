{{ define "primary-key" }}/************************** PRIMARY KEY *************************/
ALTER TABLE `{{.Name}}` 
    ADD CONSTRAINT pk_{{.Name}} 
    PRIMARY KEY (`{{.PK.Name}}`);
/******************************* {{ c1gi }} ******************************/
/****************************************************************/{{ end }}