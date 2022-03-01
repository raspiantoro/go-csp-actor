package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/raspiantoro/go-actor/actor"
	"github.com/raspiantoro/go-actor/csp"
)

func main() {
	fmt.Println("execute csp")
	csp.ExecuteGoroutine()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("execute actor")
	actor.ExecuteActor()
	_, _ = console.ReadLine()
}
