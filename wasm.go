package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("helloWasm", js.FuncOf(helloWasm))
	select {}
}

func helloWasm(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Hello from WebAssembly!")
	return nil
}
