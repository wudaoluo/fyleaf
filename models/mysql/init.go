package mysql




import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(mysqlconn string,maxconn,idleconn int) {
	var err error
	db, err = sql.Open("mysql", mysqlconn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(idleconn)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
