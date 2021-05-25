//Package structures...
package structures

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
)

type Export struct {
	Index int
	ItemAgrupado ItemAgrupado
	Procedimento Procedimento
	Values Value
}

//Calculate Totals...
func CalculateTotals(data []Export) {
	sum := make(map[int]float64)
	for _, item := range data {
		sum[item.ItemAgrupado.CodItem] += item.Values.Quant
	}

	for k, v := range sum {
		// fmt.Println(k, v)
		for index, item := range data {
			if item.ItemAgrupado.CodItem == k {
				item := &data[index]
				item.Values.Total = v
				if item.Values.Quant > 0 {
					item.Values.Percent = item.Values.Quant / item.Values.Total * 100
				}
			}
		}
	}
}

//CreateCSV...
func CreateCSV(data []Export) {
	file, err := os.Create("result.csv")

	if err != nil {
        log.Fatal(err)
    }

	defer file.Close()

	writer := csv.NewWriter(file)
    defer writer.Flush()

	file.WriteString("CodItem;Descricao;Anatomo;Dias;Quant;Total;Percent\n")
	for _, export := range data {
		_, err := file.WriteString(fmt.Sprintf(
			"%d;\"%s\";%d;%d;%f;%f;%f\n",
			//export.Index, 
			export.ItemAgrupado.CodItem, 
			export.ItemAgrupado.Descricao, 
			export.Procedimento.Codigo,
			export.Values.Days,
			export.Values.Quant,
			export.Values.Total,
			export.Values.Percent))

		if err != nil {
			log.Fatal(err)
		}
    }
}