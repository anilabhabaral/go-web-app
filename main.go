package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/greet", func() string {
		return "Hello world! \n"
	})
	m.RunOnAddr(":8080")
}
