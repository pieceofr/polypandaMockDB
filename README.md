# README 

## Purpose
This progam is used to for Development stage of polypanda project. 
This program is used to create mock data in database info for polypanda project.

## Usage

- num is the number of records to create

```
    ./polypandaMockDB -num=100
```

## Config
create a file named "conf.yml" and fill out below info

```
    --- 
    database: 
    polypandadb: dbname
    #sqlendpoint: sqlhostname
    sqlpwd: sqluserpassword
    sqluser: sqluser
    pandatable: tabletoinsertdata
```