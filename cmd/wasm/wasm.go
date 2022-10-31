//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"

	"syscall/js"
)

func formatJSON(input string) (string, error) {
	var raw any

	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return "", err
	}

	formatted, err := json.MarshalIndent(raw, "", "\t")
	if err != nil {
		return "", err
	}

	return string(formatted), nil
}

func jsonWrapper() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}

			return result
		}

		jsDoc := js.Global().Get("document")

		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}

			return result
		}

		jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output text area",
			}

			return result
		}

		inputJSON := args[0].String()

		fmt.Printf("input %s\n", inputJSON)

		pretty, err := formatJSON(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)

			result := map[string]any{
				"error": errStr,
			}

			return result
		}

		jsonOuputTextArea.Set("value", pretty)

		return nil
	})

	return jsonfunc
}

func main() {
	fmt.Println("Go Web Assembly")

	js.Global().Set("formatJSON", jsonWrapper())

	<-make(chan bool)
}