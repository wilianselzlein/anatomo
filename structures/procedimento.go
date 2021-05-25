package structures

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
)

type Procedimento struct {
	Codigo int
	// Descricao string
	// Grupo string
	// Subgrupo Strin
}

//List Anatomos...
func ListAnatomos() []Procedimento {
	var result []Procedimento
	
	//var item Procedimento	
	//item.Codigo = 40601137
	//result = append(result, item)
    //result = append(result, Procedimento{40601013})

	file, err := os.Open("anatomopatologicos.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		//result = append(result, Procedimento{40601137})
		i, _ := strconv.Atoi(scanner.Text())
		result = append(result, Procedimento{i})
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	log.Println(fmt.Sprintf("%d procedimentos...", len(result)))
	
	return result
 }
