package parser

import (
	"go-in-progress/parser/dependencies"
	cntnt "go-in-progress/parser/processor/content"
	"go-in-progress/parser/processor/flag"
)

func Go() {
	provision := dependencies.ProvideDependencies()
	context := provision.Context
	timeout := provision.Timeout
	content := provision.Content

	flag.ValidateInput(context)
	channel := context.Execute(content)
	cntnt.ProcessContent(channel, timeout)
}
