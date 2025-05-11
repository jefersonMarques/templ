//go:build js && wasm

package main

import (
	"bytes"
	"context"
	"syscall/js"

	"github.com/jefersonMarques/templ/ui/pages"
)

func main() {
	document := js.Global().Get("document")
	app := document.Call("getElementById", "app")

	var buf bytes.Buffer
	err := pages.Landing().Render(context.Background(), &buf)
	if err != nil {
		js.Global().Call("console.error", "Erro ao renderizar Landing:", err.Error())
		return
	}

	if app.Truthy() {
		app.Set("innerHTML", buf.String())
	}

	select {}
}
