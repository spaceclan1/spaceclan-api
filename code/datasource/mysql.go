package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	MainDb *sql.DB
)

func InitMysql() {
	_, found := os.LookupEnv("db_user")
	if !found {
		log.Fatal("environment Variable 'db_user' not found")
	}
	_, found = os.LookupEnv("db_pass")
	if !found {
		log.Fatal("environment Variable 'db_pass' not found")
	}
	_, found = os.LookupEnv("db_host")
	if !found {
		log.Fatal("environment Variable 'db_host' not found")
	}
	_, found = os.LookupEnv("db_port")
	if !found {
		log.Fatal("environment Variable 'db_port' not found")
	}
	var err error
	dbURL := os.ExpandEnv("${db_user}:${db_pass}@tcp(${db_host}:${db_port})/${db_name}?autocommit=true")
	//fmt.Print(dbURL)
	MainDb, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
}
