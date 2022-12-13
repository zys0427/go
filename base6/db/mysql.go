package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlZys struct {
	db *sql.DB
}

func (m *MysqlZys) Init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/zys?tls=skip-verify&autocommit=true")
	m.checkErr(err)
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)

	err = db.Ping()
	m.checkErr(err)
	m.db = db
}

func (m *MysqlZys) Query(sql string) {
	rows, err := m.db.Query(sql)
	m.checkErr(err)
	result := make(map[int]map[string]interface{})
	t := make([]map[string]interface{}, 0)
	i := 0
	for rows.Next() {
		r := make(map[string]interface{})
		var id int
		var name string
		err = rows.Scan(&id, &name)
		m.checkErr(err)
		r["id"] = id
		r["name"] = name
		i++
		result[i] = r
		t = append(t, r)
	}
	//return t
	//for k, v := range t {
	//	fmt.Println(k, v, v["id"], v["name"])
	//}
}

func (m *MysqlZys) GetOneUrlByCode(code string) string {
	s := "select url from url_info where code=(?)"
	var url string
	m.db.QueryRow(s, code).Scan(&url)
	defer m.db.Close()
	return url
}


func (m *MysqlZys) ExecZys(code string, url string) error {
	s := "insert into url_info(url,code) value(?,?)"
	_, err := m.db.Exec(s, url, code)
	m.checkErr(err)
	return err
}

func (m *MysqlZys) checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
