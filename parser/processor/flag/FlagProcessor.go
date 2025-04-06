package flag

import (
	"flag"
	"fmt"
	"go-in-progress/parser/common"
	ctx "go-in-progress/parser/strategy/context"
	"go-in-progress/parser/strategy/factory"
)

const (
	DirectoriesFlag = "d"
	FilesFlag       = "f"
	SearchFlag      = "s"
)

func NewFlags() common.Flags {
	flags := getFlags()
	return flags
}

func getFlags() common.Flags {
	directories := flag.String(DirectoriesFlag, "", "Enable directory mode (concise)")
	files := flag.String(FilesFlag, "", "Enable file mode (concise)")
	search := flag.String(SearchFlag, "", "Enable search processor (concise)")

	flag.Parse()

	return common.Flags{
		Directories: *directories,
		Files:       *files,
		Search:      *search,
	}
}

func ValidateInput(context ctx.Context) {
	flags := NewFlags()
	strategy, err := factory.NewStrategyFromFlags(flags)

	if err != nil {
		fmt.Printf("Validation Error: [%s]\n", err)
	} else {
		context.SetStrategy(strategy)
	}
}
