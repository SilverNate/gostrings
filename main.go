package main

import "github.com/navarasu/gostrings"

func main() {
	text := gostrings.Squish(" foo   bar    \\n   \\t   boo")
	println(text)
}
