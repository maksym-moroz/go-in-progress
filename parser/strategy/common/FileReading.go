package common

import (
	"fmt"
	"os"
)

func ReadFile(channel chan string, file string) {
	fileContent, err := os.ReadFile(file)

	if err != nil {
		fmt.Printf("Unable to read the file [%s], error [%s]\n", file, err)
		return
	}

	channel <- string(fileContent)
}
