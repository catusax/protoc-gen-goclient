package main

import (
	"github.com/catusax/protoc-gen-goclient/protobuf"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	protogen.Options{
		ParamFunc: protobuf.Flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			protobuf.GenerateClientFile(gen, f)
		}
		return nil
	})
}
