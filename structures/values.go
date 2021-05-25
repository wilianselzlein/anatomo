package structures

import (
	"log"
	"fmt"
	execSql "anatomo/execsql"
)

const FlNmValues string = "anatomo.sql"

type Value struct {
	Procedimento int
	Days int
	Quant float64
	Total float64
	Percent float64
}

//List Value...
func ListValue(procedimento int) []Value {
	var result []Value
	var value Value

	log.Println(fmt.Sprintf("%d\t%s\t", procedimento, "anatamo - get itens"))

	sqlStatementValue := execSql.Select(fmt.Sprintf(execSql.ReadSQL(FlNmValues), procedimento, procedimento))
	for sqlStatementValue.Next() {
		sqlStatementValue.Scan(&value.Procedimento, &value.Days, &value.Quant)
		result = append(result, value)
	}
	log.Println(fmt.Sprintf("%d valores...", len(result)))

	return result
 }