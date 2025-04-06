package content

import (
	"fmt"
	"time"
)

func ProcessContent(channel chan string, timeout <-chan time.Time) {
	for {
		select {
		case msg, ok := <-channel:
			if !ok {
				return
			}
			fmt.Printf("Msg: %s\n", msg)
		case <-timeout:
			fmt.Println("Timeout waiting for messages")
			return
		}
	}
}
