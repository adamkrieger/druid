package fileio

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"../walker"
)

//ReadFileContents - Reads the given path to []byte
func ReadFileContents(path string) (fileContents []byte, err error) {
	fileHndl, err := os.Open(path)
	lines := []string{}

	if err != nil {
		return fileContents, err
	}

	defer fileHndl.Close()

	rdr := bufio.NewReader(fileHndl)

	currentLine := []byte{}

	for {
		partial, isPrefix, err := rdr.ReadLine()

		if err != nil {
			break
		} else {
			currentLine = append(currentLine, partial...)
		}

		if isPrefix {
			continue
		} else if len(currentLine) == 0 {
			continue
		} else {
			lines = append(lines, string(currentLine))
			currentLine = []byte{}
		}

	}

	allLines := strings.Join(lines, "\n")

	return []byte(allLines), nil
}

//WriteNodesToJSON - Writes node array to destination path
func WriteNodesToJSON(entityMap map[string]*walker.Node) {
	entityFile, _ := os.Create("../seed_data/entity_generated.json")
	defer entityFile.Close()

	entityFile.Write([]byte{'['})
	first := true
	for _, entity := range entityMap {
		if !first {
			entityFile.Write([]byte{','})
		} else {
			first = false
		}
		marshaled, _ := json.Marshal(entity)
		entityFile.Write(marshaled)
	}
	entityFile.Write([]byte{']'})
}

//WriteEdgesToJSON - Writes edge array to destination path
func WriteEdgesToJSON(edgesMap map[string]*walker.Edge) {
	edgeFile, _ := os.Create("../seed_data/edges_generated.json")
	defer edgeFile.Close()

	edgeFile.Write([]byte{'['})
	first := true
	for _, entity := range edgesMap {
		if !first {
			edgeFile.Write([]byte{','})
		} else {
			first = false
		}
		marshaled, _ := json.Marshal(entity)
		edgeFile.Write(marshaled)
	}
	edgeFile.Write([]byte{']'})
}
