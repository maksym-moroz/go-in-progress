package dependencies

import (
	"go-in-progress/parser/strategy/context"
	"time"
)

type Dependencies struct {
	Context context.Context
	Timeout <-chan time.Time
	Content chan string
}
