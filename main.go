package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"github.com/MJKWoolnough/gopherjs/files"
)

const layout string = `<div class="container"><p>asd</p><button id="btn1" type="button" class="btn btn-secondary">Klick me!</button>
<div class="custom-file">
  <input type="file" class="custom-file-input" id="bsfile">
  <label class="custom-file-label" for="bsfile" id="bsfile-label">Datei</label>
</div>

<input type="file" id="nobsfile">

</div>`

func main() {
	js.Global.Get("document").Call("write", layout)

	d := dom.GetWindow().Document()
	h := d.GetElementByID("btn1")

	h.AddEventListener("click", false, func(event dom.Event) {
		event.PreventDefault()
		h.SetInnerHTML("I am Clicked")
		println("This message is printed in browser console")
	})

	bsfilechooser := d.GetElementByID("bsfile")
	bsfilechooser.AddEventListener("change", false, func(event dom.Event) {
		input := event.Target().(*dom.HTMLInputElement)
		if len(input.Files()) > 0 {
			file := files.NewFile(input.Files()[0])
			println(file)
			dom.GetWindow().Document().GetElementByID("bsfile-label").SetInnerHTML(file.Name())
		} else {
		dom.GetWindow().Document().GetElementByID("bsfile-label").SetInnerHTML("Choose File")
		}


	})

}