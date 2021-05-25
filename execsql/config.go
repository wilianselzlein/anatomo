package execsql

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type Yaml struct {
    Schema     string `yaml:"schema"`
    ID         string `yaml:"id"`
    Version    string `yaml:"version"`
    Dependency []Dependency
}

type Dependency struct {
    Name    string
    Days 	int 
    Start 	string // date?
    End 	string
    Groups []string
}

var y *Yaml

func init() {
	log.Printf("Yaml: %+v\n", load())
}

func load() *Yaml {
	//y := Yaml{}
	if y == nil {

		yamlFile, err := ioutil.ReadFile("config.yaml") 
		if err != nil {
			log.Printf("yamlFile.Get err  #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &y) 

		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	return y
}

// Replace...
func Replace(commandSql string) string {

	y := load()

	if len(y.Dependency) == 0 {
		log.Fatal("No Dependency configuration in yaml file.")
	} else if len(y.Dependency[0].Groups) == 0 {
		log.Fatal("No Groups configuration in Dependency in yaml file.")
	}

	var result string = commandSql
	result = strings.ReplaceAll(result, "|days|", strconv.Itoa(y.Dependency[0].Days))
	result = strings.ReplaceAll(result, "|start|", y.Dependency[0].Start)
	result = strings.ReplaceAll(result, "|end|", y.Dependency[0].End)
	//fmt.Println(y.Dependency[0].Groups[0])

	return result

}