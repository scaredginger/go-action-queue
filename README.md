# Go Action Queue

Allows a user to create a long list of actions, then serve them one by one or all at once

## Example usage

	package main

	import "fmt"
	import "strconv"
	import "github.com/scaredginger/go-action-queue"

	type printer struct {
		s string
	}

	func (p *printer) Execute() {
		fmt.Println(p.s)
	}

	func main() {
		var q = queue.MakeActionQueue()
		for i := 0; i < 9500; i++ {
			q.Append(&printer{s: strconv.Itoa(i)})
		}
		q.Empty()
	}

The queue accepts any object with an Execute method that accepts no arguments.
