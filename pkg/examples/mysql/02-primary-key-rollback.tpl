{{ define "primary-key-rollback" }}/************************** PRIMARY KEY *************************/
ALTER TABLE `{{.Name}}` DROP PRIMARY KEY;
/******************************* {{ c1gi }} ******************************/
/****************************************************************/{{ end }}