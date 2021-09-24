package database

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	Db       string
}

var GetConnectionString = func(Config Config) (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", Config.User, Config.Password, Config.Host, Config.Db)
	return
}
