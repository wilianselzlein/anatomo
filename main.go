package main

import (
	structures "anatomo/structures"
	"sync"
)
var exp structures.Export
var data []structures.Export
var listValue []structures.Value
var idx int = 0 

func main() {
	var listItemAgrupado []structures.ItemAgrupado

	var wg sync.WaitGroup

	listProcedimentos := structures.ListAnatomos()
	for _, procedimento := range listProcedimentos {
		exp.Procedimento = procedimento
		listItemAgrupado = structures.ListItemAgrupado(procedimento.Codigo) 
		listValue = structures.ListValue(procedimento.Codigo) 

		for  _, itemAgrupado := range listItemAgrupado {
			exp.ItemAgrupado = itemAgrupado

			for day := 0; day <= 5; day++ {
				wg.Add(1)
				go worker(day, itemAgrupado.CodItem, &wg)
			}
			wg.Wait()
		}
	}

	structures.CalculateTotals(data)
	structures.CreateCSV(data)
}

func worker(day int, itemAgrupado int, wg *sync.WaitGroup) {
    defer wg.Done()

	for  _, value := range listValue {
		if value.Procedimento == itemAgrupado && value.Days == day{
			idx++
			exp.Index = idx
			value.Total = 0.00
			value.Percent = 0.00
			exp.Values = value
			data = append(data, exp)
		}
	}
}