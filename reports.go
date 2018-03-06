package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (r DBWrapper) GetGamertagByDiscordId(id string) string {
	rows, err := r.db.Query("Select gamertag from users where discordId = ?", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		return tag
	}
	return ""
}
func (r DBWrapper) SetGamertagByDiscordId(id, tag string) {
	var stmt string
	if r.GetGamertagByDiscordId(id) != "" {
		stmt = "DELETE FROM users WHERE discordId = ?"
		r.db.Exec(stmt, id)
	}
	stmt = "INSERT INTO users ( gamertag, discordId ) values (? ,?)"
	r.db.Exec(stmt, tag, id)
}
func (r DBWrapper) GetGamertagByAlias(alias string) string {
	rows, err := r.db.Query("Select gamertag from aliases where alias = ?", alias)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		return tag
	}
	return ""
}
func (r DBWrapper) SetGamertagByAlias(alias, tag string) {
	var stmt string
	if r.GetGamertagByAlias(alias) != "" {
		stmt = "DELETE FROM aliases WHERE alias = ?"
		r.db.Exec(stmt, alias)
	}
	stmt = "INSERT INTO aliases ( gamertag, alias ) values (? ,?)"
	r.db.Exec(stmt, tag, alias)
}
func (r DBWrapper) GetStat(stat, gamertag string) string {
	stmt := "SELECT " + stat + " FROM stats where gamertag = \"" + gamertag + "\" ORDER BY id DESC"
	rows, err := r.db.Query(stmt)
	if err != nil {
		fmt.Println(err.Error())
	}
	var value string
	if rows.Next() {
		rows.Scan(&value)
		return value
	}
	return ""
}
func (r DBWrapper) GetStats(gamertag string) string { return "" }
func (r DBWrapper) SetStats(details UserDetails) {
	s := r.GetStat("Matches_Played", details.Gamertag)
	if s == "" {
		s = "0"
	}
	cur, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println(err)
	}
	in, err := details.CheckStat("Matches Played")
	if err != nil {
		fmt.Println(err)
	}
	if cur < float64(in) {

		m := make(map[string]interface{})
		cols := make([]string, 0)
		vals := make([]interface{}, 0)
		m["gamertag"] = details.Gamertag
		cols = append(cols, "gamertag")
		vals = append(vals, details.Gamertag)
		colNum := 0
		for _, group := range details.Stats {
			for _, stat := range group.Data {
				colNum++
				n := strings.Join(strings.Split(stat.Name, " "), "_")
				cols = append(cols, n)
				vals = append(vals, stat.Value)
				m[n] = stat.Value
			}
		}
		fmt.Println(m)
		p := make([]string, 0)
		for range vals {
			p = append(p, "?")
		}
		stmt := "INSERT INTO stats ( " + strings.Join(cols, " , ") + " ) values ( " + strings.Join(p, " , ") + " )"
		res, err := r.db.Exec(stmt, vals...)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)
	} else {
		fmt.Println("No new matches")
	}
}
func (r DBWrapper) DiffStat(stat, gamertag, start, end string) string { return "" }
func (r DBWrapper) DiffSatas(gamertag, start, end string) string      { return "" }
func (r DBWrapper) AddWatch(gamertag string) {
	watching := r.ListWatch()
	for _, tag := range watching {
		if tag == gamertag {
			return
		}
	}
	stmt := "INSERT INTO watch ( gamertag ) values ( ? )"
	r.db.Exec(stmt, gamertag)
}
func (r DBWrapper) RemoveWatch(gamertag string) {
	stmt := "DELETE FROM watch WHERE gamertag = ?"
	r.db.Exec(stmt, gamertag)
}
func (r DBWrapper) ListWatch() []string {
	list := make([]string, 0)
	stmt := "SELECT * FROM watch"
	rows, err := r.db.Query(stmt)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		list = append(list, tag)
	}
	fmt.Println(list)
	return list
}

func (r DBWrapper) ReportStats(u UserDetails) string {
	if u.XUID == "" {
		return "User '" + u.Gamertag + "' not found."
	}
	report := "\n"
	for _, stat := range u.Stats {
		report += stat.Title + "\n"
		for _, data := range stat.Data {
			report += data.Name + ":" + strconv.FormatFloat(float64(data.Value), 'f', -1, 32) + "\n"
		}
	}
	return report
}
