package mysqlex

import (
	"container/list"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
	"time"
)

//----------------------------------------------------------------------------

var ERR_EMPTY_DATA_SOURCE_NAME = errors.New("Empty database source name.")
var ERR_EMPTY_SQL = errors.New("Empty SQL statement.")
var ERR_EMPTY_TABLE = errors.New("Empty table.")
var ERR_UNEXPECTED_ROWS = errors.New("Unexpected rows affected.")

//----------------------------------------------------------------------------

type Database struct {
	client *sql.DB
}

func NewDatabase(dataSourceName string, maxConns int) (*Database, error) {
	if len(dataSourceName) == 0 {
		return nil, ERR_EMPTY_DATA_SOURCE_NAME
	}

	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	client.SetMaxOpenConns(maxConns)
	client.SetMaxIdleConns(maxConns / 2)

	db := new(Database)
	db.client = client

	return db, nil
}

//----------------------------------------------------------------------------

func (db *Database) Close() error {
	return db.client.Close()
}

//----------------------------------------------------------------------------

func (db *Database) Count(table string) (int, error) {
	sql := "SELECT COUNT(*) FROM " + table + ";"

	rows, err := db.Select(sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, ERR_EMPTY_TABLE
	}

	n := 0
	err = rows.Scan(&n)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

//----------------------------------------------------------------------------

func (db *Database) Select(sql string) (*sql.Rows, error) {
	if len(sql) == 0 {
		return nil, ERR_EMPTY_SQL
	}

	return db.client.Query(sql)
}

//----------------------------------------------------------------------------

func (db *Database) Insert(sql string, rows int64) (int64, error) {
	if len(sql) == 0 {
		return 0, ERR_EMPTY_SQL
	}

	result, err := db.client.Exec(sql)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return affected, err
	}
	if affected != rows {
		return affected, ERR_UNEXPECTED_ROWS
	}

	id, err := result.LastInsertId()
	if err != nil {
		return affected, err
	}

	return id, nil
}

//----------------------------------------------------------------------------

func (db *Database) Exec(sql string) (int64, error) {
	if len(sql) == 0 {
		return 0, ERR_EMPTY_SQL
	}

	result, err := db.client.Exec(sql)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return affected, err
	}

	return affected, nil
}

//----------------------------------------------------------------------------

func (db *Database) Transaction() (*sql.Tx, error) {
	return db.client.Begin()
}

//----------------------------------------------------------------------------

func (db *Database) TypedSelect(sql string) (*sql.Rows, error) {
	if len(sql) == 0 {
		return nil, ERR_EMPTY_SQL
	}

	stmt, err := db.client.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query()
}

//----------------------------------------------------------------------------

func (db *Database) Encode(s string) string {
	r := strings.Replace(s, "\t", "\\t", -1)
	r = strings.Replace(r, "\n", "\\n", -1)
	r = strings.Replace(r, "\r", "\\r", -1)
	r = strings.Replace(r, "'", "\\'", -1)
	return r
}

func (db *Database) Dump(filename string, maxSize int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	ls, err := (func() (*list.List, error) {
		rows, err := db.Select("SHOW TABLE STATUS;")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		tables := list.New()
		cols, _ := rows.Columns()

		for rows.Next() {
			values := make([]interface{}, len(cols))
			valuePtrs := make([]interface{}, len(cols))
			for i := range cols {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				return tables, err
			}

			var m = make(map[string]interface{})
			for i, col := range cols {
				buf, okay := values[i].([]byte)
				if okay {
					m[col] = string(buf)
				} else {
					m[col] = values[i]
				}
			}
			if m["Engine"] != nil {
				tables.PushBack(m["Name"].(string))
			}
		}

		return tables, nil
	})()
	if err != nil {
		return err
	}

	err = (func() error {
		for e := ls.Front(); e != nil; e = e.Next() {
			name, okay := e.Value.(string)
			if !okay {
				continue
			}

			rows, err := db.Select("SHOW CREATE TABLE `" + name + "`;")
			if err != nil {
				return err
			}
			defer rows.Close()

			k := ""
			v := ""
			for rows.Next() {
				if err := rows.Scan(&k, &v); err != nil {
					return err
				}
				f.WriteString(v + ";\n")
			}
		}

		return nil
	})()
	if err != nil {
		return err
	}

	err = (func() error {
		for e := ls.Front(); e != nil; e = e.Next() {
			name, okay := e.Value.(string)
			if !okay {
				continue
			}

			sSQL := "SELECT * FROM `" + name + "`"
			if maxSize > 0 {
				sSQL += " LIMIT 0," + strconv.Itoa(maxSize)
			}
			sSQL += ";"

			rows, err := db.TypedSelect(sSQL)
			if err != nil {
				return err
			}
			defer rows.Close()

			cols, err := rows.ColumnTypes()

			s := "INSERT INTO `" + name + "` ("
			for i := 0; i < len(cols); i++ {
				if i > 0 {
					s += ","
				}
				s += "`" + cols[i].Name() + "`"
			}
			s += ")"
			f.WriteString(s)

			values := make([]interface{}, len(cols))
			valuePtrs := make([]interface{}, len(cols))
			for i := range cols {
				valuePtrs[i] = &values[i]
			}

			first := true
			for rows.Next() {
				if err := rows.Scan(valuePtrs...); err != nil {
					return err
				}

				s := ""
				if first {
					first = false
				} else {
					s += ","
				}

				s += "\nVALUES ("
				for i := 0; i < len(values); i++ {
					if i > 0 {
						s += ","
					}

					switch values[i].(type) {
					case bool:
						if values[i].(bool) {
							s += "true"
						} else {
							s += "false"
						}

					case int:
						s += strconv.Itoa(values[i].(int))
					case int64:
						s += strconv.FormatInt(values[i].(int64), 10)
					case float32:
						s += strconv.FormatFloat((float64)(values[i].(float32)), 'f', -1, 32)
					case float64:
						s += strconv.FormatFloat(values[i].(float64), 'f', -1, 64)

					case string:
						s += "'" + db.Encode(values[i].(string)) + "'"
					case []byte:
						s += "'" + db.Encode(string(values[i].([]byte))) + "'"

					case time.Time:
						s += "'" + db.Encode(values[i].(string)) + "'"

					default:
						s += "NULL"
					}
				}
				s += ")"
				f.WriteString(s)
			}
			f.WriteString(";\n")
		}
		return nil
	})()

	return nil
}

//----------------------------------------------------------------------------
