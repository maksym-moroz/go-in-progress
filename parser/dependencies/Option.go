package dependencies

import (
	"go-in-progress/parser/strategy/context"
	"go-in-progress/parser/strategy/help"
	"time"
)

type Option func(dependencies *Dependencies)

func NewProvision(options ...Option) *Dependencies {
	provision := &Dependencies{}

	for _, option := range options {
		option(provision)
	}

	return provision
}

func WithStrategy(strategy *help.Strategy) Option {
	return func(dependencies *Dependencies) {
		dependencies.Context = context.NewFlagContext(strategy)
	}
}

func WithTimeout(timeout <-chan time.Time) Option {
	return func(dependencies *Dependencies) {
		dependencies.Timeout = timeout
	}
}

func WithContent(content chan string) Option {
	return func(dependencies *Dependencies) {
		dependencies.Content = content
	}
}
