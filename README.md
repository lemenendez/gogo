# Gogo

A code generator in go. Input files must be written in [YAML](https://en.wikipedia.org/wiki/YAML).

## Parameters

1. Configuration file.
2. Entity file.

## Configuration File

The configuration file is a **yml** file with the settings for the generators.

Example:

```yaml
name: gogo config file
generators:
  mysql:
    description: MySQL DDL Generator
    template-path: /home/leo/cloud/gogo/examples/mysql
    map: {big_int: bigint, date: date, date_time: datetime, int: int, medium_text: mediumtext, small_int: smallint, string: varchar, uuid: binary(16)}
  gorm:
    description: GO type with Generator with GORM 
    template-path: /home/leo/cloud/gogo/examples/gorm
    map: {big_int: bigint, date: date, date_time: datetime, int: int, medium_text: mediumtext, small_int: smallint, string: varchar, uuid: binary(16)}
```
Gogo can be used to generate code for different use cases. Each case is under *generators* property.
A generator consists of:
- A list of templates (*.tpl*) specified by **template-path**
- A field map specified by **map**

During the code generation all the field types specified in the entity file are replaced with the values specified in the map seccion of the generator.
For example the type *big_int* for the generator **A** can be mapped to *bigint* and to *int64* for the generator **B**

## Entity File

Example:

```yaml
name: user
comment: User entity
props: {charset: utf8mb4, collate: utf8mb4_unicode_ci}
fields:
  - name: id
    type: big_int
    pk: true
    unsigned: true
  - name: uid
    type: uuid
    required: true
    unique: true
  - name: email
    type: String
    size: 128
    required: true
    unique: true
  - name: updated_at
    type: date_time
    required: true
  - name: created_at
    type: date_time
    required: true
```
A entity file specifies the **name**, **comments**, **fields** and additional **props**.
All the props are available for the templates during the generation.

## Build

`go build -o ./gogo cmd/gogo.go`

## Command Line

`./gogo -c examples/gogo.yml examples/account.yml mysql migration`

- the **c** flag specifies the configuration file, if not provided then uses $HOME\gogo.yml
- **examples/account** param is the entity file
- **mysql** indicates the name of the generator
- **migration** indicates the name of the template

