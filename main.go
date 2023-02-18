package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
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
	var inputFilename string
	app := &cli.App{
		Name:  "FilePath",
		Usage: "Provide the proper path of the your env file",
		Action: func(cCtx *cli.Context) error {
			inputFilename = cCtx.Args().Get(0)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	input, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	body, _ := io.ReadAll(input)
	var p any
	if _, err := toml.Decode(string(body), &p); err != nil {
		panic(err)
	}
	walk(p)
	outputFilename := "env.example.toml"
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
