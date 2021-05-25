// Package dbconfig...
package dbconfig

import (
    "fmt"
    "log"
)

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "postgres"

const DbName = "sc_antiglosa"

var Basic int

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func init() {
    log.Println("DB: " + DataSourceName)
}