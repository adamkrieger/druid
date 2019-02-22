package walker

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

func ProcessTrees(inputData []byte) (nodes map[string]*Node, edges map[string]*Edge) {

	entityTrees := EntityTrees{}

	yaml.Unmarshal(inputData, &entityTrees)

	log.Println(entityTrees)
}
