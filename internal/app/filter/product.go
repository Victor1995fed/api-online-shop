package filter

import (
	"api-online-store/internal/app/model"
	"fmt"
	"strconv"
	"strings"
)

//Product ...
type Product struct {
	Count  int
	Offset int
	Tags   []int
	Sort   string
}

//Apply ...
func (p *Product) Apply(s *Product) string {
	m := &model.Product{}
	sql := "where true "
	if s.Tags != nil {
		sql += s.GetTagSql(s.Tags)

	}
	if s.Sort != "" {
		sql += s.GetSort(s.Sort)
	}
	sql += s.GetGroupBy(m.GetTableName() + ".id")
	if s.Count != 0 {
		sql += s.GetLimitSql(s.Count)
	}

	if s.Offset != 0 {
		sql += s.GetOffsetSql(s.Offset)
	}

	fmt.Println(sql)
	return sql
}

//GetTagSql ...
func (p *Product) GetTagSql(t []int) string {
	var sql string
	if len(t) > 0 {
		m := &model.Product{}
		id := strings.Trim(strings.Replace(fmt.Sprint(t), " ", ",", -1), "[]")
		sql = "and " + m.GetTableName() + ".id in " +
			"(select id from " + m.GetViewTags() + " where tag_id in (" + id + "))"
	}
	return sql

}

//GetCountSql ...
func (p *Product) GetLimitSql(c int) string {
	if c < 0 {
		c = 1
	}
	if c > 100 {
		c = 100
	}
	count := strconv.Itoa(c)
	return " limit " + count
}

//GetOffsetSql ...
func (p *Product) GetOffsetSql(o int) string {
	offset := strconv.Itoa(o)
	return " offset " + offset
}

//GetSort ...
func (p *Product) GetSort(s string) string {

	m := map[string]bool{
		"id":    true,
		"title": true,
		"price": true,
	}

	var sql string
	var sortField string
	var sortType string

	if s[0:1] == "-" {
		sortField = s[1:]
		sortType = " desc "
	} else {
		sortField = s[0:]
		sortType = " asc "
	}

	if _, ok := m[sortField]; ok {
		sql = sql + " order by "
		sql = sql + sortField + sortType
	}

	return sql

}

//GetGroupBy ...
func (p *Product) GetGroupBy(field string) string {
	return " GROUP  by  " + field
}
