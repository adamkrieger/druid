package input

import (
	"errors"
	"fmt"
	"os"
)

//GetInputs - Evaluates CLI arguments
func GetInputs(args []string) (sourceFilePath, destDirPath string, err error) {
	if len(os.Args) != 3 {

		fmt.Println(`
Invalid Input. Expecting the following usage:

druid <SRC_FILE> <DEST_PATH>
- SRC_FILE: The data source
- DEST_PATH: The output directory to save generated data`)

		return "", "", errors.New("Invalid Input")
	}

	return os.Args[1], os.Args[2], nil
}
