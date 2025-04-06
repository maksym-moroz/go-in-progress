package factory

import (
	"errors"
	"go-in-progress/parser/common"
	"go-in-progress/parser/strategy/composite"
	"go-in-progress/parser/strategy/directory"
	"go-in-progress/parser/strategy/file"
	flagStrategy "go-in-progress/parser/strategy/flag"
	"strings"
)

// NewStrategyFromFlags is a factory function that creates a Strategy based on the provided flags.
func NewStrategyFromFlags(flags common.Flags) (flagStrategy.Strategy, error) {
	switch {
	case flags.DirectoriesPresent() && flags.FilesPresent():
		return composite.Strategy{
			Strategies: []flagStrategy.Strategy{
				directory.TraversalStrategy{
					Directories: strings.Split(flags.Directories, " "),
					SearchItem:  flags.Search,
				},
				file.TraversalStrategy{
					Files:      strings.Split(flags.Files, " "),
					SearchItem: flags.Search,
				},
			},
		}, nil
	case flags.DirectoriesPresent():
		return directory.TraversalStrategy{
			Directories: strings.Split(flags.Directories, " "),
			SearchItem:  flags.Search,
		}, nil
	case flags.FilesPresent():
		return file.TraversalStrategy{
			Files:      strings.Split(flags.Files, " "),
			SearchItem: flags.Search,
		}, nil
	default:
		return nil, errors.New("no required processor passed")
	}
}
