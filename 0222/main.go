package main

import "fmt"

func main() {
	var g Greeter
	g = Helloer{}
	g.Greet()
}

type Greeter interface{ Greet() }

type Helloer struct{}
type Goodbyer struct{}

var _ Greeter = Helloer{}  // Helloer implements Greeter
var _ Greeter = Goodbyer{} // Goodbyer implemnets Greeter

func (Helloer) Greet()  { hello() }
func (Goodbyer) Greet() { goodbye() }

func hello()   { fmt.Println("hello") }
func goodbye() { fmt.Println("goodbye") }
