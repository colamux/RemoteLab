package DAL

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	info_table = "users"
	auth_table = "user_auths"
)

var (
	Db *sqlx.DB
)

// TODO: 连接池
func init() {
	db, err := sqlx.Connect("mysql", "db1:Db1_passwd@tcp(10.15.192.237:3306)/student")
	if err != nil {
		log.Panicln("MySql: ", err.Error())
		// fmt.Errorf("mysql connect error!")
		return
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	Db = db
	// defer db.Close()
}

// 查询正确但是数据为空时也应返回error
func GetPasswd(identity_type string, identifier string) (string, error) {
	querySql := `select password from user_auths where identity_type = "` +
		identity_type + `" and identifier = "` + identifier + `";`
	var ans []string
	if err := Db.Select(&ans, querySql); err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if len(ans) != 1 {
		return "", nil
	}

	return ans[0], nil
}

func GetVM(dest *[]string, name string) error {
	querySql := `select vm from users where name = "` + name + `";`

	return Db.Select(dest, querySql)
}

// 先前存在错误：httpstatus=500
// 1. 及时处理查询错误
// 2. sql语句中字符串加引号
// 3. dest为slice型的地址

// type of dest is *[]string
func GetAuth(dest *[]string, identity_type string, identifier string) error {
	querySql := `select password from user_auths where identity_type = "` +
		identity_type + `" and identifier = "` + identifier + `";`

	return Db.Select(dest, querySql)
}
