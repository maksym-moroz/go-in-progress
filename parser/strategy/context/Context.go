package context

import (
	"go-in-progress/parser/strategy/flag"
)

type Context interface {
	SetStrategy(strategy flag.Strategy)
	GetStrategy() flag.Strategy
	Execute(channel chan string) chan string
}
