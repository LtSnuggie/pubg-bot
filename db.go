package main

import (
	"database/sql"
	"fmt"
)

type MysqlConnector struct {
	sql.DB
	config SqlConfig
	tables map[string]Table
}

func NewMysqlConnector(conf Config) *MysqlConnector {
	c := conf.Mysql
	d, err := sql.Open("mysql", c.User+":"+c.Pass+"@/"+c.DatabaseName)
	if err != nil {
		panic(err.Error())
	}
	t := make(map[string]Table)
	tables := c.Tables
	fmt.Println(tables)
	for _, table := range tables {
		t[table.Name] = table

	}
	m := &MysqlConnector{*d, c, t}
	return m
}
