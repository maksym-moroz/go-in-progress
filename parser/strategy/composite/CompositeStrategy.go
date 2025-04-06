package composite

import (
	"go-in-progress/parser/strategy/flag"
)

type Strategy struct {
	Strategies []flag.Strategy
}

func (receiver Strategy) Execute(channel chan string) chan string {
	for _, strategy := range receiver.Strategies {
		channel = strategy.Execute(channel)
	}
	return channel
}
