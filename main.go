package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type any map[string]interface{}

// func getType(x interface{}) string {
// 	return reflect.TypeOf(x).String()
// }

func walk(m map[string]interface{}) {
	for key, val := range m {
		switch val := val.(type) {
		case map[string]interface{}:
			walk(val)
		case bool:
			m[key] = true
		case int64:
			m[key] = 1
		case float64:
			m[key] = 3.1416
		case []interface{}:
			m[key] = []interface{}{}
		default:
			m[key] = "placeholder"
		}
	}
}

func main() {
	inputFilename := "env.toml"
	input, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var p any
	if _, err := toml.DecodeReader(input, &p); err != nil {
		panic(err)
	}
	walk(p)
	outputFilename := "output.toml"
	output, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	if err := toml.NewEncoder(output).Encode(p); err != nil {
		panic(err)
	}
	fmt.Printf("Example written to %s\n", outputFilename)
}
