package main

import (
	"parallax-wasm/controller"
	"syscall/js"
)

var (
	doc, star, moon, mountains_behind, text, btn, mountains_front, header js.Value
	elements                                                              []js.Value
)

func main() {

	doc = js.Global().Get("document")

	star = doc.Call("getElementById", "stars")
	elements = append(elements, star)

	moon = doc.Call("getElementById", "moon")
	elements = append(elements, moon)

	mountains_behind = doc.Call("getElementById", "mountains_behind")
	elements = append(elements, mountains_behind)

	text = doc.Call("getElementById", "text")
	elements = append(elements, text)

	btn = doc.Call("getElementById", "btn")
	elements = append(elements, btn)

	mountains_front = doc.Call("getElementById", "mountains_front")
	elements = append(elements, mountains_front)

	header = doc.Call("querySelector", "header")
	elements = append(elements, header)

	done := make(chan struct{})

	scrollEvent := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		value := js.Global().Get("window").Get("scrollY").Float()

		controller.Update(value, elements)

		return nil
	})

	defer scrollEvent.Release()

	doc.Call("addEventListener", "scroll", scrollEvent)

	<-done

}
