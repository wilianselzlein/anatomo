package structures

import (
	"log"
	"fmt"
	execSql "anatomo/execsql"
)

const FlNmItemAgrupado string = "itens_agrupados.sql"

type ItemAgrupado struct {
	CodItem int
	Descricao string
}

//List Item Agrupado...
func ListItemAgrupado(procedimento int) []ItemAgrupado {
	var result []ItemAgrupado
	var itemAgrupado ItemAgrupado

	log.Println(fmt.Sprintf("%d\t%s", procedimento, "anatamo"))

	sqlStatementItemAgrupadoList := execSql.Select(fmt.Sprintf(execSql.ReadSQL(FlNmItemAgrupado), procedimento))
	for sqlStatementItemAgrupadoList.Next() {
		sqlStatementItemAgrupadoList.Scan(&itemAgrupado.CodItem, &itemAgrupado.Descricao)
		result = append(result, itemAgrupado)
	}
	log.Println(fmt.Sprintf("%d itens agrupados...", len(result)))

	return result
 }
