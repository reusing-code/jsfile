package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"github.com/MJKWoolnough/gopherjs/files"
	"bufio"
	"bytes"
	"encoding/base64"
	"strings"
)

const layout string = `<div class="container"><h1>UPPERCASINATOR</h1>
<div class="custom-file">
  <input type="file" class="custom-file-input" id="bsfile">
  <label class="custom-file-label" for="bsfile" id="bsfile-label">Uppercase me, bitch!</label>
</div>


</div>`


func readFile(file files.File) {
	buf := bytes.Buffer{}
	writer := bufio.NewWriter(base64.NewEncoder(base64.URLEncoding, &buf))
	reader := files.NewFileReader(file);
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	output := ""
	for scanner.Scan() {
		output += scanner.Text()
		writer.WriteString(strings.ToUpper(scanner.Text()) + "\n")
	}
	writer.Flush()
	elem := dom.GetWindow().Document().CreateElement("a")
	elem.SetAttribute("download", file.Name())
	elem.SetAttribute("href", "data:text/plain;base64," + buf.String())
	elem.(*dom.HTMLAnchorElement).Click()
}

func main() {
	js.Global.Get("document").Call("write", layout)

	d := dom.GetWindow().Document()

	bsfilechooser := d.GetElementByID("bsfile")
	bsfilechooser.AddEventListener("change", false, func(event dom.Event) {
		input := event.Target().(*dom.HTMLInputElement)
		if len(input.Files()) > 0 {
			file := files.NewFile(input.Files()[0])
			println(file)
			dom.GetWindow().Document().GetElementByID("bsfile-label").SetInnerHTML(file.Name())
			go readFile(file)
		} else {
		dom.GetWindow().Document().GetElementByID("bsfile-label").SetInnerHTML("Choose File")
		}


	})

}