package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	MainDb *sql.DB
)

func init() {
	_, found := os.LookupEnv("db_user")
	if !found {
		fmt.Println("environment Variable 'db_user' not found")
		os.Exit(1)
	}
	_, found = os.LookupEnv("db_pass")
	if !found {
		fmt.Println("environment Variable 'db_pass' not found")
		os.Exit(1)
	}
	_, found = os.LookupEnv("db_host")
	if !found {
		fmt.Println("environment Variable 'db_host' not found")
		os.Exit(1)
	}
	_, found = os.LookupEnv("db_port")
	if !found {
		fmt.Println("environment Variable 'db_port' not found")
		os.Exit(1)
	}
	var err error
	dbURL := os.ExpandEnv("${db_user}:${db_pass}@tcp(${db_host}:${db_port})/${db_name}?autocommit=true")
	//fmt.Print(dbURL)
	MainDb, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
}
