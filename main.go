package main

import (
	"io/ioutil"
	"os"

	"github.com/scukonick/gocloud/lib"
)

func main() {
	code, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("wtf")
	}

	comp := lib.NewComputer()
	result, err := comp.Run(string(code))
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}

	os.Stdout.Write([]byte(result))
	os.Exit(0)
}
