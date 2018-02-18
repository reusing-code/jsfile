package main

import "github.com/gopherjs/gopherjs/js"

const layout string = `<div class="container"><p>asd</p><button type="button" class="btn btn-secondary">Klick me!</button></div>`

func main() {
	js.Global.Get("document").Call("write", layout)
}