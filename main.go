// Copyright 2022 mozyy. All rights reserved.
// Use of this source code is governed by a Apache-2.0
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	gengo "github.com/mozyy/protoc-gen-gorm/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {

	var (
		flags flag.FlagSet
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
