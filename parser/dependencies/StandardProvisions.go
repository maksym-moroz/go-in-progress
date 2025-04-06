package dependencies

import (
	"go-in-progress/parser/strategy/help"
	"time"
)

var StandardStrategy = &help.Strategy{}
var StandardTimeout = time.After(time.Millisecond * 80)
var StandardContent = make(chan string)
