// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qbd

import (
	"fmt"
	"strconv"
	"strings"
)

// CommaSpace is the separation
const CommaSpace = ", "

// MySQLQueryBuilder is the SQL build
type MySQLQueryBuilder struct {
	tokens []string
}

// Select will join the fields
func (qb *MySQLQueryBuilder) Select(fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "SELECT", strings.Join(fields, CommaSpace))
	return qb
}

// ForUpdate add the FOR UPDATE clause
func (qb *MySQLQueryBuilder) ForUpdate() QueryBuilder {
	qb.tokens = append(qb.tokens, "FOR UPDATE")
	return qb
}

// From join the tables
func (qb *MySQLQueryBuilder) From(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "FROM", strings.Join(tables, CommaSpace))
	return qb
}

// InnerJoin INNER JOIN the table
func (qb *MySQLQueryBuilder) InnerJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "INNER JOIN", table)
	return qb
}

// LeftJoin LEFT JOIN the table
func (qb *MySQLQueryBuilder) LeftJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "LEFT JOIN", table)
	return qb
}

// RightJoin RIGHT JOIN the table
func (qb *MySQLQueryBuilder) RightJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "RIGHT JOIN", table)
	return qb
}

// On join with on ternary
func (qb *MySQLQueryBuilder) On(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "ON", cond)
	return qb
}

// Where join the Where ternary
func (qb *MySQLQueryBuilder) Where(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "WHERE", cond)
	return qb
}

// And join the and ternary
func (qb *MySQLQueryBuilder) And(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "AND", cond)
	return qb
}

// Or join the or ternary
func (qb *MySQLQueryBuilder) Or(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "OR", cond)
	return qb
}

// InSQL join the IN (sql)
// Deprecate There is an injection risk here, recommended segmented execution
func (qb *MySQLQueryBuilder) InSQL(sql string) QueryBuilder {
	qb.tokens = append(qb.tokens, "IN", "(", sql, ")")
	return qb
}

// In join the IN (vals)
func (qb *MySQLQueryBuilder) In(vLen int) QueryBuilder {
	var vals []string
	for i := 0; i < vLen; i++ {
		vals = append(vals, "?")
	}
	qb.tokens = append(qb.tokens, "IN", "(", strings.Join(vals, CommaSpace), ")")
	return qb
}

// OrderBy join the Order by fields
// There is an injection risk here
func (qb *MySQLQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "ORDER BY", strings.Join(fields, CommaSpace))
	return qb
}

// Asc join the asc
func (qb *MySQLQueryBuilder) Asc() QueryBuilder {
	qb.tokens = append(qb.tokens, "ASC")
	return qb
}

// Desc join the desc
func (qb *MySQLQueryBuilder) Desc() QueryBuilder {
	qb.tokens = append(qb.tokens, "DESC")
	return qb
}

// Limit join the limit num
func (qb *MySQLQueryBuilder) Limit(limit int) QueryBuilder {
	qb.tokens = append(qb.tokens, "LIMIT", strconv.Itoa(limit))
	return qb
}

// Offset join the offset num
func (qb *MySQLQueryBuilder) Offset(offset int) QueryBuilder {
	qb.tokens = append(qb.tokens, "OFFSET", strconv.Itoa(offset))
	return qb
}

// GroupBy join the Group by fields
// There is an injection risk here
func (qb *MySQLQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "GROUP BY", strings.Join(fields, CommaSpace))
	return qb
}

// Having join the Having ternary
// There is an injection risk here
func (qb *MySQLQueryBuilder) Having(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "HAVING", cond)
	return qb
}

// Update join the update table
func (qb *MySQLQueryBuilder) Update(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "UPDATE", strings.Join(tables, CommaSpace))
	return qb
}

// Set join the set kv
func (qb *MySQLQueryBuilder) Set(kv ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "SET", strings.Join(kv, CommaSpace))
	return qb
}

// Delete join the Delete tables
func (qb *MySQLQueryBuilder) Delete(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "DELETE")
	if len(tables) != 0 {
		qb.tokens = append(qb.tokens, strings.Join(tables, CommaSpace))
	}
	return qb
}

// InsertInto join the insert SQL
func (qb *MySQLQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "INSERT INTO", table)
	if len(fields) != 0 {
		fieldsStr := strings.Join(fields, CommaSpace)
		qb.tokens = append(qb.tokens, "(", fieldsStr, ")")
	}
	return qb
}

// Values join the Values(vals)
func (qb *MySQLQueryBuilder) Values(vals ...string) QueryBuilder {
	valsStr := strings.Join(vals, CommaSpace)
	qb.tokens = append(qb.tokens, "VALUES", "(", valsStr, ")")
	return qb
}

// Subquery join the sub as alias
func (qb *MySQLQueryBuilder) Subquery(sub string, alias string) string {
	return fmt.Sprintf("(%s) AS %s", sub, alias)
}

// String join all tokens
func (qb *MySQLQueryBuilder) String(multiplex ...bool) string {
	var p = false
	if len(multiplex) > 0 {
		p = multiplex[0]
	}
	s := strings.Join(qb.tokens, " ")
	if p {
		qb.tokens = qb.tokens[:0]
	}
	return s
}
