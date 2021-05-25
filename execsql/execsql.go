//Package execsql...
package execsql

import (
	"database/sql"
	"fmt"
	dbConfig "anatomo/dbconfig"
	_ "github.com/lib/pq"
	"io"
	"os"
	"bytes"
	"log"
)

var db *sql.DB
var err error

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Select...
func Select(commandSql string) *sql.Rows {
	log.Println(fmt.Sprintf("Accessing %s ... ", dbConfig.DbName))
     db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

    if err != nil {
        panic(err.Error())
    }  else {
        log.Println("Connected! Executing SQL...")
    }

	sqlStatement, err := db.Query(Replace(commandSql))
	checkErr(err)

    defer db.Close()
    return sqlStatement
 }


// ReadSQL... 
 func ReadSQL(fileNameCommandSql string) string {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open("execsql/" + fileNameCommandSql)
	if err != nil{
		log.Fatal(err) //panic(errs)
	}
	io.Copy(buf, f) 
	defer f.Close()
	//  b, err := ioutil.ReadAll(file)
	return string(buf.Bytes())
 }