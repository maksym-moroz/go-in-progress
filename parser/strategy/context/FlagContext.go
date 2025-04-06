package context

import (
	"go-in-progress/parser/strategy/flag"
)

type FlagContext struct {
	strategy flag.Strategy
}

func (context *FlagContext) GetStrategy() flag.Strategy {
	return context.strategy
}

func (context *FlagContext) SetStrategy(strategy flag.Strategy) {
	context.strategy = strategy
}

func (context *FlagContext) Execute(channel chan string) chan string {
	return context.strategy.Execute(channel)
}

func NewFlagContext(strategy flag.Strategy) Context {
	return &FlagContext{
		strategy: strategy,
	}

}
