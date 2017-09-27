#!/usr/bin/env bash

mysql-orm-gen -sql_file=./fin-stock.sql -orm_file=./fin-stock-gen.go -package_name="gen"