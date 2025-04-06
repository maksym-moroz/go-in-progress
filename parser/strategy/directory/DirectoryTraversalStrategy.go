package directory

import (
	"fmt"
	"go-in-progress/parser/const"
	"go-in-progress/parser/strategy/common"
	"os"
	"path/filepath"
)

type TraversalStrategy struct {
	Directories []string
	SearchItem  string
}

func (receiver TraversalStrategy) Execute(channel chan string) chan string {
	for _, entry := range receiver.Directories {
		go traverseDirectory(channel, entry)
	}
	return channel
}

func traverseDirectory(channel chan string, entry string) {
	directories, err := os.ReadDir(entry)

	if err != nil {
		fmt.Printf("Unable to read the directory: [%s]\n", err)
		return
	}

	for _, directory := range directories {
		if directory.Name() == _const.MacOsMetadataFolder {
			continue
		}

		fullPath := filepath.Join(entry, directory.Name())

		if directory.IsDir() {
			traverseDirectory(channel, fullPath)
		} else {
			common.ReadFile(channel, fullPath)
		}
	}
}
