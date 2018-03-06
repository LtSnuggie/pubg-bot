package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DBWrapper struct {
	db     *sql.DB
	Config SqlConfig
}

type SqlConfig struct {
	User         string  `json:"user"`
	Pass         string  `json:"pass"`
	DatabaseName string  `json:"database"`
	Tables       []Table `json:"tables"`
}

type Table struct {
	Key      string   `json:"key'"`
	Name     string   `json:"name"`
	ColNames []string `json:"col_names"`
	ColTypes []string `json:"col_types"`
}

func NewDBWrapper(c Config) DBWrapper {
	d, err := sql.Open("mysql", c.Mysql.User+":"+c.Mysql.Pass+"@/"+c.Mysql.DatabaseName)
	if err != nil {
		panic(err.Error())
	}
	t := make(map[string]Table)
	tables := c.Mysql.Tables
	for _, table := range tables {
		t[table.Name] = table

	}
	w := DBWrapper{}
	w.db = d
	w.ShowTables()
	return w
}

func (d *DBWrapper) Close() {
	d.db.Close()
}

func (m *DBWrapper) ShowTables() {
	res, err := m.db.Query("show tables")
	if err != nil {
		fmt.Println(err.Error())
	}
	t := make(map[string]interface{}, 0)
	for res.Next() {
		var s string
		res.Scan(&s)
		t[s] = true
	}
	for _, table := range m.Config.Tables {
		if t[table.Name] == nil {
			m.CreateTable(table)
		}
	}
}

func (m *DBWrapper) CreateTable(table Table) {
	s := make([]string, 0)
	for i, name := range table.ColNames {
		if name != table.Key {
			tmp := strings.Split(name, " ")
			name = strings.Join(tmp, "_")
			s = append(s, name+" "+table.ColTypes[i])
		} else if table.Key != "" {
			s = append(s, name+" "+table.ColTypes[i]+" AUTO_INCREMENT, PRIMARY KEY ("+table.Key+")")
		}
	}
	stmt := "CREATE TABLE " + table.Name + " ( " + strings.Join(s, ", ") + " )"
	_, err := m.db.Exec(stmt)
	if err != nil {
		panic(err)
	}
}

func (m *DBWrapper) DropTable(table string) {
	stmt := "DROP TABLE " + table
	m.db.Exec(stmt)
}

func (m *DBWrapper) DeleteRow(query string) {

}
