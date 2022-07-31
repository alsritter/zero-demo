#!/bin/bash

table=zero_demo
host=127.0.0.1
port=3306
dbname=my_test_db
username=root
passwd=123456
modeldir=./genModel

goctl model mysql datasource --url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" --table="${table}" --dir="${modeldir}" --style=goZero