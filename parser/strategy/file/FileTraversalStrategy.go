package file

import (
	"go-in-progress/parser/strategy/common"
)

type TraversalStrategy struct {
	Files      []string
	SearchItem string
}

func (receiver TraversalStrategy) Execute(channel chan string) chan string {
	for _, file := range receiver.Files {
		go common.ReadFile(channel, file)
	}

	return channel
}
