package config

import (
	"fmt"
	"time"
)

const (
	SQL_DATETIME_FORMAT = "2006-01-02 15:04:05"
	SQL_DATE_FORMAT     = "2006-01-02"
	API_ENDPOINT        = "https://api.waxsweden.org/v2"
	REDISTTL            = time.Duration(time.Hour * 24 * 90)
)

func BuildUrl(name string, params ...interface{}) string {
	switch name {
	case "get_actions":
		return fmt.Sprintf(API_ENDPOINT+"/history/get_actions?limit=%v&account=heroestaking&sort=asc&after=%v", params...)
	default:
		return ""
	}
}
