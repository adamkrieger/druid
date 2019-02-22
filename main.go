package main

import (
	"os"

	"./fileio"
	"./input"
)

func main() {
	sourceFile, destDir, err := input.GetInputs(os.Args)

	if err != nil {
		os.Exit(1)
	}

	fileContents, err := fileio.ReadFileContents(sourceFile)

	if err != nil {
		os.Exit(2)
	}

	nodes, edges, err := processTrees(fileContents)

	if err != nil {
		os.Exit(4)
	}

	err := writeOutputToDestinationFiles(nodes, edges, destDir)

	if err != nil {
		os.Exit(5)
	} else {
		os.Exit(0)
	}
}
