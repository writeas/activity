package main

import (
	"io/ioutil"

	"github.com/writeas/activity/tools/defs"
	"github.com/writeas/activity/tools/streams/gen"
)

func main() {
	allTypes := append(defs.AllCoreTypes, defs.AllExtendedTypes...)
	files, err := gen.GenerateConvenienceTypes(allTypes)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		err = ioutil.WriteFile(f.Name, f.Content, 0666)
		if err != nil {
			panic(err)
		}
	}
}
