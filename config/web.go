package config

import "fmt"

const DB_HOST = "127.0.0.1"
const DB_PORT = "1000"
const DB_USER = "user"
const DB_PASS = "passwd"
const DB_NAME = "whoknowsmeapp"

const RDB_HOST = "127.0.0.1"
const RDB_PORT = "1001"

const BASE_URL = ""
const FULL_URL = "http://localhost:1323"+BASE_URL

func UrlFor(url string) string {
	return fmt.Sprintf("%s%s", FULL_URL, url)
}
